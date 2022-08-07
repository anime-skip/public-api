package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
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

var (
	VALID   = true
	INVALID = false
)

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

	alerter := alerts.NewDiscordWebhookClient(
		config.DiscordAlertsURL(),
	)

	anilist := http.NewAnilistService()
	episodeService := postgres.NewEpisodeService(db)
	timestampService := postgres.NewTimestampService(db)
	timestampTypeService := postgres.NewTimestampTypeService(db)
	showService := postgres.NewShowService(db, anilist)

	episodes, err := episodeService.List(ctx, internal.EpisodesFilter{})
	episodeCount := len(episodes)
	utils.CheckErr(err)

	f, err := os.Create("report.log")
	utils.CheckErr(err)
	report := bufio.NewWriter(f)
	defer report.Flush()

	invalidCount := 0
	for i, episode := range episodes {
		log.I("%d/%d (%.2f%%)", i, episodeCount, 100*float64(i)/float64(episodeCount))
		isValid, err := validateEpisode(
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
			invalidCount++
		}
		// if i >= 10 {
		// 	break
		// }
	}

	summary := fmt.Sprintf("Processed %d episodes, %d invalid", len(episodes), invalidCount)
	fmt.Fprintf(report, "---\n%s\n", summary)
	if invalidCount > 0 {
		alerter.Notify(summary)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Fprintf(report, "Started: %v\nFinished: %v\nDuration: %v\n", start, end, duration)

	os.Exit(invalidCount)
}

func validateEpisode(
	ctx context.Context,
	report io.Writer,
	timestampService internal.TimestampService,
	timestampTypeService internal.TimestampTypeService,
	showService internal.ShowService,
	episode internal.Episode,
) (isValid bool, err error) {
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

	timestamps, isValid = validateEpisodeTimestamps(report, episode, timestamps)
	if isValid == INVALID {
		return INVALID, nil
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

	return VALID, nil
}

func validateEpisodeTimestamps(
	report io.Writer,
	episode internal.Episode,
	timestamps []internal.Timestamp,
) ([]internal.Timestamp, bool) {
	timestamps, isValid, errs := validation.EpisodeTimestamps(episode, timestamps)
	if !isValid {
		for _, err := range errs {
			fmt.Fprintf(report, "[VALIDATION ERROR]: %v\n", err)
		}
	}
	return timestamps, isValid
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

// Caches

var episodeCache = sync.Map{}

func getEpisode(ctx context.Context, episodeService internal.EpisodeService, id uuid.UUID) (internal.Episode, error) {
	key := id.String()
	cached, ok := episodeCache.Load(key)
	if ok {
		return cached.(internal.Episode), nil
	}
	v, err := episodeService.Get(ctx, internal.EpisodesFilter{ID: &id})
	if err != nil {
		return internal.Episode{}, err
	}
	episodeCache.Store(key, v)
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
