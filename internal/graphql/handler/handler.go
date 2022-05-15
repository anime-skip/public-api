package handler

import (
	"context"
	"time"

	"github.com/vektah/gqlparser/v2/gqlerror"

	"anime-skip.com/public-api/internal"
	myerrors "anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/directives"
	"anime-skip.com/public-api/internal/graphql/resolvers"
	"anime-skip.com/public-api/internal/log"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphqlHandler(db internal.Database, services internal.Services, enableIntrospection bool) internal.GraphQLHandler {
	log.D("Building GraphQL Server...")
	config := graphql.Config{
		Resolvers: &resolvers.Resolver{
			Services: &services,
			DB:       db,
		},
		Directives: graphql.DirectiveRoot{
			Authenticated: directives.Authenticated,
			HasRole:       directives.HasRole,
			IsShowAdmin:   directives.IsShowAdmin,
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

		if err, ok := paniced.(*myerrors.PanicedError); ok {
			return gqlerror.Errorf(err.Error())
		}
		return gqlerror.Errorf("Internal server error")
	})

	if enableIntrospection {
		srv.Use(extension.Introspection{})
	}

	return internal.GraphQLHandler{
		Handler:             srv,
		EnableIntrospection: enableIntrospection,
	}
}
