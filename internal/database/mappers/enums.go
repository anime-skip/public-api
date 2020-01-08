package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
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
