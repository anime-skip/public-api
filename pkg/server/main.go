package server

import (
	"log"

	"github.com/aklinker1/anime-skip-backend/internal/handlers"
	"github.com/aklinker1/anime-skip-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var host, port string

func init() {
	host = utils.EnvString("HOST")
	port = utils.EnvString("PORT")
}

// Run the web server
func Run() {
	server := gin.Default()

	server.GET("/ping", handlers.Status())

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(server.Run(host + ":" + port))
}
