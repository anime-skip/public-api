package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToGraphqlEpisodeURL(entity internal.EpisodeURL) graphql.EpisodeURL {
	return graphql.EpisodeURL{
		URL:             entity.URL,
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: &entity.CreatedByUserID,
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: &entity.UpdatedByUserID,

		Source:           ToEpisodeSourceEnum(entity.Source),
		Duration:         entity.Duration,
		TimestampsOffset: entity.TimestampsOffset,
		EpisodeID:        &entity.EpisodeID,
	}
}

func toGraphqlEpisodeURLPointer(timestamp internal.EpisodeURL) *graphql.EpisodeURL {
	value := ToGraphqlEpisodeURL(timestamp)
	return &value
}

func ToGraphqlEpisodeURLPointers(episodeURLs []internal.EpisodeURL) []*graphql.EpisodeURL {
	result := []*graphql.EpisodeURL{}
	for _, episodeURL := range episodeURLs {
		result = append(result, toGraphqlEpisodeURLPointer(episodeURL))
	}
	return result
}

func ApplyGraphqlInputEpisodeURL(input graphql.InputEpisodeURL, output *internal.EpisodeURL) {
	output.URL = input.URL
	output.Duration = input.Duration
	output.TimestampsOffset = input.TimestampsOffset
	output.Source = urlToEpisodeSourceInt(input.URL)
}
