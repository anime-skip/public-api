package mappers

import (
	"fmt"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToRoleEnum(i int) graphql.Role {
	switch i {
	case internal.ROLE_DEV:
		return graphql.RoleDev
	case internal.ROLE_ADMIN:
		return graphql.RoleAdmin
	case internal.ROLE_USER:
		return graphql.RoleUser
	}
	panic(fmt.Errorf("Unknown role integer: %d", i))
}

func ToRoleInt(role graphql.Role) int {
	switch role {
	case graphql.RoleDev:
		return internal.ROLE_DEV
	case graphql.RoleAdmin:
		return internal.ROLE_ADMIN
	case graphql.RoleUser:
		return internal.ROLE_USER
	}
	panic(fmt.Errorf("Unknown role enum: %s", role))
}
