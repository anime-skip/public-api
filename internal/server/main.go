package server

import (
	"fmt"
	"time"

	database "anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server/handlers"
	"anime-skip.com/backend/internal/utils"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/gin-gonic/gin"
)

const GRAPHQL_PATH = "/graphql"

// Run the web server
func Run(orm *database.ORM, startedAt time.Time) {
	server := gin.New()
	if utils.ENV.IS_DEV {
		server.Use(log.RequestLogger)
	}
	server.Use(gin.Recovery())

	// Middleware
	server.Use(headerMiddleware)
	server.Use(ginContextToContextMiddleware)
	server.Use(corsMiddleware)

	// REST endpoints
	server.GET("/status", handlers.Status())

	// GraphQL endpoints
	if utils.ENV.ENABLE_PLAYGROUND {
		server.GET("/graphiql", handlers.PlaygroundHandler(GRAPHQL_PATH))
	}
	server.POST(GRAPHQL_PATH, handlers.GraphQLHandler(orm))

	port := fmt.Sprintf(":%d", utils.ENV.PORT)
	log.I("Started web server in %s @ %s", time.Since(startedAt), port)
	log.Panic(server.Run(port))
}
