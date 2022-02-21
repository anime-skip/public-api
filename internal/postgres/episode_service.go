package postgres

import "anime-skip.com/timestamps-service/internal"

type episodeService struct {
	db internal.Database
}

func NewEpisodeService(db internal.Database) internal.EpisodeService {
	return &episodeService{db}
}
