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

func (s *episodeURLService) Create(ctx context.Context, newEpisodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	return insertEpisodeURL(ctx, s.db, newEpisodeURL)
}

func (s *episodeURLService) Update(ctx context.Context, newEpisodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	return updateEpisodeURL(ctx, s.db, newEpisodeURL)
}

func (s *episodeURLService) Delete(ctx context.Context, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	panic("episodeURLService.Delete not implemented")
}
