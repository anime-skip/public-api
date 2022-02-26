package directives

import (
	ctx "context"
	"errors"

	"anime-skip.com/timestamps-service/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

func authenticate(ctx ctx.Context) (ctx.Context, error) {
	if context.GetAlreadyAuthenticated(ctx) {
		return ctx, nil
	}

	services := context.GetDirectiveServices(ctx)
	token := context.GetAuthToken(ctx)
	if token == "" {
		return nil, errors.New("Unauthorized: Authorization header must be 'Bearer <token>'")
	}
	details, err := services.AuthService.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}
	ctx = context.WithAuthClaims(ctx, details)
	ctx = context.WithAlreadyAuthenticated(ctx, true)
	return ctx, nil
}

func Authenticated(ctx ctx.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	ctx, err := authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return next(ctx)
}
