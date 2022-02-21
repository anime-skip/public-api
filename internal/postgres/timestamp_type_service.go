package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type timestampTypeService struct {
	db internal.Database
}

func NewTimestampTypeService(db internal.Database) internal.TimestampTypeService {
	return &timestampTypeService{db}
}

func (s *timestampTypeService) GetByID(ctx context.Context, id uuid.UUID) (internal.TimestampType, error) {
	return getTimestampTypeByID(ctx, s.db, id)
}

func (s *timestampTypeService) GetAll(ctx context.Context) ([]internal.TimestampType, error) {
	panic("timestampTypeService.GetAll not implemented")
}
