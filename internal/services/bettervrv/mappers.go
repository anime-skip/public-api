package bettervrv

import (
	"fmt"
	"math"
	"strings"

	"anime-skip.com/backend/internal/graphql/models"
	. "anime-skip.com/backend/internal/services/bettervrv/models"
	"anime-skip.com/backend/internal/utils/constants"
)

type Section struct {
	Start *models.ThirdPartyTimestamp
	End   *models.ThirdPartyTimestamp
}

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
func (betterVRVService betterVRVServiceInterface) StandardizeEpisodeName(title string) (standardized string) {
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
		str := fmt.Sprintf("%v", *input.AmbiguosNumber)
		number = &str
	}
	var season *string
	if input.Season != nil {
		str := fmt.Sprintf("%v", *input.Season)
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
	epsiodeTitle := BetterVRV.StandardizeEpisodeName(input.EpisodeTitle)
	return &models.ThirdPartyEpisode{
		AbsoluteNumber: number,
		Name:           &epsiodeTitle,
		Number:         nil,
		Season:         season,
		Source:         timestampTypeBetterVRVPtr,
		Timestamps:     timestamps,
		ShowID:         input.Series.ObjectId,
	}
}

func MapBetterVRVShowToThirdPartyShow(input *BetterVRVShow) *models.ThirdPartyShow {
	return &models.ThirdPartyShow{
		Name:      input.Title,
		CreatedAt: &input.CreatedAt,
		UpdatedAt: &input.UpdatedAt,
	}
}
