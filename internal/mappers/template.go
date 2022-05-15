package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputTemplate(input internal.InputTemplate, output *internal.Template) {
	output.ShowID = input.ShowID
	output.Type = input.Type
	output.Seasons = input.Seasons
	output.SourceEpisodeID = input.SourceEpisodeID
}
