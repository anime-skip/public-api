package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ApplyGraphqlInputTemplateTimestamp(input internal.InputTemplateTimestamp, output *internal.TemplateTimestamp) {
	output.TemplateID = input.TemplateID
	output.TimestampID = input.TimestampID
}
