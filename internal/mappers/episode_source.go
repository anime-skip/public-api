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
