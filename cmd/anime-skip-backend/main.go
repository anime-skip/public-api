package main

import (
	"time"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/pkg/server"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

func main() {
	log.V("Starting server")
	now := time.Now()
	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm, now)
}
