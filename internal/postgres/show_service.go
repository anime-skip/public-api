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
