package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/aklinker1/anime-skip-backend/internal/gql"
	"github.com/aklinker1/anime-skip-backend/internal/gql/resolvers"
	"github.com/gin-gonic/gin"
)

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the internal/gql/generated.go file
	config := gql.Config{
		Resolvers: &resolvers.Resolver{},
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
