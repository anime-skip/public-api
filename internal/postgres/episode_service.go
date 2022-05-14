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

func (s *episodeService) GetRecentlyAdded(ctx context.Context, params internal.GetRecentlyAddedFilter) ([]internal.Episode, error) {
	return getRecentlyAddedEpisodes(ctx, s.db, params)
}

func (s *episodeService) GetByID(ctx context.Context, id uuid.UUID) (internal.Episode, error) {
	return getEpisodeByID(ctx, s.db, id)
}

func (s *episodeService) GetByShowID(ctx context.Context, showID uuid.UUID) ([]internal.Episode, error) {
	return getEpisodesByShowID(ctx, s.db, showID)
}

func (s *episodeService) Create(ctx context.Context, newEpisode internal.Episode) (internal.Episode, error) {
	return insertEpisode(ctx, s.db, newEpisode)
}

func (s *episodeService) Update(ctx context.Context, newEpisode internal.Episode) (internal.Episode, error) {
	return updateEpisode(ctx, s.db, newEpisode)
}

func (s *episodeService) Delete(ctx context.Context, id uuid.UUID) (internal.Episode, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.Episode{}, err
	}
	defer tx.Rollback()

	existing, err := getEpisodeByIDInTx(ctx, tx, id)
	if err != nil {
		return internal.Episode{}, err
	}

	deleted, err := deleteCascadeEpisode(ctx, tx, existing)
	if err != nil {
		return internal.Episode{}, err
	}
	tx.Commit()
	return deleted, nil
}
