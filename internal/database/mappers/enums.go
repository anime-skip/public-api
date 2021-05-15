package mappers

import (
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/log"
)

func RoleEnumToInt(role models.Role) int {
	switch role {
	case models.RoleDev:
		return constants.ROLE_DEV
	case models.RoleAdmin:
		return constants.ROLE_ADMIN
	case models.RoleUser:
		return constants.ROLE_USER
	}
	log.E("Invalid role enum: %v", role)
	return -1
}

func RoleIntToEnum(value int) models.Role {
	switch value {
	case constants.ROLE_DEV:
		return models.RoleDev
	case constants.ROLE_ADMIN:
		return models.RoleAdmin
	case constants.ROLE_USER:
		return models.RoleUser
	}
	log.E("Invalid role int: %d", value)
	return models.RoleUser
}

func EpisodeSourceEnumToInt(source models.EpisodeSource) int {
	switch source {
	case models.EpisodeSourceFunimation:
		return constants.EPISODE_SOURCE_FUNIMATION
	case models.EpisodeSourceVrv:
		return constants.EPISODE_SOURCE_VRV
	}
	return constants.EPISODE_SOURCE_UNKNOWN
}

func EpisodeSourceIntToEnum(value int) models.EpisodeSource {
	switch value {
	case constants.EPISODE_SOURCE_FUNIMATION:
		return models.EpisodeSourceFunimation
	case constants.EPISODE_SOURCE_VRV:
		return models.EpisodeSourceVrv
	}
	return models.EpisodeSourceUnknown
}

func TimestampSouceEnumToInt(value *models.TimestampSource) int {
	if value != nil {
		switch *value {
		case models.TimestampSourceBetterVrv:
			return constants.TIMESTAMP_SOURCE_BETTER_VRV
		}
	}
	return constants.TIMESTAMP_SOURCE_ANIME_SKIP
}

func TimestampSouceIntToEnum(value int) models.TimestampSource {
	switch value {
	case constants.TIMESTAMP_SOURCE_BETTER_VRV:
		return models.TimestampSourceBetterVrv
	}
	return models.TimestampSourceAnimeSkip
}

func TemplateTypeEnumToInt(templateType models.TemplateType) int {
	switch templateType {
	case models.TemplateTypeShow:
		return constants.TEMPLATE_TYPE_SHOW
	case models.TemplateTypeSeasons:
		return constants.TEMPLATE_TYPE_SEASONS
	}
	log.E("Invalid template type enum: %v", templateType)
	return -1
}

func TemplateTypeIntToEnum(value int) models.TemplateType {
	switch value {
	case constants.TEMPLATE_TYPE_SHOW:
		return models.TemplateTypeShow
	case constants.TEMPLATE_TYPE_SEASONS:
		return models.TemplateTypeSeasons
	}
	log.E("Invalid template type int: %d", value)
	return models.TemplateTypeShow
}
