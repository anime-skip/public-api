package server

import (
	"os"
	"time"

	database "anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server/handlers"
	"anime-skip.com/backend/internal/utils"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/gin-gonic/gin"
)

var host, port, graphqlPath string
var enablePlayground, isDev bool

func init() {
	host = os.Getenv("HOST")
	port = utils.EnvString("PORT")
	graphqlPath = "/graphql"
	enablePlayground = utils.EnvBool("ENABLE_PLAYGROUND")
	isDev = utils.EnvBool("IS_DEV")
}

// Run the web server
func Run(orm *database.ORM, startedAt time.Time) {
	server := gin.New()
	if isDev {
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
	if enablePlayground {
		server.GET("/graphiql", handlers.PlaygroundHandler(graphqlPath))
	}
	server.POST(graphqlPath, handlers.GraphQLHandler(orm))

	log.D("Started web server in %s @ %s:%s", time.Since(startedAt), host, port)
	log.Panic(server.Run(host + ":" + port))
}
