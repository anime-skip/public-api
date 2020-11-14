package handlers

import (
	"anime-skip.com/backend/internal/database"
	gql "anime-skip.com/backend/internal/graphql"
	"anime-skip.com/backend/internal/graphql/directives"
	"anime-skip.com/backend/internal/graphql/resolvers"
	"anime-skip.com/backend/internal/utils"
	"github.com/99designs/gqlgen-contrib/prometheus"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

var isIntrospectionEnabled bool

func init() {
	isIntrospectionEnabled = utils.EnvBool("ENABLE_INTROSPECTION")
}

// GraphQLHandler defines the handler for the generated GraphQL server
func GraphQLHandler(orm *database.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the internal/graphql/generated.go file
	config := gql.Config{
		Resolvers: resolvers.ResolverWithORM(orm),
	}

	gqlHandler.Use(prometheus.Tracer{})

	// Apply and serve the schema
	handler := handler.GraphQL(
		gql.NewExecutableSchema(config),
		handler.IntrospectionEnabled(isIntrospectionEnabled),
		// handler.ComplexityLimit(5),
	)
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
