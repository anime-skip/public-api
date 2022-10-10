package directives

import (
	context1 "context"
	"fmt"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/samber/lo"
)

var (
	DEV_ROLES      = []int64{internal.ROLE_DEV}
	ADMIN_ROLES    = []int64{internal.ROLE_DEV, internal.ROLE_ADMIN}
	REVIEWER_ROLES = []int64{internal.ROLE_DEV, internal.ROLE_ADMIN, internal.ROLE_REVIEWER}
)

func HasRole(ctx context1.Context, obj any, next graphql.Resolver, requiredRole internal.Role) (res any, err error) {
	ctx, err = authenticate(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	hasRole := false
	switch requiredRole {
	case internal.RoleDev:
		hasRole = lo.Contains(DEV_ROLES, auth.Role)
	case internal.RoleAdmin:
		hasRole = lo.Contains(ADMIN_ROLES, auth.Role)
	case internal.RoleReviewer:
		hasRole = lo.Contains(REVIEWER_ROLES, auth.Role)
	}

	if !hasRole {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: fmt.Sprintf("Forbidden - you don't have the required role to perform this action (required: %s, has: %d)", requiredRole, auth.Role),
			Op:      "hasRole",
		}
	}
	return next(ctx)
}
