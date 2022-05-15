package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type showAdminService struct {
	db internal.Database
}

func NewShowAdminService(db internal.Database) internal.ShowAdminService {
	return &showAdminService{db}
}

func (s *showAdminService) Get(ctx context.Context, filter internal.ShowAdminsFilter) (internal.ShowAdmin, error) {
	return inTx(ctx, s.db, false, internal.ZeroShowAdmin, func(tx internal.Tx) (internal.ShowAdmin, error) {
		return findShowAdmin(ctx, tx, filter)
	})
}

func (s *showAdminService) List(ctx context.Context, filter internal.ShowAdminsFilter) ([]internal.ShowAdmin, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.ShowAdmin, error) {
		return findShowAdmins(ctx, tx, filter)
	})
}

func (s *showAdminService) Create(ctx context.Context, newShowAdmin internal.ShowAdmin, createdBy uuid.UUID) (internal.ShowAdmin, error) {
	return inTx(ctx, s.db, true, internal.ZeroShowAdmin, func(tx internal.Tx) (internal.ShowAdmin, error) {
		return createShowAdmin(ctx, tx, newShowAdmin, createdBy)
	})
}

func (s *showAdminService) Update(ctx context.Context, newShowAdmin internal.ShowAdmin, updatedBy uuid.UUID) (internal.ShowAdmin, error) {
	return inTx(ctx, s.db, true, internal.ZeroShowAdmin, func(tx internal.Tx) (internal.ShowAdmin, error) {
		return updateShowAdmin(ctx, tx, newShowAdmin, updatedBy)
	})
}

func (s *showAdminService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.ShowAdmin, error) {
	return inTx(ctx, s.db, true, internal.ZeroShowAdmin, func(tx internal.Tx) (internal.ShowAdmin, error) {
		existing, err := findShowAdmin(ctx, tx, internal.ShowAdminsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroShowAdmin, err
		}
		return deleteCascadeShowAdmin(ctx, tx, existing, deletedBy)
	})
}
