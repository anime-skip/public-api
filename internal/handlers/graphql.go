package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/aklinker1/anime-skip-backend/internal/database"
	gql "github.com/aklinker1/anime-skip-backend/internal/graphql"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/directives"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/resolvers"
	"github.com/gin-gonic/gin"
)

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler(orm *database.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the internal/graphql/generated.go file
	config := gql.Config{
		Resolvers: resolvers.ResolverWithORM(orm),
	}

	// Set the directives
	config.Directives.Authorized = directives.Authorized
	config.Directives.HasRole = directives.HasRole
	config.Directives.IsShowAdmin = directives.IsShowAdmin

	// Apply and serve the schema
	handler := handler.GraphQL(gql.NewExecutableSchema(config))
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler defines the handler to expose the GraphQL playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	handler := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request.WithContext(c))
	}
}
