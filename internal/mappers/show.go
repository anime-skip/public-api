package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputShow(input internal.InputShow, output *internal.Show) {
	output.Name = input.Name
	output.OriginalName = input.OriginalName
	output.Website = input.Website
	output.Image = input.Image
}
