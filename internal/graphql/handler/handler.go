package handler

import (
	"time"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/directives"
	"anime-skip.com/timestamps-service/internal/graphql/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphqlHandler(services internal.Services, enableIntrospection bool) internal.GraphQLHandler {
	println("Defining GraphQL Server...")
	config := graphql.Config{
		Resolvers: &resolvers.Resolver{
			Services: &services,
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

	if enableIntrospection {
		srv.Use(extension.Introspection{})
	}

	return internal.GraphQLHandler{
		Handler:             srv,
		EnableIntrospection: enableIntrospection,
	}
}
