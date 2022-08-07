package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/alerts"
	"anime-skip.com/public-api/internal/config"
	"anime-skip.com/public-api/internal/http"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres"
	"anime-skip.com/public-api/internal/utils"
	"anime-skip.com/public-api/internal/validation"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
)

type EpisodeResult struct {
	episode        internal.Episode
	validationErrs []error
}

func main() {
	start := time.Now()
	// godotenv.Load(".env.local")
	godotenv.Load(".env.prod")

	ctx := context.Background()

	db := postgres.Open(
		config.DatabaseURL(),
		config.DatabaseDisableSSL(),
		nil,
		false,
	)

	alerter, err := alerts.NewDiscordAPIClient(
		config.DiscordBotToken(),
		config.DiscordAlertChannelID(),
	)
	utils.CheckErr(err)

	anilist := http.NewAnilistService()
	episodeService := postgres.NewEpisodeService(db)
	userService := postgres.NewUserService(db)
	timestampService := postgres.NewTimestampService(db)
	timestampTypeService := postgres.NewTimestampTypeService(db)
	showService := postgres.NewShowService(db, anilist)
	episodeURLService := postgres.NewEpisodeURLService(db)

	episodes, err := episodeService.List(ctx, internal.EpisodesFilter{})
	episodeCount := len(episodes)
	utils.CheckErr(err)

	f, err := os.Create("report.log")
	utils.CheckErr(err)
	report := bufio.NewWriter(f)
	defer report.Flush()

	invalidEpisodes := []EpisodeResult{}
	for i, episode := range episodes {
		log.I("%d/%d (%.2f%%)", i, episodeCount, 100*float64(i)/float64(episodeCount))
		isValid, err, validationErrs := validateEpisode(
			ctx,
			report,
			timestampService,
			timestampTypeService,
			showService,
			episode,
		)
		if err != nil {
			continue
		} else if !isValid {
			invalidEpisodes = append(invalidEpisodes, EpisodeResult{
				episode:        episode,
				validationErrs: validationErrs,
			})
		}
	}

	summary := fmt.Sprintf("Processed %d episodes, %d invalid", len(episodes), len(invalidEpisodes))
	fmt.Fprintf(report, "---\n%s\n", summary)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Fprintf(report, "Started: %v\nFinished: %v\nDuration: %v\n", start, end, duration)

	if len(invalidEpisodes) > 0 {
		sendNotification(
			ctx,
			alerter,
			userService,
			showService,
			timestampService,
			episodeURLService,
			summary,
			invalidEpisodes,
		)
	}
	report.Flush()
	os.Exit(len(invalidEpisodes))
}

func validateEpisode(
	ctx context.Context,
	report io.Writer,
	timestampService internal.TimestampService,
	timestampTypeService internal.TimestampTypeService,
	showService internal.ShowService,
	episode internal.Episode,
) (isValid bool, err error, validationErrs []error) {
	fmt.Fprintf(report, "---\n")
	fmt.Fprintf(report, "Episode ID: %s\n", episode.ID.String())
	fmt.Fprintf(report, "Name: %v\n", lo.If(episode.Name == nil, "<nil>").ElseF(func() string {
		return *episode.Name
	}))
	fmt.Fprintf(report, "Show ID: %s\n", episode.ShowID.String())

	timestamps, err := timestampService.List(ctx, internal.TimestampsFilter{
		EpisodeID: episode.ID,
	})
	if err != nil {
		fmt.Fprintf(report, "[ERROR] %v\n", err)
		return
	}
	originalTimestamps := utils.CopySlice(timestamps)
	fmt.Fprintf(report, "Timestamp count: %d\n", len(originalTimestamps))
	printTimestamps(ctx, report, timestampTypeService, "Original Timestamps", originalTimestamps)

	timestamps, isValid, validationErrs = validateEpisodeTimestamps(report, episode, timestamps)
	if !isValid {
		return false, nil, validationErrs
	}

	_, toCreate, toUpdate, toDelete := utils.ComputeSliceDiffs(
		timestamps,
		originalTimestamps,
		func(ts internal.Timestamp) string {
			if ts.ID == nil {
				// Timestamps we create during validation don't have an ID
				return utils.RandomString(16)
			}
			return ts.ID.String()
		},
		func(l internal.Timestamp, r internal.Timestamp) bool {
			return l.At != r.At || l.TypeID.String() != r.TypeID.String()
		},
	)
	if len(toCreate) > 0 || len(toUpdate) > 0 || len(toDelete) > 0 {
		printTimestamps(ctx, report, timestampTypeService, "WARN Timestamps Changed", timestamps)
	}
	if len(toCreate) > 0 {
		printTimestamps(ctx, report, timestampTypeService, "Creating", toCreate)
	}
	if len(toUpdate) > 0 {
		printTimestamps(ctx, report, timestampTypeService, "Updating", toUpdate)
	}
	if len(toDelete) > 0 {
		printTimestamps(ctx, report, timestampTypeService, "Deleting", toDelete)
	}

	return true, nil, nil
}

func validateEpisodeTimestamps(
	report io.Writer,
	episode internal.Episode,
	timestamps []internal.Timestamp,
) ([]internal.Timestamp, bool, []error) {
	timestamps, isValid, errs := validation.EpisodeTimestamps(episode, timestamps)
	if !isValid {
		for _, err := range errs {
			fmt.Fprintf(report, "[VALIDATION ERROR]: %v\n", err)
		}
	}
	return timestamps, isValid, errs
}

