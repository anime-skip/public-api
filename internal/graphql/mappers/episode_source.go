package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToEpisodeSourceInt(source graphql.EpisodeSource) int {
	switch source {
	case graphql.EpisodeSourceFunimation:
		return internal.EPISODE_SOURCE_FUNIMATION
	case graphql.EpisodeSourceVrv:
		return internal.EPISODE_SOURCE_VRV
	}
	return internal.EPISODE_SOURCE_UNKNOWN
}

func ToEpisodeSourceEnum(value int) graphql.EpisodeSource {
	switch value {
	case internal.EPISODE_SOURCE_FUNIMATION:
		return graphql.EpisodeSourceFunimation
	case internal.EPISODE_SOURCE_VRV:
		return graphql.EpisodeSourceVrv
	}
	return graphql.EpisodeSourceUnknown
}
