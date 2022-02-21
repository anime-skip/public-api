package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"github.com/gofrs/uuid"
)

type episodeService struct {
	db internal.Database
}

func NewEpisodeService(db internal.Database) internal.EpisodeService {
	return &episodeService{db}
}

func (s *episodeService) GetRecentlyAdded(ctx context.Context, params internal.GetRecentlyAddedParams) ([]internal.Episode, error) {
	return getRecentlyAddedEpisodes(ctx, s.db, params)
}

func (s *episodeService) GetByID(ctx context.Context, id uuid.UUID) (internal.Episode, error) {
	return getEpisodeByID(ctx, s.db, id)
}

func (s *episodeService) GetByShowID(ctx context.Context, showID uuid.UUID) ([]internal.Episode, error) {
	return getEpisodesByShowID(ctx, s.db, showID)
}
