package server

import (
	database "github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/handlers"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/gin-gonic/gin"
	"time"
)

var host, port, graphqlPath string
var enablePlayground, isDev bool

func init() {
	host = utils.EnvString("HOST")
	port = utils.EnvString("PORT")
	graphqlPath = "/graphql"
	enablePlayground = utils.EnvBool("ENABLE_PLAYGROUND")
	isDev = utils.EnvBool("IS_DEV")
}

// Run the web server
func Run(orm *database.ORM, startedAt time.Time) {
	server := gin.New()
	if isDev {
		server.Use(customLogger)
	}
	server.Use(gin.Recovery())

	// Middleware
	server.Use(headerMiddleware)
	server.Use(ginContextToContextMiddleware)

	// REST endpoints
	server.GET("/status", handlers.Status())

	// GraphQL endpoints
	if enablePlayground {
		server.GET("/graphiql", handlers.PlaygroundHandler(graphqlPath))
	}
	server.POST(graphqlPath, handlers.GraphQLHandler(orm))

	log.D("Started web server in %s @ \x1b[4mhttp://%s:%s", time.Since(startedAt), host, port)
	log.Panic(server.Run(host + ":" + port))
}