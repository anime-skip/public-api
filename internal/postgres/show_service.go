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

func (s *showService) Create(ctx context.Context, newShow internal.Show) (internal.Show, error) {
	return insertShow(ctx, s.db, newShow)
}

func (s *showService) Update(ctx context.Context, newShow internal.Show) (internal.Show, error) {
	return updateShow(ctx, s.db, newShow)
}

func (s *showService) Delete(ctx context.Context, show internal.Show) (internal.Show, error) {
	panic("showService.Delete not implemented")
}
