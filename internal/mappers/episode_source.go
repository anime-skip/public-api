package mappers

import (
	"strings"

	"anime-skip.com/public-api/internal"
)

func urlToEpisodeSource(url string) internal.EpisodeSource {
	if strings.Contains(url, "vrv") {
		return internal.EpisodeSourceVrv
	}
	if strings.Contains(url, "funimation") {
		return internal.EpisodeSourceFunimation
	}
	if strings.Contains(url, "crunchyroll") {
		return internal.EpisodeSourceCrunchyroll
	}
	return internal.EpisodeSourceUnknown
}

func ToEpisodeSourceInt(source internal.EpisodeSource) int {
	switch source {
	case internal.EpisodeSourceFunimation:
		return internal.EPISODE_SOURCE_FUNIMATION
	case internal.EpisodeSourceVrv:
		return internal.EPISODE_SOURCE_VRV
	case internal.EpisodeSourceCrunchyroll:
		return internal.EPISODE_SOURCE_CRUNCHYROLL
	}
	return internal.EPISODE_SOURCE_UNKNOWN
}

func ToEpisodeSourceEnum(value int) internal.EpisodeSource {
	switch value {
	case internal.EPISODE_SOURCE_FUNIMATION:
		return internal.EpisodeSourceFunimation
	case internal.EPISODE_SOURCE_VRV:
		return internal.EpisodeSourceVrv
	case internal.EPISODE_SOURCE_CRUNCHYROLL:
		return internal.EpisodeSourceCrunchyroll
	}
	return internal.EpisodeSourceUnknown
}
