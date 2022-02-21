package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/log"
)

func ToTemplateTypeInt(templateType graphql.TemplateType) int {
	switch templateType {
	case graphql.TemplateTypeShow:
		return internal.TEMPLATE_TYPE_SHOW
	case graphql.TemplateTypeSeasons:
		return internal.TEMPLATE_TYPE_SEASONS
	}
	log.E("Invalid template type enum: %v", templateType)
	return -1
}

func ToTemplateTypeEnum(value int) graphql.TemplateType {
	switch value {
	case internal.TEMPLATE_TYPE_SHOW:
		return graphql.TemplateTypeShow
	case internal.TEMPLATE_TYPE_SEASONS:
		return graphql.TemplateTypeSeasons
	}
	log.E("Invalid template type int: %d", value)
	return graphql.TemplateTypeShow
}
