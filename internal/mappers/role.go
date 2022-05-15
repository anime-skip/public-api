package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/errors"
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
	panic(errors.NewPanicedError("Unknown role integer: %d", i))
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
	panic(errors.NewPanicedError("Unknown role enum: %s", role))
}
