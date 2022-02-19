package graphql

import (
	ctx "context"
	"errors"

	"anime-skip.com/timestamps-service/internal/context"
	"github.com/99designs/gqlgen/graphql"
)

func authenticated(ctx ctx.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if !context.GetAlreadyAuthenticated(ctx) {
		authenticator := context.GetAuthenticator(ctx)
		token := context.GetAuthToken(ctx)
		if token == "" {
			return nil, errors.New("Unauthorized: Authorization header must be 'Bearer <token>'")
		}
		details, err := authenticator.Authenticate(token)
		if err != nil {
			return nil, err
		}
		ctx = context.WithAuthenticationDetails(ctx, details)
		ctx = context.WithAlreadyAuthenticated(ctx, true)
	}
	return next(ctx)
}

func hasRole(ctx ctx.Context, obj interface{}, next graphql.Resolver, role Role) (interface{}, error) {
	panic("Not implemented")
}

func isShowAdmin(ctx ctx.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	panic("Not implemented")
}
