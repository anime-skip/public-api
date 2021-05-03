package main

import (
	"fmt"
	"net/http"
	"time"

	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server"
	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
)

func main() {
	startedAt := time.Now()

	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}

	router := server.GetRoutes(orm)
	port := fmt.Sprintf(":%d", env.PORT)
	log.I("Started web server in %s @ %s", time.Since(startedAt), port)
	err = http.ListenAndServe(port, router)
	log.E("Server crashed: %v", err)
}
