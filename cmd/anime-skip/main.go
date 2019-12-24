package main

import (
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/pkg/server"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

func main() {
	orm, err := database.Factory()
	if err != nil {
		log.Panic(err)
	}
	server.Run(orm)
}
