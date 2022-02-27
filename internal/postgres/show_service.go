package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type showService struct {
	db internal.Database
}

func NewShowService(db internal.Database) internal.ShowService {
	return &showService{db}
}

func (s *showService) GetByID(ctx context.Context, id uuid.UUID) (internal.Show, error) {
	return getShowByID(ctx, s.db, id)
}

func (s *showService) GetSeasonCount(ctx context.Context, id uuid.UUID) (int, error) {
	return getEpisodeSeasonCountByShowID(ctx, s.db, id)
}

func (s *showService) Create(ctx context.Context, newShow internal.Show) (internal.Show, error) {
	return insertShow(ctx, s.db, newShow)
}

func (s *showService) Update(ctx context.Context, newShow internal.Show) (internal.Show, error) {
	return updateShow(ctx, s.db, newShow)
}

func (s *showService) Delete(ctx context.Context, id uuid.UUID) (internal.Show, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.Show{}, err
	}
	defer tx.Rollback()

	existing, err := getShowByIDInTx(ctx, tx, id)
	if err != nil {
		return internal.Show{}, err
	}

	deleted, err := deleteCascadeShow(ctx, tx, existing)
	if err != nil {
		return internal.Show{}, err
	}
	tx.Commit()
	return deleted, nil
}
