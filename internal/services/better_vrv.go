package services

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/constants"
)

// Types

type betterVRVServiceInterface struct{}
type BetterVRVEpisode struct {
	ID             string `json:"objectId"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	VRVID          string `json:"episodeId"`
	EpisodeTitle   string `json:"episodeTitle"`
	Season         *int   `json:"seasonNumber"`
	AmbiguosNumber *int   `json:"episodeNumber"`
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

type Section struct {
	Start *models.ThirdPartyTimestamp
	End   *models.ThirdPartyTimestamp
}

// API

var BetterVRV = betterVRVServiceInterface{}

const baseURL = "https://parseapi.back4app.com"

const APP_ID_KEY = "x-parse-application-id"
const API_KEY_KEY = "X-Parse-REST-API-Key"

var APP_ID_VALUE = utils.EnvString("BETTER_VRV_APP_ID")
var API_KEY_VALUE = utils.EnvString("BETTER_VRV_API_KEY")

var localCache map[string]*models.ThirdPartyEpisode
var UNKOWN_EPISODE = &models.ThirdPartyEpisode{}

func createRequest(endpoint string, query map[string]string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
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

	req.Header.Add(APP_ID_KEY, APP_ID_VALUE)
	req.Header.Add(API_KEY_KEY, API_KEY_VALUE)
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}

func (betterVRVService betterVRVServiceInterface) FetchEpisodeByName(episodeName string) (*models.ThirdPartyEpisode, error) {
	if cachedResult, ok := localCache[episodeName]; ok {
		return cachedResult, nil
	}

	remoteResult, err := betterVRVService.fetchRemoteEpisodeByName(episodeName)
	if err != nil || remoteResult == nil {
		return nil, err
	}
	localCache[episodeName] = remoteResult
	return remoteResult, nil
}

func (_ betterVRVServiceInterface) fetchRemoteEpisodeByName(episodeName string) (*models.ThirdPartyEpisode, error) {
	req, err := createRequest("/classes/Timestamps", nil, nil)
	if err != nil {
		return nil, err
	}
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	bvrvEpisode := &BetterVRVEpisode{}
	err = json.Unmarshal(body, bvrvEpisode)
	if err != nil {
		return nil, err
	}

	return MapBetterVRVEpisodeToThirdPartyEpisode(bvrvEpisode), nil
}

// Mappers

var timestampTypeBetterVRV = models.TimestampSourceBetterVrv
var timestampTypeBetterVRVPtr = &timestampTypeBetterVRV

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

func createSection(hasSection *bool, sectionStart *float64, sectionEnd *float64, sectionDuration float64, typeID string) Section {
	var start *models.ThirdPartyTimestamp
	var end *models.ThirdPartyTimestamp

	if hasSection != nil && *hasSection {
		if sectionStart != nil && sectionEnd != nil {
			// Have both
			start = &models.ThirdPartyTimestamp{
				At:     *sectionStart,
				TypeID: typeID,
			}
			end = &models.ThirdPartyTimestamp{
				At:     *sectionEnd,
				TypeID: constants.TIMESTAMP_ID_UNKNOWN,
			}
		} else if sectionStart != nil {
			// Only have start
			start = &models.ThirdPartyTimestamp{
				At:     *sectionStart,
				TypeID: typeID,
			}
			if sectionDuration > 0 {
				end = &models.ThirdPartyTimestamp{
					At:     *sectionStart + sectionDuration,
					TypeID: constants.TIMESTAMP_ID_UNKNOWN,
				}
			}
		} else if sectionEnd != nil {
			// Only have end
			end = &models.ThirdPartyTimestamp{
				At:     *sectionEnd,
				TypeID: constants.TIMESTAMP_ID_UNKNOWN,
			}

			if sectionDuration > 0 {
				start = &models.ThirdPartyTimestamp{
					At:     *sectionEnd - sectionDuration,
					TypeID: typeID,
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
func standardizeEpisodeTitle(title string) (standardized string) {
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

func MapBetterVRVEpisodeToThirdPartyEpisode(input *BetterVRVEpisode) *models.ThirdPartyEpisode {
	if input == nil {
		return nil
	}

	var number *string
	if input.AmbiguosNumber != nil {
		str := string(*input.AmbiguosNumber)
		number = &str
	}
	var season *string
	if input.Season != nil {
		str := string(*input.Season)
		season = &str
	}

	// Parse sections
	intro := createSection(input.HasIntro, input.IntroStart, input.IntroEnd, 90, constants.TIMESTAMP_ID_INTRO)
	outro := createSection(input.HasOutro, input.OutroStart, input.OutroEnd, 90, constants.TIMESTAMP_ID_CREDITS)
	postCredits := createSection(input.HasPostCredit, input.PostCreditStart, input.PostCreditEnd, 0, constants.TIMESTAMP_ID_CANON)
	preview := createSection(input.HasPreview, input.PreviewStart, input.PreviewEnd, 0, constants.TIMESTAMP_ID_PREVIEW)

	// Combine Sections
	timestamps := []*models.ThirdPartyTimestamp{}
	addPreview := func() {
		if preview.Start != nil {
			timestamps = append(timestamps, preview.Start)
		}
		if preview.End != nil {
			timestamps = append(timestamps, preview.End)
		}
	}

	// RECAP
	isRecap := preview.Start != nil && preview.Start.At == 0
	if isRecap {
		if preview.Start != nil {
			preview.Start.TypeID = constants.TIMESTAMP_ID_RECAP
		}
		addPreview()
	}

	// INTRO
	if intro.Start != nil && intro.End != nil {
		if intro.Start.At < 0 {
			intro.Start.At = 0
		}
		timestamps = append(timestamps, intro.Start, intro.End)
	}

	// OUTRO
	if outro.Start != nil && outro.End != nil {
		timestamps = append(timestamps, outro.Start, outro.End)
	}

	// PREVIEW/POST CREDIT CANON
	if !isRecap {
		if postCredits.Start != nil {
			timestamps = append(timestamps, postCredits.Start)
		}
		if postCredits.End != nil {
			timestamps = append(timestamps, postCredits.End)
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
				if second.TypeID != constants.TIMESTAMP_ID_UNKNOWN {
					typeID = second.TypeID
				}
				combinedTimestamp := &models.ThirdPartyTimestamp{
					At:     math.Min(first.At, second.At),
					TypeID: typeID,
				}
				// Replace two timestamps with one
				firstHalf := timestamps[:firstIndex]
				secondHalf := append([]*models.ThirdPartyTimestamp{combinedTimestamp}, timestamps[secondIndex+1:]...)
				timestamps = append(firstHalf, secondHalf...)
			}
		}
	}

	// Set 0 if necessary
	if len(timestamps) > 0 && timestamps[0].At != 0 {
		zeroTimestamp := &models.ThirdPartyTimestamp{
			At:     0,
			TypeID: constants.TIMESTAMP_ID_UNKNOWN,
		}
		timestamps = append([]*models.ThirdPartyTimestamp{zeroTimestamp}, timestamps...)
	}

	if len(timestamps) == 0 {
		return nil
	}
	epsiodeTitle := standardizeEpisodeTitle(input.EpisodeTitle)
	return &models.ThirdPartyEpisode{
		AbsoluteNumber: number,
		Name:           &epsiodeTitle,
		Number:         nil,
		Season:         season,
		Source:         timestampTypeBetterVRVPtr,
		Timestamps:     timestamps,
	}
}
