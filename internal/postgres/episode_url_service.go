package postgres

import "anime-skip.com/timestamps-service/internal"

type episodeURLService struct {
	db internal.Database
}

func NewEpisodeURLService(db internal.Database) internal.EpisodeURLService {
	return &episodeURLService{db}
}
