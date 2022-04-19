package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlTemplateTimestamp(templateTimestamp internal.TemplateTimestamp) graphql.TemplateTimestamp {
	return graphql.TemplateTimestamp{
		TemplateID:  &templateTimestamp.TemplateID,
		TimestampID: &templateTimestamp.TimestampID,
	}
}

func ApplyGraphqlInputTemplateTimestamp(input graphql.InputTemplateTimestamp, output *internal.TemplateTimestamp) {
	output.TemplateID = *input.TemplateID
	output.TimestampID = *input.TimestampID
}
