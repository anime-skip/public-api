package directives

import (
	"context"

	"anime-skip.com/backend/internal/utils/context_utils"
	"github.com/99designs/gqlgen/graphql"
)

func Authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if err := context_utils.AuthError(ctx); err != nil {
		return nil, err
	}
	return next(ctx)
}
