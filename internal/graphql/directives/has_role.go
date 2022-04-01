package directives

import (
	ctx "context"
	"fmt"

	"anime-skip.com/public-api/internal/context"
	gql "anime-skip.com/public-api/internal/graphql"
	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx ctx.Context, obj interface{}, next graphql.Resolver, role gql.Role) (interface{}, error) {
	ctx, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	hasRole := false
	if role == gql.RoleAdmin {
		hasRole = auth.IsAdmin || auth.IsDev
	} else if role == gql.RoleDev {
		hasRole = auth.IsDev
	}

	if !hasRole {
		return nil, fmt.Errorf("Forbidden - you don't have the required role to perform this action (%s)", role)
	}
	return next(ctx)
}
