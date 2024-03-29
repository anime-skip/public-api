package handler

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/directives"
	"anime-skip.com/public-api/internal/graphql/resolvers"
	"anime-skip.com/public-api/internal/log"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphqlHandler(db internal.Database, services internal.Services, rateLimiter internal.RateLimiter, enableIntrospection bool) internal.GraphQLHandler {
	log.D("Building GraphQL Server...")
	config := graphql.Config{
		Resolvers: &resolvers.Resolver{
			Services: &services,
			DB:       db,
		},
		Directives: graphql.DirectiveRoot{
			Authenticated:         directives.Authenticated,
			OptionalAuthenticated: directives.OptionalAuthenticated,
			HasRole:               directives.HasRole,
			IsShowAdmin:           directives.IsShowAdmin,
		},
	}
	srv := handler.New(graphql.NewExecutableSchema(config))
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.POST{})

	srv.SetRecoverFunc(func(ctx context.Context, paniced any) error {
		// TODO notify bugsnag

		log.E("Recovered panic: %v\n%s", paniced, string(debug.Stack()))
		return gqlerror.Errorf(internal.ErrorMessage(paniced))
	})
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		if e, ok := err.(*gqlerror.Error); ok {
			unwrapped := e.Unwrap()
			if unwrapped != nil {
				err = unwrapped
			} else {
				return e
			}
		}
		log.E("Error: %v", err)
		return gqlerror.Errorf(internal.ErrorMessage(err))
	})
	if rateLimiter != nil {
		extension := rateLimiter.GqlMiddleware()
		if extension != nil {
			srv.Use(extension)
		}
	}

	if enableIntrospection {
		srv.Use(extension.Introspection{})
	}

	return internal.GraphQLHandler{
		Handler:             srv,
		EnableIntrospection: enableIntrospection,
	}
}
