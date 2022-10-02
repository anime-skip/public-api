package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type episodeURLService struct {
	db internal.Database
}

func NewEpisodeURLService(db internal.Database) internal.EpisodeURLService {
	return &episodeURLService{db}
}

func (s *episodeURLService) Get(ctx context.Context, filter internal.EpisodeURLsFilter) (internal.EpisodeURL, error) {
	return inTx(ctx, s.db, false, internal.ZeroEpisodeURL, func(tx internal.Tx) (internal.EpisodeURL, error) {
		return findEpisodeURL(ctx, tx, filter)
	})
}

func (s *episodeURLService) List(ctx context.Context, filter internal.EpisodeURLsFilter) ([]internal.EpisodeURL, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.EpisodeURL, error) {
		return findEpisodeURLs(ctx, tx, filter)
	})
}

func (s *episodeURLService) Create(ctx context.Context, newEpisodeURL internal.EpisodeURL, createdBy uuid.UUID) (internal.EpisodeURL, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisodeURL, func(tx internal.Tx) (internal.EpisodeURL, error) {
		return createEpisodeURL(ctx, tx, newEpisodeURL, createdBy)
	})
}

func (s *episodeURLService) Update(ctx context.Context, newEpisodeURL internal.EpisodeURL, updatedBy uuid.UUID) (internal.EpisodeURL, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisodeURL, func(tx internal.Tx) (internal.EpisodeURL, error) {
		return updateEpisodeURL(ctx, tx, newEpisodeURL, updatedBy)
	})
}

func (s *episodeURLService) Delete(ctx context.Context, url string) (internal.EpisodeURL, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisodeURL, func(tx internal.Tx) (internal.EpisodeURL, error) {
		existing, err := findEpisodeURL(ctx, tx, internal.EpisodeURLsFilter{
			URL: &url,
		})
		if err != nil {
			return internal.ZeroEpisodeURL, err
		}
		return deleteCascadeEpisodeURL(ctx, tx, existing)
	})
}

func (s *episodeURLService) Count(ctx context.Context) (int, error) {
	return inTx(ctx, s.db, false, 0, func(tx internal.Tx) (int, error) {
		return count(ctx, tx, "SELECT COUNT(*) FROM episode_urls")
	})
}
