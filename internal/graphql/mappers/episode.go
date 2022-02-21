package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
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

func ToGraphqlEpisodes(episodes []internal.Episode) []graphql.Episode {
	result := []graphql.Episode{}
	for _, episode := range episodes {
		result = append(result, ToGraphqlEpisode(episode))
	}
	return result
}

func ToGraphqlEpisodePointers(episodes []internal.Episode) []*graphql.Episode {
	result := []*graphql.Episode{}
	for _, episode := range ToGraphqlEpisodes(episodes) {
		result = append(result, &episode)
	}
	return result
}
