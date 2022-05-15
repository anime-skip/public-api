package mappers

import (
	"fmt"

	"anime-skip.com/public-api/internal"
)

func ToRoleEnum(i int) internal.Role {
	switch i {
	case internal.ROLE_DEV:
		return internal.RoleDev
	case internal.ROLE_ADMIN:
		return internal.RoleAdmin
	case internal.ROLE_USER:
		return internal.RoleUser
	}
	panic(&internal.Error{
		Code:    internal.EINVALID,
		Message: fmt.Sprintf("Unknown role integer: %d", i),
		Op:      "ToColorThemeEnum",
	})
}

func ToRoleInt(role internal.Role) int {
	switch role {
	case internal.RoleDev:
		return internal.ROLE_DEV
	case internal.RoleAdmin:
		return internal.ROLE_ADMIN
	case internal.RoleUser:
		return internal.ROLE_USER
	}
	panic(&internal.Error{
		Code:    internal.EINVALID,
		Message: fmt.Sprintf("Unknown role enum: %s", role),
		Op:      "ToColorThemeInt",
	})
}
