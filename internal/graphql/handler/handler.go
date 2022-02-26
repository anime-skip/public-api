package handler

import (
	"time"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/directives"
	"anime-skip.com/timestamps-service/internal/graphql/resolvers"
	"anime-skip.com/timestamps-service/internal/log"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewGraphqlHandler(db internal.Database, services internal.Services, enableIntrospection bool, enableShowAdminDirective bool) internal.GraphQLHandler {
	log.D("Building GraphQL Server...")
	isShowAdmin := directives.AllowShowAdmin
	if enableShowAdminDirective {
		log.I("Enabling the @isShowAdmin directive")
		isShowAdmin = directives.IsShowAdmin
	}

	config := graphql.Config{
		Resolvers: &resolvers.Resolver{
			Services: &services,
			DB:       db,
		},
		Directives: graphql.DirectiveRoot{
			Authenticated: directives.Authenticated,
			HasRole:       directives.HasRole,
			IsShowAdmin:   isShowAdmin,
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
