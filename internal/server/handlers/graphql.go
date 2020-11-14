package handlers

import (
	"anime-skip.com/backend/internal/database"
	gql "anime-skip.com/backend/internal/graphql"
	"anime-skip.com/backend/internal/graphql/directives"
	"anime-skip.com/backend/internal/graphql/resolvers"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler(orm *database.ORM) gin.HandlerFunc {
	schema := gql.NewExecutableSchema(gql.Config{
		Resolvers: resolvers.ResolverWithORM(orm),
		Directives: gql.DirectiveRoot{
			Authorized:  directives.Authorized,
			HasRole:     directives.HasRole,
			IsShowAdmin: directives.IsShowAdmin,
		},
	})

	gqlHandler := handler.New(schema)
	gqlHandler.AddTransport(transport.POST{})

	return gin.WrapH(gqlHandler)
}

// PlaygroundHandler defines the handler to expose the GraphQL playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	handler := playground.Handler("Anime Skip Playground", path)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request.WithContext(c))
	}
}
