package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Constants

const (
	betterVRVBaseURL      = "https://parseapi.back4app.com"
	betterVRVAppIDHeader  = "x-parse-application-id"
	betterVRVAPIKeyHeader = "X-Parse-REST-API-Key"
)

// API Types

type betterVRVShow struct {
	ID        string    `json:"objectId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	VRVID     string    `json:"seriesId"`
	Title     string    `json:"seriesTitle"`
}

type betterVRVEpisode struct {
	ID             string              `json:"objectId"`
	CreatedAt      string              `json:"createdAt"`
	UpdatedAt      string              `json:"updatedAt"`
	VRVID          string              `json:"episodeId"`
	EpisodeTitle   string              `json:"episodeTitle"`
	Season         *int                `json:"seasonNumber"`
	AmbiguosNumber *int                `json:"episodeNumber"`
	Series         betterVRVSeriesLink `json:"series"`
	// When HasIntro=true, IntroStart and IntroEnd may exist, but they don't have to
	HasIntro   *bool    `json:"hasIntro"`
	IntroStart *float64 `json:"introStart"`
	IntroEnd   *float64 `json:"introEnd"`
	// When HasOutro=true, OutroStart and OutroEnd may exist, but they don't have to
	HasOutro   *bool    `json:"hasOutro"`
	OutroStart *float64 `json:"outroStart"`
	OutroEnd   *float64 `json:"outroEnd"`
	// When HasPostCredit=true, PostCreditStart and PostCreditEnd may exist, but they don't have to
	HasPostCredit   *bool    `json:"hasPostScene"`
	PostCreditStart *float64 `json:"postSceneStart"`
	PostCreditEnd   *float64 `json:"postSceneEnd"`
	// When HasPreview=true, PreviewStart and PreviewEnd may exist, but they don't have to
	HasPreview   *bool    `json:"hasPreview"`
	PreviewStart *float64 `json:"previewStart"`
	PreviewEnd   *float64 `json:"previewEnd"`
}

type betterVRVSeriesLink struct {
	Type      string `json:"__type"`
	ClassName string `json:"className"`
	ObjectID  string `json:"objectId"`
}

type betterVRVEpisodeResponse struct {
	Results []betterVRVEpisode `json:"results"`
}

type betterVRVShowResponse struct {
	Results []betterVRVShow `json:"results"`
}

// Service Definition

type BetterVRVThirdPartyService struct {
	appId  string
	apiKey string
	client *http.Client
}

func NewBetterVRVThirdPartyService(appId string, apiKey string) internal.ThirdPartyService {
	return &BetterVRVThirdPartyService{
		appId:  appId,
		apiKey: apiKey,
		client: &http.Client{
			Timeout: 1 * time.Second,
		},
	}
}

var (
	EmptyThirdPartyShow = internal.ThirdPartyShow{}
)

func (s *BetterVRVThirdPartyService) Name() string {
	return "http.BetterVRVThirdPartyService"
}

func (s *BetterVRVThirdPartyService) FindEpisodeByName(ctx context.Context, name string) ([]internal.ThirdPartyEpisode, error) {
	if s.apiKey == "" {
		log.W("BetterVRV skipped, apiKey is not set")
		return []internal.ThirdPartyEpisode{}, nil
	}
	if s.appId == "" {
		log.W("BetterVRV skipped, appId is not set")
		return []internal.ThirdPartyEpisode{}, nil
	}

	inputName := strings.ReplaceAll(name, "\"", "\\\"")
	inputName = strings.ReplaceAll(inputName, "\\", "\\\\")
	queryParams := map[string]string{
		"where": fmt.Sprintf("{ \"episodeTitle\": \"%s\" }", inputName),
	}
	req, err := s.createRequest(ctx, "/classes/Timestamps", queryParams, nil)
	if err != nil {
		return nil, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		log.W("Better VRV request failed: %v", err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Parse response
	response := &betterVRVEpisodeResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	// Map list
	episodes := []internal.ThirdPartyEpisode{}
	if response.Results != nil {
		for _, item := range response.Results {
			mappedItem := parseBetterVRVEpisode(item)
			if mappedItem != nil {
				episodes = append(episodes, *mappedItem)
			}
		}
	}

	return episodes, nil
}

func (s *BetterVRVThirdPartyService) createRequest(ctx context.Context, endpoint string, query map[string]string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", betterVRVBaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add(betterVRVAppIDHeader, s.appId)
	req.Header.Add(betterVRVAPIKeyHeader, s.apiKey)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req.WithContext(ctx), nil
}

func (s *BetterVRVThirdPartyService) fetchRemoteShowById(ctx context.Context, id string) (internal.ThirdPartyShow, error) {
	queryParams := map[string]string{
		"where": fmt.Sprintf("{ \"objectId\": \"%s\" }", id),
	}
	req, err := s.createRequest(ctx, "/classes/Series", queryParams, nil)
	if err != nil {
		return EmptyThirdPartyShow, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		log.W("Better VRV request failed: %v", err)
		return EmptyThirdPartyShow, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Parse response
	response := &betterVRVShowResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return EmptyThirdPartyShow, err
	}
	if len(response.Results) == 0 {
		return EmptyThirdPartyShow, nil
	}

	// Return the first matching show
	return mapBetterVRVShow(response.Results[0]), nil
}

// Mappers

func mapBetterVRVShow(episode betterVRVShow) internal.ThirdPartyShow {
	return internal.ThirdPartyShow{
		Name:      episode.Title,
		CreatedAt: &episode.CreatedAt,
		UpdatedAt: &episode.UpdatedAt,
	}
}

type Section struct {
	Start *internal.ThirdPartyTimestamp
	End   *internal.ThirdPartyTimestamp
}

const SAME_DIFF_THRESHOLD_SECONDS = 2

func (firstSection Section) isSame(secondSection Section) bool {
	if firstSection.Start != nil && secondSection.Start != nil && math.Abs(firstSection.Start.At-secondSection.Start.At) < SAME_DIFF_THRESHOLD_SECONDS {
		return true
	}
	if firstSection.End != nil && secondSection.End != nil && math.Abs(firstSection.End.At-secondSection.End.At) < SAME_DIFF_THRESHOLD_SECONDS {
		return true
	}
	return false
}

func createSection(hasSection *bool, sectionStart *float64, sectionEnd *float64, sectionDuration float64, typeID uuid.UUID) Section {
	var start *internal.ThirdPartyTimestamp
	var end *internal.ThirdPartyTimestamp

	if hasSection != nil && *hasSection {
		if sectionStart != nil && sectionEnd != nil {
			// Have both
			start = &internal.ThirdPartyTimestamp{
				At:     *sectionStart,
				TypeID: &typeID,
			}
			end = &internal.ThirdPartyTimestamp{
				At:     *sectionEnd,
				TypeID: &internal.TIMESTAMP_ID_UNKNOWN,
			}
		} else if sectionStart != nil {
			// Only have start
			start = &internal.ThirdPartyTimestamp{
				At:     *sectionStart,
				TypeID: &typeID,
			}
			if sectionDuration > 0 {
				end = &internal.ThirdPartyTimestamp{
					At:     *sectionStart + sectionDuration,
					TypeID: &internal.TIMESTAMP_ID_UNKNOWN,
				}
			}
		} else if sectionEnd != nil {
			// Only have end
			end = &internal.ThirdPartyTimestamp{
				At:     *sectionEnd,
				TypeID: &internal.TIMESTAMP_ID_UNKNOWN,
			}

			if sectionDuration > 0 {
				start = &internal.ThirdPartyTimestamp{
					At:     *sectionEnd - sectionDuration,
					TypeID: &typeID,
				}
			}
		}
	}
	return Section{
		Start: start,
		End:   end,
	}
}

// This should only be used when returning data, better vrv searches should NOT be standardized since it stores the data non-standard
func BetterVRVStandardizeEpisodeName(title string) (standardized string) {
	standardized = strings.ReplaceAll(title, "’", "'")
	standardized = strings.ReplaceAll(standardized, "‘", "'")
	standardized = strings.ReplaceAll(standardized, "–", "-")
	standardized = strings.ReplaceAll(standardized, "　", " ")
	standardized = strings.ReplaceAll(standardized, "“", "\"")
	standardized = strings.ReplaceAll(standardized, "”", "\"")
	standardized = strings.ReplaceAll(standardized, "…", "...")
	standardized = strings.ReplaceAll(standardized, "‼", "!!")
	return strings.TrimSpace(standardized)
}

func parseBetterVRVEpisode(input betterVRVEpisode) *internal.ThirdPartyEpisode {
	var number *string
	if input.AmbiguosNumber != nil {
		str := fmt.Sprintf("%v", *input.AmbiguosNumber)
		number = &str
	}
	var season *string
	if input.Season != nil {
		str := fmt.Sprintf("%v", *input.Season)
		season = &str
	}

	// Parse sections
	intro := createSection(input.HasIntro, input.IntroStart, input.IntroEnd, 90, internal.TIMESTAMP_ID_INTRO)
	outro := createSection(input.HasOutro, input.OutroStart, input.OutroEnd, 90, internal.TIMESTAMP_ID_CREDITS)
	postCredits := createSection(input.HasPostCredit, input.PostCreditStart, input.PostCreditEnd, 0, internal.TIMESTAMP_ID_CANON)
	preview := createSection(input.HasPreview, input.PreviewStart, input.PreviewEnd, 0, internal.TIMESTAMP_ID_PREVIEW)

	// Combine Sections
	timestamps := []internal.ThirdPartyTimestamp{}
	addPreview := func() {
		if preview.Start != nil {
			timestamps = append(timestamps, *preview.Start)
		}
		if preview.End != nil {
			timestamps = append(timestamps, *preview.End)
		}
	}

	// RECAP
	isRecap := preview.Start != nil && preview.Start.At == 0
	if isRecap {
		if preview.Start != nil {
			preview.Start.TypeID = &internal.TIMESTAMP_ID_RECAP
		}
		addPreview()
	}

	// INTRO
	if intro.Start != nil && intro.End != nil {
		if intro.Start.At < 0 {
			intro.Start.At = 0
		}
		timestamps = append(timestamps, *intro.Start, *intro.End)
	}

	// OUTRO
	if outro.Start != nil && outro.End != nil {
		timestamps = append(timestamps, *outro.Start, *outro.End)
	}

	// PREVIEW/POST CREDIT CANON
	if !isRecap {
		if postCredits.Start != nil {
			timestamps = append(timestamps, *postCredits.Start)
		}
		if postCredits.End != nil {
			timestamps = append(timestamps, *postCredits.End)
		}
		if !postCredits.isSame(preview) {
			addPreview()
		}
	}

	// Remove "same" timestamp unknowns
	if len(timestamps) > 1 {
		for i := 1; i < len(timestamps)-1; i++ {
			firstIndex := i - 1
			first := timestamps[firstIndex]
			secondIndex := i
			second := timestamps[secondIndex]
			if math.Abs(first.At-second.At) < SAME_DIFF_THRESHOLD_SECONDS {
				var typeID = first.TypeID
				if *second.TypeID != internal.TIMESTAMP_ID_UNKNOWN {
					typeID = second.TypeID
				}
				combinedTimestamp := internal.ThirdPartyTimestamp{
					At:     math.Min(first.At, second.At),
					TypeID: typeID,
				}
				// Replace two timestamps with one
				firstHalf := timestamps[:firstIndex]
				secondHalf := append([]internal.ThirdPartyTimestamp{combinedTimestamp}, timestamps[secondIndex+1:]...)
				timestamps = append(firstHalf, secondHalf...)
			}
		}
	}

	// Set 0 if necessary
	if len(timestamps) > 0 && timestamps[0].At != 0 {
		zeroTimestamp := internal.ThirdPartyTimestamp{
			At:     0,
			TypeID: &internal.TIMESTAMP_ID_UNKNOWN,
		}
		timestamps = append([]internal.ThirdPartyTimestamp{zeroTimestamp}, timestamps...)
	}

	if len(timestamps) == 0 {
		return nil
	}
	episodeTitle := BetterVRVStandardizeEpisodeName(input.EpisodeTitle)
	return &internal.ThirdPartyEpisode{
		AbsoluteNumber: number,
		Name:           &episodeTitle,
		Number:         nil,
		Season:         season,
		Source:         utils.Ptr(internal.TimestampSourceBetterVrv),
		Timestamps:     utils.PtrSlice(timestamps),
		ShowID:         input.Series.ObjectID,
	}
}
