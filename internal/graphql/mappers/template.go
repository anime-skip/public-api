package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/utils"
)

func ToGraphqlTemplate(template internal.Template) graphql.Template {
	return graphql.Template{
		ID:              &template.ID,
		CreatedAt:       template.CreatedAt,
		CreatedByUserID: &template.CreatedByUserID,
		UpdatedAt:       template.UpdatedAt,
		UpdatedByUserID: &template.UpdatedByUserID,
		DeletedAt:       template.DeletedAt,
		DeletedByUserID: template.DeletedByUserID,

		ShowID:          &template.ShowID,
		Type:            ToTemplateTypeEnum(template.Type),
		Seasons:         utils.SliceOrNil(template.Seasons),
		SourceEpisodeID: &template.SourceEpisodeID,
	}
}

func toGraphqlTemplatePointer(template internal.Template) *graphql.Template {
	value := ToGraphqlTemplate(template)
	return &value
}

func ToGraphqlTemplatePointers(templates []internal.Template) []*graphql.Template {
	result := []*graphql.Template{}
	for _, template := range templates {
		result = append(result, toGraphqlTemplatePointer(template))
	}
	return result
}

func ApplyGraphqlInputTemplate(input graphql.InputTemplate, output *internal.Template) {
	output.ShowID = *input.ShowID
	output.Type = ToTemplateTypeInt(input.Type)
	output.Seasons = input.Seasons
	output.SourceEpisodeID = *input.SourceEpisodeID
}
