package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type episodeURLService struct {
	db internal.Database
}

func NewEpisodeURLService(db internal.Database) internal.EpisodeURLService {
	return &episodeURLService{db}
}

func (s *episodeURLService) GetByURL(ctx context.Context, url string) (internal.EpisodeURL, error) {
	return getEpisodeURLByURL(ctx, s.db, url)
}

func (s *episodeURLService) GetByEpisodeId(ctx context.Context, episodeID uuid.UUID) ([]internal.EpisodeURL, error) {
	return getEpisodeURLsByEpisodeID(ctx, s.db, episodeID)
}
