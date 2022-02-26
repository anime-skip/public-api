package mappers

import (
	"strings"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func urlToEpisodeSourceInt(url string) int {
	if strings.Contains(url, "vrv") {
		return internal.EPISODE_SOURCE_VRV
	}
	if strings.Contains(url, "funimation") {
		return internal.EPISODE_SOURCE_FUNIMATION
	}
	if strings.Contains(url, "crunchyroll") {
		return internal.EPISODE_SOURCE_CRUNCHYROLL
	}
	return internal.EPISODE_SOURCE_UNKNOWN
}

func ToEpisodeSourceInt(source graphql.EpisodeSource) int {
	switch source {
	case graphql.EpisodeSourceFunimation:
		return internal.EPISODE_SOURCE_FUNIMATION
	case graphql.EpisodeSourceVrv:
		return internal.EPISODE_SOURCE_VRV
	case graphql.EpisodeSourceCrunchyroll:
		return internal.EPISODE_SOURCE_CRUNCHYROLL
	}
	return internal.EPISODE_SOURCE_UNKNOWN
}

func ToEpisodeSourceEnum(value int) graphql.EpisodeSource {
	switch value {
	case internal.EPISODE_SOURCE_FUNIMATION:
		return graphql.EpisodeSourceFunimation
	case internal.EPISODE_SOURCE_VRV:
		return graphql.EpisodeSourceVrv
	case internal.EPISODE_SOURCE_CRUNCHYROLL:
		return graphql.EpisodeSourceCrunchyroll
	}
	return graphql.EpisodeSourceUnknown
}
