package directives

import (
	ctx "context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

func authenticate(ctx ctx.Context) (ctx.Context, error) {
	if context.GetAlreadyAuthenticated(ctx) {
		return ctx, nil
	}

	services := context.GetServices(ctx)
	token := context.GetAuthToken(ctx)
	if token == "" {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: "Unauthorized: Authorization header must be 'Bearer <token>'",
			Op:      "authenticate",
		}
	}
	details, err := services.AuthService.ValidateAccessToken(ctx, token)
	if err != nil {
		return nil, err
	}
	ctx = context.WithAuthClaims(ctx, details)
	ctx = context.WithAlreadyAuthenticated(ctx, true)
	return ctx, nil
}

func Authenticated(ctx ctx.Context, obj any, next graphql.Resolver) (any, error) {
	ctx, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return next(ctx)
}
