package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputEpisodeURL(input internal.InputEpisodeURL, output *internal.EpisodeURL) {
	output.URL = input.URL
	output.Duration = input.Duration
	output.TimestampsOffset = input.TimestampsOffset
	output.Source = urlToEpisodeSource(input.URL)
}
