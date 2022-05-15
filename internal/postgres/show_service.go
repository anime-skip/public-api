package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type showService struct {
	db internal.Database
}

func NewShowService(db internal.Database) internal.ShowService {
	return &showService{db}
}

func (s *showService) Get(ctx context.Context, filter internal.ShowsFilter) (internal.Show, error) {
	return inTx(ctx, s.db, false, internal.ZeroShow, func(tx internal.Tx) (internal.Show, error) {
		return findShow(ctx, tx, filter)
	})
}

func (s *showService) GetSeasonCount(ctx context.Context, id uuid.UUID) (int, error) {
	return inTx(ctx, s.db, false, 0, func(tx internal.Tx) (int, error) {
		return countShowSeasons(ctx, tx, id)
	})
}

func (s *showService) List(ctx context.Context, filter internal.ShowsFilter) ([]internal.Show, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.Show, error) {
		return findShows(ctx, tx, filter)
	})
}

func (s *showService) Create(ctx context.Context, newShow internal.Show, createdBy uuid.UUID) (internal.Show, error) {
	return inTx(ctx, s.db, true, internal.ZeroShow, func(tx internal.Tx) (internal.Show, error) {
		return createShow(ctx, tx, newShow, createdBy)
	})
}

func (s *showService) Update(ctx context.Context, newShow internal.Show, updatedBy uuid.UUID) (internal.Show, error) {
	return inTx(ctx, s.db, true, internal.ZeroShow, func(tx internal.Tx) (internal.Show, error) {
		return updateShow(ctx, tx, newShow, updatedBy)
	})
}

func (s *showService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.Show, error) {
	return inTx(ctx, s.db, true, internal.ZeroShow, func(tx internal.Tx) (internal.Show, error) {
		existing, err := findShow(ctx, tx, internal.ShowsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroShow, err
		}
		return deleteCascadeShow(ctx, tx, existing, deletedBy)
	})
}
