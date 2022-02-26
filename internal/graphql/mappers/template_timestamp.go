package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func TemplateTimestampEntityToModel(entity *internal.TemplateTimestamp) *graphql.TemplateTimestamp {
	return &graphql.TemplateTimestamp{
		TemplateID:  &entity.TemplateID,
		TimestampID: &entity.TimestampID,
	}
}

func ApplyGraphqlInputTemplateTimestamp(input graphql.InputTemplateTimestamp, output *internal.TemplateTimestamp) {
	output.TemplateID = *input.TemplateID
	output.TimestampID = *input.TimestampID
}
