package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type showAdminService struct {
	db internal.Database
}

func NewShowAdminService(db internal.Database) internal.ShowAdminService {
	return &showAdminService{db}
}

func (s *showAdminService) GetByID(ctx context.Context, id uuid.UUID) (internal.ShowAdmin, error) {
	return getShowAdminByID(ctx, s.db, id)
}

func (s *showAdminService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]internal.ShowAdmin, error) {
	return getShowAdminsByUserID(ctx, s.db, userID)
}

func (s *showAdminService) GetByShowID(ctx context.Context, showID uuid.UUID) ([]internal.ShowAdmin, error) {
	return getShowAdminsByShowID(ctx, s.db, showID)
}
