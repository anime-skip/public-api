package validation

import (
	"errors"
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/config"
	"github.com/gofrs/uuid"
	"github.com/samber/lo"
)

var (
	INTRO_DURATION_MIN      = 50.0  // s
	INTRO_DURATION_MAX      = 160.0 // s
	INTRO_COUNT_MAX         = 1
	TRANSITION_DURATION_MIN = 0.0  // s
	TRANSITION_DURATION_MAX = 20.0 // s
	TITLE_CARD_COUNT_MAX    = 2    // Series title, episode title
	PREVIEW_COUNT_MAX       = 1
	CREDITS_COUNT_MAX       = 3
)

func EpisodeTimestamps(
	episode internal.Episode,
	timestamps []internal.Timestamp,
) (newTimestamps []internal.Timestamp, isValid bool, validationErrs []error) {
	validationErrs = []error{}
	isValid = true
	var validationErr error
	for _, rule := range rules {
		timestamps, validationErr = rule.check(episode, timestamps)
		if validationErr != nil {
			validationErrs = append(validationErrs, fmt.Errorf("%s: %v", rule.name, validationErr))
			isValid = false
		}
	}
	return timestamps, isValid, lo.Ternary(len(validationErrs) != 0, validationErrs, nil)
}

// rule is a function that accepts an episode and list of timestamps, and returns a new list and nil
// if there are no problems. Some rules can "fix" the timestamps. The "fixed" list is the returned
// list.
//
// rule returns the original list and a validation error message when the list cannot be fixed an needs manually purged.
type rule struct {
	name  string
	check func(episode internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error)
}

var noopCheck = func(episode internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error) {
	return timestamps, nil
}

var rules = []rule{
	// Order and timing validations
	{
		name: "Ensure there's a timestamp at 0",
		check: func(episode internal.Episode, timestamps []internal.Timestamp) (newTimestamps []internal.Timestamp, err error) {
			if len(timestamps) == 0 || timestamps[0].At <= 0 {
				return timestamps, nil
			}
			newTimestamps = make([]internal.Timestamp, len(timestamps)+1)
			newTimestamps[0] = internal.Timestamp{
				At:        0,
				Source:    internal.TimestampSourceAnimeSkip,
				TypeID:    lo.ToPtr(uuid.FromStringOrNil(config.TIMESTAMP_ID_UNKNOWN)),
				EpisodeID: episode.ID,
			}
			for i, ts := range timestamps {
				newTimestamps[i+1] = ts
			}
			return newTimestamps, nil
		},
	},
	{
		name: "Merge timestamps with the same type and time",
		check: func(_ internal.Episode, timestamps []internal.Timestamp) (newTimestamps []internal.Timestamp, err error) {
			return lo.Reduce(timestamps, func(res []internal.Timestamp, ts internal.Timestamp, i int) []internal.Timestamp {
				prev, _ := lo.Last(res)
				if i == 0 || prev.At != ts.At || prev.TypeID.String() != ts.TypeID.String() {
					return append(res, ts)
				}
				return res
			}, []internal.Timestamp{}), nil
		},
	},
	{
		name: "Timestamps cannot be at the same time",
		check: func(_ internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error) {
			var prev internal.Timestamp
			duplicateMessages := []string{}
			for i, ts := range timestamps {
				if i != 0 && prev.At == ts.At {
					duplicateMessages = append(
						duplicateMessages,
						fmt.Sprintf("%s+%s@%.2f", prev.TypeID.String(), ts.TypeID.String(), prev.At),
					)
				}
				prev = ts
			}
			if len(duplicateMessages) > 0 {
				return timestamps, errors.New(strings.Join(duplicateMessages, ", "))
			}
			return timestamps, nil
		},
	},
	{
		name: "The same timestamp type should not be used twice in a row",
		check: func(_ internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error) {
			var prev internal.Timestamp
			duplicateTypes := []string{}
			for i, ts := range timestamps {
				if i != 0 && prev.TypeID.String() == ts.TypeID.String() {
					duplicateTypes = append(duplicateTypes, ts.TypeID.String())
				}
				prev = ts
			}
			if len(duplicateTypes) > 0 {
				return timestamps, errors.New(strings.Join(duplicateTypes, ", "))
			}
			return timestamps, nil
		},
	},
	// Everything is in order, there are no duplicates - start checking specific timestamp types
	maxTimestampTypeCountRule("intro", isIntro, INTRO_COUNT_MAX),
	timestampTypeDurationRule("intro", isIntro, INTRO_DURATION_MIN, INTRO_DURATION_MAX),
	timestampTypeDurationRule("transition", isTransition, TRANSITION_DURATION_MIN, TRANSITION_DURATION_MAX),
	maxTimestampTypeCountRule("title card", isTitleCard, TITLE_CARD_COUNT_MAX),
	maxTimestampTypeCountRule("preview", isPreview, PREVIEW_COUNT_MAX),
	maxTimestampTypeCountRule("credits", isCredits, CREDITS_COUNT_MAX),
}

