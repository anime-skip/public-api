package directives

import (
	context1 "context"
	"fmt"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context1.Context, obj any, next graphql.Resolver, role internal.Role) (res any, err error) {
	ctx, err = authenticate(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	hasRole := false
	if role == internal.RoleAdmin {
		hasRole = auth.IsAdmin || auth.IsDev
	} else if role == internal.RoleDev {
		hasRole = auth.IsDev
	}

	if !hasRole {
		return nil, fmt.Errorf("Forbidden - you don't have the required role to perform this action (%s)", role)
	}
	return next(ctx)
}
