package main

import (
	"time"

	"anime-skip.com/backend/internal/database"
	"anime-skip.com/backend/internal/server"
	log "anime-skip.com/backend/internal/utils/log"
)

func main() {
	log.I("Initializing server...")
	now := time.Now()
	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm, now)
}