func printTimestamps(
	ctx context.Context,
	report io.Writer,
	timestampTypeService internal.TimestampTypeService,
	label string,
	timestamps []internal.Timestamp,
) {
	fmt.Fprintf(report, "%s:\n", label)
	for i, ts := range timestamps {
		var t internal.TimestampType
		t, err := getTimestampType(ctx, timestampTypeService, *ts.TypeID)
		if err != nil {
			fmt.Fprintf(report, "[ERROR] %v\n", err)
			return
		}
		fmt.Fprintf(report, "  %d. %s @ %.2f\n", i+1, t.Name, ts.At)
	}
}

func sendNotification(
	ctx context.Context,
	alerter internal.Alerter,
	userService internal.UserService,
	showService internal.ShowService,
	timestampService internal.TimestampService,
	episodeURLService internal.EpisodeURLService,
	summary string,
	invalidEpisodes []EpisodeResult,
) {
	log.I("Sending notification")
	threadID, err := alerter.CreateThread(
		summary,
		fmt.Sprintf("Timestamp Validation Results %s", time.Now().Format(time.RFC822)),
	)
	utils.CheckErr(err)
	fmt.Println(len(invalidEpisodes))
	for i, res := range invalidEpisodes {
		message := []string{
			"▁▁▁▁▁▁▁▁▁▁",
			fmt.Sprintf("_%d/%d_", i+1, len(invalidEpisodes)),
			fmt.Sprintf("**Episode**: %s (id=`%s`)", utils.ValueOr(res.episode.Name, "<nil>"), res.episode.ID.String()),
			fmt.Sprintf("**Number**: %s", utils.ValueOr(res.episode.Number, "<nil>")),
			fmt.Sprintf("**Absolute Number**: %s", utils.ValueOr(res.episode.AbsoluteNumber, "<nil>")),
			fmt.Sprintf("**Season**: %s", utils.ValueOr(res.episode.Season, "<nil>")),
		}
		user, err := getUser(ctx, userService, *res.episode.CreatedByUserID)
		if err != nil {
			message = append(message, fmt.Sprintf("**Created By**: ERR=%v (id=`%s`)", err, res.episode.CreatedByUserID.String()))
		} else {
			message = append(message, fmt.Sprintf("**Created By**: %s (id=`%s`)", user.Username, res.episode.CreatedByUserID.String()))
		}
		show, err := getShow(ctx, showService, *res.episode.ShowID)
		if err != nil {
			message = append(message, fmt.Sprintf("**Show**: ERR=%v (id=`%s`)", err, res.episode.ShowID.String()))
		} else {
			message = append(message, fmt.Sprintf("**Show**: %s (id=`%s`)", show.Name, res.episode.ShowID.String()))
		}

		timestamps, err := timestampService.List(ctx, internal.TimestampsFilter{EpisodeID: res.episode.ID})
		if err == nil {
			message = append(message, "**Contributors**")
			tsUserIDs := lo.Values(
				lo.Reduce(timestamps, func(m map[string]uuid.UUID, ts internal.Timestamp, i int) map[string]uuid.UUID {
					m[ts.UpdatedByUserID.String()] = *ts.UpdatedByUserID
					return m
				}, map[string]uuid.UUID{}),
			)
			for _, id := range tsUserIDs {
				u, err := getUser(ctx, userService, id)
				if err == nil {
					message = append(message, fmt.Sprintf(" • %s (id=`%s`)", u.Username, id.String()))
				}
			}
		}

		message = append(message, "**Validation Errors**")
		for _, err := range res.validationErrs {
			message = append(message, fmt.Sprintf(" • %v", err))
		}
		urls, err := episodeURLService.List(ctx, internal.EpisodeURLsFilter{EpisodeID: res.episode.ID})
		if err != nil {
			message = append(message, "**URLs**: <nil>")
		} else {
			message = append(message, "**URLs**")
			for _, url := range urls {
				message = append(message, fmt.Sprintf(" • %s", url.URL))
			}
		}

		message = append(message, "_Delete this message once fixed_", "▔▔▔▔▔▔▔▔▔▔")
		alerter.SendToThread(threadID, strings.Join(message, "\n"))
		time.Sleep(time.Second)
	}
}

// Caches

var showCache = sync.Map{}

func getShow(ctx context.Context, showService internal.ShowService, id uuid.UUID) (internal.Show, error) {
	key := id.String()
	cached, ok := showCache.Load(key)
	if ok {
		return cached.(internal.Show), nil
	}
	v, err := showService.Get(ctx, internal.ShowsFilter{ID: &id})
	if err != nil {
		return internal.Show{}, err
	}
	showCache.Store(key, v)
	return v, nil
}

var timestampTypeCache = sync.Map{}

func getTimestampType(ctx context.Context, timestampTypeService internal.TimestampTypeService, id uuid.UUID) (internal.TimestampType, error) {
	key := id.String()
	cached, ok := timestampTypeCache.Load(key)
	if ok {
		return cached.(internal.TimestampType), nil
	}
	v, err := timestampTypeService.Get(ctx, internal.TimestampTypesFilter{ID: &id})
	if err != nil {
		return internal.TimestampType{}, err
	}
	timestampTypeCache.Store(key, v)
	return v, nil
}

var userCache = sync.Map{}

func getUser(ctx context.Context, userService internal.UserService, id uuid.UUID) (internal.FullUser, error) {
	key := id.String()
	cached, ok := userCache.Load(key)
	if ok {
		return cached.(internal.FullUser), nil
	}
	v, err := userService.Get(ctx, internal.UsersFilter{ID: &id})
	if err != nil {
		return internal.FullUser{}, err
	}
	userCache.Store(key, v)
	return v, nil
}
