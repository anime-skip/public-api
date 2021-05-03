package handlers

import (
	"net/http"
	"time"

	"anime-skip.com/backend/internal/database"
	gql "anime-skip.com/backend/internal/graphql"
	"anime-skip.com/backend/internal/graphql/directives"
	"anime-skip.com/backend/internal/graphql/resolvers"
	"anime-skip.com/backend/internal/utils/env"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler(orm *database.ORM) http.HandlerFunc {
	schema := gql.NewExecutableSchema(gql.Config{
		Resolvers: resolvers.ResolverWithORM(orm),
		Directives: gql.DirectiveRoot{
			Authorized:  directives.Authorized,
			HasRole:     directives.HasRole,
			IsShowAdmin: directives.IsShowAdmin,
		},
		// Complexity: ,
	})

	gqlHandler := newServer(schema)

	return gqlHandler.ServeHTTP
}

// Based off handler.NewDefaultServer
func newServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	// srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	if env.ENABLE_INTROSPECTION {
		srv.Use(extension.Introspection{})
	}
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}

// PlaygroundHandler defines the handler to expose the GraphQL playground
func PlaygroundHandler(graphqlPath string) http.HandlerFunc {
	return playground.Handler("Anime Skip Playground", graphqlPath).ServeHTTP
}
