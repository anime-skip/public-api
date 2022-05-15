package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

func ToTemplateTypeInt(templateType internal.TemplateType) int {
	switch templateType {
	case internal.TemplateTypeShow:
		return internal.TEMPLATE_TYPE_SHOW
	case internal.TemplateTypeSeasons:
		return internal.TEMPLATE_TYPE_SEASONS
	}
	log.E("Invalid template type enum: %v", templateType)
	return -1
}

func ToTemplateTypeEnum(value int) internal.TemplateType {
	switch value {
	case internal.TEMPLATE_TYPE_SHOW:
		return internal.TemplateTypeShow
	case internal.TEMPLATE_TYPE_SEASONS:
		return internal.TemplateTypeSeasons
	}
	log.E("Invalid template type int: %d", value)
	return internal.TemplateTypeShow
}
