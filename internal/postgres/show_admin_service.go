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

func (s *showAdminService) Create(ctx context.Context, newShowAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	return insertShowAdmin(ctx, s.db, newShowAdmin)
}

func (s *showAdminService) Update(ctx context.Context, newShowAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	return updateShowAdmin(ctx, s.db, newShowAdmin)
}

func (s *showAdminService) Delete(ctx context.Context, id uuid.UUID) (internal.ShowAdmin, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	defer tx.Rollback()

	existing, err := getShowAdminByIDInTx(ctx, tx, id)
	if err != nil {
		return internal.ShowAdmin{}, err
	}

	deleted, err := deleteCascadeShowAdmin(ctx, tx, existing)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	tx.Commit()
	return deleted, nil
}
