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

func ToGraphqlEpisodeURLs(episodeURLs []internal.EpisodeURL) []graphql.EpisodeURL {
	result := []graphql.EpisodeURL{}
	for _, episodeURL := range episodeURLs {
		result = append(result, ToGraphqlEpisodeURL(episodeURL))
	}
	return result
}

func ToGraphqlEpisodeURLPointers(episodeURLs []internal.EpisodeURL) []*graphql.EpisodeURL {
	result := []*graphql.EpisodeURL{}
	for _, episodeURL := range ToGraphqlEpisodeURLs(episodeURLs) {
		result = append(result, &episodeURL)
	}
	return result
}
