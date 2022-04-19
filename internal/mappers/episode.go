package mappers

import (
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlEpisode(episode internal.Episode) graphql.Episode {
	return graphql.Episode{
		ID:              &episode.ID,
		CreatedAt:       episode.CreatedAt,
		CreatedByUserID: &episode.CreatedByUserID,
		UpdatedAt:       episode.UpdatedAt,
		UpdatedByUserID: &episode.UpdatedByUserID,
		DeletedAt:       episode.DeletedAt,
		DeletedByUserID: episode.DeletedByUserID,

		Name:           episode.Name,
		Season:         episode.Season,
		Number:         episode.Number,
		AbsoluteNumber: episode.AbsoluteNumber,
		BaseDuration:   episode.BaseDuration,
		ShowID:         &episode.ShowID,
	}
}

func toGraphqlEpisodePointer(timestamp internal.Episode) *graphql.Episode {
	value := ToGraphqlEpisode(timestamp)
	return &value
}

func ToGraphqlEpisodePointers(episodes []internal.Episode) []*graphql.Episode {
	result := []*graphql.Episode{}
	for _, episode := range episodes {
		result = append(result, toGraphqlEpisodePointer(episode))
	}
	return result
}

func ApplyGraphqlInputEpisode(input graphql.InputEpisode, output *internal.Episode) {
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
	output.BaseDuration = input.BaseDuration
}
