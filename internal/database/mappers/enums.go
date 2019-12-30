package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/aklinker1/anime-skip-backend/internal/utils/constants"
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
