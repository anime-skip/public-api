package graphql

import (
	"time"

	"anime-skip.com/timestamps-service/internal"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphqlHandler(db internal.Database, enableIntrospection bool) internal.GraphQLHandler {
	println("Defining GraphQL Server...")
	config := Config{
		Resolvers: &Resolver{
			db: db,
		},
		Directives: DirectiveRoot{
			Authenticated: authenticated,
			HasRole:       hasRole,
			IsShowAdmin:   isShowAdmin,
		},
	}
	srv := handler.New(NewExecutableSchema(config))
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
