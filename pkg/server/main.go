package server

import (
	database "github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/handlers"
	"github.com/aklinker1/anime-skip-backend/pkg/utils"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
	"github.com/gin-gonic/gin"
)

var host, port, graphqlPath, playgroundPath string
var enablePlayground bool

func init() {
	host = utils.EnvString("HOST")
	port = utils.EnvString("PORT")
	graphqlPath = "/graphql"
	enablePlayground = utils.EnvBool("ENABLE_PLAYGROUND")
	playgroundPath = utils.EnvString("PLAYGROUND_PATH")
}

// Run the web server
func Run(orm *database.ORM) {
	server := gin.Default()

	// REST endpoints
	server.GET("/status", handlers.Status())

	// GraphQL endpoints
	if enablePlayground {
		server.GET(playgroundPath, handlers.PlaygroundHandler(graphqlPath))
	}
	server.POST(graphqlPath, handlers.GraphQLHandler())

	log.Panic(server.Run(host + ":" + port))
}
