package server

import (
	"fmt"
	"time"

	database "anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server/handlers"
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/gin-gonic/gin"
)

const GRAPHQL_PATH = "/graphql"

// Run the web server
func Run(orm *database.ORM, startedAt time.Time) {
	server := gin.New()
	server.Use(gin.Recovery())

	// Middleware
	if env.LOG_LEVEL >= constants.LOG_LEVEL_VERBOSE {
		server.Use(loggerMiddleware)
	}
	server.Use(banIPMiddleware)
	server.Use(headerMiddleware)
	server.Use(ginContextToContextMiddleware)
	server.Use(corsMiddleware)

	// REST endpoints
	server.GET("/status", handlers.Status())

	// GraphQL endpoints
	if env.ENABLE_PLAYGROUND {
		server.GET("/graphiql", handlers.PlaygroundHandler(GRAPHQL_PATH))
	}
	server.POST(GRAPHQL_PATH, handlers.GraphQLHandler(orm))

	port := fmt.Sprintf(":%d", env.PORT)
	log.I("Started web server in %s @ %s", time.Since(startedAt), port)
	log.V("---")
	log.Panic(server.Run(port))
}
