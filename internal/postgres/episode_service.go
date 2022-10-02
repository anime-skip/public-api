package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"github.com/gofrs/uuid"
)

type episodeService struct {
	db internal.Database
}

func NewEpisodeService(db internal.Database) internal.EpisodeService {
	return &episodeService{db}
}

func (s *episodeService) Get(ctx context.Context, filter internal.EpisodesFilter) (internal.Episode, error) {
	return inTx(ctx, s.db, false, internal.ZeroEpisode, func(tx internal.Tx) (internal.Episode, error) {
		return findEpisode(ctx, tx, filter)
	})
}

func (s *episodeService) List(ctx context.Context, filter internal.EpisodesFilter) ([]internal.Episode, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.Episode, error) {
		return findEpisodes(ctx, tx, filter)
	})
}

func (s *episodeService) ListRecentlyAdded(ctx context.Context, filter internal.RecentlyAddedEpisodesFilter) ([]internal.Episode, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.Episode, error) {
		return findRecentlyAddedEpisodes(ctx, tx, filter)
	})
}

func (s *episodeService) Create(ctx context.Context, newEpisode internal.Episode, createdBy uuid.UUID) (internal.Episode, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisode, func(tx internal.Tx) (internal.Episode, error) {
		return createEpisode(ctx, tx, newEpisode, createdBy)
	})
}

func (s *episodeService) Update(ctx context.Context, newEpisode internal.Episode, updatedBy uuid.UUID) (internal.Episode, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisode, func(tx internal.Tx) (internal.Episode, error) {
		return updateEpisode(ctx, tx, newEpisode, updatedBy)
	})
}

func (s *episodeService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.Episode, error) {
	return inTx(ctx, s.db, true, internal.ZeroEpisode, func(tx internal.Tx) (internal.Episode, error) {
		existing, err := findEpisode(ctx, tx, internal.EpisodesFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroEpisode, err
		}
		return deleteCascadeEpisode(ctx, tx, existing, deletedBy)
	})
}

func (s *episodeService) Count(ctx context.Context) (int, error) {
	return inTx(ctx, s.db, false, 0, func(tx internal.Tx) (int, error) {
		return count(ctx, tx, "SELECT COUNT(*) FROM episodes WHERE deleted_at IS NULL")
	})
}