func maxTimestampTypeCountRule(
	targetTypeName string,
	isType func(typeID *uuid.UUID) bool,
	maxCount int,
) rule {
	return rule{
		name: fmt.Sprintf("At most, there should be %d %s(s)", maxCount, targetTypeName),
		check: func(_ internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error) {
			instances := lo.Reduce(timestamps, func(sum int, ts internal.Timestamp, _ int) int {
				return lo.Ternary(isType(ts.TypeID), sum+1, sum)
			}, 0)
			if instances > maxCount {
				return timestamps, fmt.Errorf("%d %s(s) found", instances, targetTypeName)
			}
			return timestamps, nil
		},
	}
}

func timestampTypeDurationRule(
	targetTypeName string,
	isType func(typeID *uuid.UUID) bool,
	min float64,
	max float64,
) rule {
	return rule{
		name: fmt.Sprintf("If %ss exists, they should be between %.2f-%.2fs long", targetTypeName, min, max),
		check: func(episode internal.Episode, timestamps []internal.Timestamp) ([]internal.Timestamp, error) {
			ts, i, ok := lo.FindIndexOf(timestamps, func(ts internal.Timestamp) bool {
				return isType(ts.TypeID)
			})
			if !ok {
				return timestamps, nil
			}
			var nextAt float64
			if i+1 < len(timestamps) {
				nextAt = timestamps[i+1].At
			} else if episode.BaseDuration != nil {
				nextAt = *episode.BaseDuration
			} else {
				// We can't find the end, so give a value resulting in a valid duration
				nextAt = ts.At + (min+max)/2
			}

			duration := nextAt - ts.At
			if duration < min || duration > max {
				return timestamps, fmt.Errorf("%.2fs", duration)
			}

			return timestamps, nil
		},
	}
}

func isIntro(typeID *uuid.UUID) bool {
	if typeID == nil {
		return false
	}
	str := typeID.String()
	return str == config.TIMESTAMP_ID_INTRO || str == config.TIMESTAMP_ID_NEW_INTRO || str == config.TIMESTAMP_ID_MIXED_INTRO
}

func isCredits(typeID *uuid.UUID) bool {
	if typeID == nil {
		return false
	}
	str := typeID.String()
	return str == config.TIMESTAMP_ID_CREDITS || str == config.TIMESTAMP_ID_NEW_CREDITS || str == config.TIMESTAMP_ID_MIXED_CREDITS
}

func isTransition(typeID *uuid.UUID) bool {
	if typeID == nil {
		return false
	}
	return typeID.String() == config.TIMESTAMP_ID_TRANSITION
}

func isTitleCard(typeID *uuid.UUID) bool {
	if typeID == nil {
		return false
	}
	return typeID.String() == config.TIMESTAMP_ID_TITLE_CARD
}

func isPreview(typeID *uuid.UUID) bool {
	if typeID == nil {
		return false
	}
	return typeID.String() == config.TIMESTAMP_ID_PREVIEW
}
