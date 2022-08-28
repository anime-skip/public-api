package directives

import (
	ctx "context"

	"anime-skip.com/public-api/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

func OptionalAuthenticated(ctx ctx.Context, obj any, next graphql.Resolver) (any, error) {
	token := context.GetAuthToken(ctx)
	if token == "" {
		return next(ctx)
	} else {
		return Authenticated(ctx, obj, next)
	}
}
