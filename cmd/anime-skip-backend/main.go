package main

import (
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/server"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

func main() {
	log.W("Starting server")
	now := time.Now()
	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm, now)
}
