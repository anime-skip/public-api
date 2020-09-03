package services

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"

	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/constants"
)

// Types

type betterVRVInterface struct{}
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

var BetterVRV = betterVRVInterface{}

const baseURL = "https://parseapi.back4app.com"

const APP_ID_KEY = "x-parse-application-id"
const APP_ID_VALUE = "CfnxYFbrcy0Eh517CcjOAlrAOH9hfe7dpOqfMcJj"

const API_KEY_KEY = "X-Parse-REST-API-Key"
const API_KEY_VALUE = "qUx8aEYYk8m9wajFvXPzHwf0nrxSG0hF4ekycmV3"

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

func (betterVRV betterVRVInterface) fetchEpisodeByName() (*models.ThirdPartyEpisode, error) {
	req, err := createRequest("/classes/Timestamps", nil, nil)
	if err != nil {
		return nil, err
	}
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	bvrvEpisode := &BetterVRVEpisode{}
	err = json.Unmarshal(body, bvrvEpisode)
	if err != nil {
		return nil, err
	}

	return mapBetterVRVEpisodeToThirdPartyEpisode(bvrvEpisode), nil
}

// Mappers

func (firstSection Section) endsWith(secondSection Section) bool {
	if firstSection.End != nil && secondSection.End != nil {
		return math.Abs(firstSection.Start.At-secondSection.End.At) < 2
	}
	return false
}

func (firstSection Section) isBefore(secondSection Section) bool {
	if firstSection.End != nil && secondSection.Start != nil {
		return firstSection.End.At < secondSection.Start.At
	} else if firstSection.End != nil {
		return true
	}
	return false
}

func (firstSection Section) isAfter(secondSection Section) bool {
	if firstSection.Start != nil && secondSection.End != nil {
		return firstSection.Start.At > secondSection.End.At
	} else if secondSection.End != nil {
		return true
	}
	return false
}

func (firstSection Section) isSame(secondSection Section) bool {
	if firstSection.Start != nil && secondSection.Start != nil && math.Abs(firstSection.Start.At-secondSection.Start.At) < 2 {
		return true
	}
	if firstSection.End != nil && secondSection.End != nil && math.Abs(firstSection.End.At-secondSection.End.At) < 2 {
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
				var startTime float64 = 0
				if *sectionEnd > sectionDuration {
					startTime = *sectionEnd - sectionDuration
				}
				start = &models.ThirdPartyTimestamp{
					At:     startTime,
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

func MapBetterVRVEpisodeToThirdPartyEpisode(input *BetterVRVEpisode) *models.ThirdPartyEpisode {
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

	// Preview is more of a recap
	usedPreview := preview.isBefore(intro)
	if usedPreview {
		timestamps = append(timestamps, preview.Start, preview.End, intro.Start, intro.End)
	} else {
		timestamps = append(timestamps, intro.Start, intro.End)
	}
	timestamps = append(timestamps, outro.Start, outro.End)
	if postCredits.isSame(preview) {
		if postCredits.isBefore(preview) {
			timestamps = append(timestamps, postCredits.Start)
		} else {
			timestamps = append(timestamps, preview.Start)
		}
		if postCredits.isAfter(preview) {

		} else {

		}
	} else if postCredits.isBefore(preview) {
		timestamps = append(timestamps, postCredits.Start)
		if preview.Start != nil {
			timestamps = append(timestamps, postCredits.End, preview.Start, preview.End)
		}
	} else if preview.isBefore(postCredits) {
		timestamps = append(timestamps, preview.Start)
		if preview.Start != nil {
			timestamps = append(timestamps, preview.End, postCredits.Start, postCredits.End)
		}
	}

	// Set 0 if necessary

	return &models.ThirdPartyEpisode{
		AbsoluteNumber: number,
		Name:           &input.EpisodeTitle,
		Number:         nil,
		Season:         season,
		Timestamps:     timestamps,
	}
}
