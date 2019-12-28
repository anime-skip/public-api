package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/resolvers"
	"github.com/gin-gonic/gin"
)

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler(orm *database.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the internal/gql/generated.go file
	config := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}
	handler := handler.GraphQL(gql.NewExecutableSchema(config))
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines the handler to expose the GraphQL playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	handler := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
