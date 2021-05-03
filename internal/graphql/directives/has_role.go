package directives

import (
	"context"
	"fmt"
	"net/http"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/context_utils"
	"github.com/99designs/gqlgen/graphql"
)

// ! Test after switching to use ctx
func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
	if err := context_utils.AuthError(ctx); err != nil {
		return nil, err
	}

	requiredRole := mappers.RoleEnumToInt(role)
	actualRole, hasRole := context_utils.Role(ctx)

	if !hasRole || actualRole < requiredRole {
		return nil, fmt.Errorf(http.StatusText(403))
	}
	return next(ctx)
}
