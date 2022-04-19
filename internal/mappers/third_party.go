package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func InternalShowToThirdPartyShow(show internal.Show) internal.ThirdPartyShow {
	return internal.ThirdPartyShow{
		Name:      show.Name,
		UpdatedAt: &show.UpdatedAt,
		CreatedAt: &show.CreatedAt,
	}
}

func InternalTimestampToThirdPartyTimestamp(timestamp internal.Timestamp) internal.ThirdPartyTimestamp {
	return internal.ThirdPartyTimestamp{
		ID:     &timestamp.ID,
		At:     timestamp.At,
		TypeID: timestamp.TypeID,
	}
}

func InternalEpisodeToThirdPartyEpisode(episode internal.Episode, show internal.ThirdPartyShow, timestamps []internal.ThirdPartyTimestamp) internal.ThirdPartyEpisode {
	var baseDuration float64
	if episode.BaseDuration != nil {
		baseDuration = *episode.BaseDuration
	}
	return internal.ThirdPartyEpisode{
		ID:             &episode.ID,
		Season:         episode.Season,
		Number:         episode.Number,
		AbsoluteNumber: episode.AbsoluteNumber,
		BaseDuration:   baseDuration, // TODO: Refactor and make it not a pointer so it can be assigned directly
		Name:           episode.Name,
		Source:         internal.TIMESTAMP_SOURCE_ANIME_SKIP,
		ShowID:         episode.ShowID.String(),
		Show:           show,
		Timestamps:     timestamps,
	}
}

func ToGraphqlThirdPartyEpisodePointers(entities []internal.ThirdPartyEpisode) []*graphql.ThirdPartyEpisode {
	result := []*graphql.ThirdPartyEpisode{}
	for _, episode := range entities {
		result = append(result, toGraphqlThirdPartyEpisodePointer(episode))
	}
	return result
}

func toGraphqlThirdPartyEpisodePointer(entity internal.ThirdPartyEpisode) *graphql.ThirdPartyEpisode {
	source := ToTimestampSourceEnum(entity.Source)
	return &graphql.ThirdPartyEpisode{
		ID:             entity.ID,
		Season:         entity.Season,
		Number:         entity.Number,
		AbsoluteNumber: entity.AbsoluteNumber,
		BaseDuration:   &entity.BaseDuration,
		Name:           entity.Name,
		Source:         &source,
		ShowID:         entity.ShowID,
		Show:           toGraphqlThirdPartyShowPointer(entity.Show),
		Timestamps:     toGraphqlThirdPartyTimestampPointers(entity.Timestamps),
	}
}

func toGraphqlThirdPartyShowPointer(entity internal.ThirdPartyShow) *graphql.ThirdPartyShow {
	return &graphql.ThirdPartyShow{
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func toGraphqlThirdPartyTimestampPointers(entities []internal.ThirdPartyTimestamp) []*graphql.ThirdPartyTimestamp {
	result := []*graphql.ThirdPartyTimestamp{}
	for _, entity := range entities {
		result = append(result, toGraphqlThirdPartyTimestampPointer(entity))
	}
	return result
}

func toGraphqlThirdPartyTimestampPointer(entity internal.ThirdPartyTimestamp) *graphql.ThirdPartyTimestamp {
	return &graphql.ThirdPartyTimestamp{
		ID:     entity.ID,
		At:     entity.At,
		TypeID: &entity.TypeID,
		// Type has a resolver
	}
}
