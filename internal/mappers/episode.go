package mappers

import (
	"strings"

	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputEpisode(input internal.InputEpisode, output *internal.Episode) {
	// Replace empty values with nil
	Name := input.Name
	if Name != nil && strings.TrimSpace(*Name) == "" {
		Name = nil
	}
	Season := input.Season
	if Season != nil && strings.TrimSpace(*Season) == "" {
		Season = nil
	}
	Number := input.Number
	if Number != nil && strings.TrimSpace(*Number) == "" {
		Number = nil
	}
	AbsoluteNumber := input.AbsoluteNumber
	if AbsoluteNumber != nil && strings.TrimSpace(*AbsoluteNumber) == "" {
		AbsoluteNumber = nil
	}

	output.Name = Name
	output.Season = Season
	output.Number = Number
	output.AbsoluteNumber = AbsoluteNumber
	output.BaseDuration = &input.BaseDuration
}
