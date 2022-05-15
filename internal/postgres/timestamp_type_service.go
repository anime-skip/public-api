package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type timestampTypeService struct {
	db internal.Database
}

func NewTimestampTypeService(db internal.Database) internal.TimestampTypeService {
	return &timestampTypeService{db}
}

func (s *timestampTypeService) Get(ctx context.Context, filter internal.TimestampTypesFilter) (internal.TimestampType, error) {
	return inTx(ctx, s.db, false, internal.ZeroTimestampType, func(tx internal.Tx) (internal.TimestampType, error) {
		return findTimestampType(ctx, tx, filter)
	})
}

func (s *timestampTypeService) List(ctx context.Context, filter internal.TimestampTypesFilter) ([]internal.TimestampType, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.TimestampType, error) {
		return findTimestampTypes(ctx, tx, filter)
	})
}

func (s *timestampTypeService) Create(ctx context.Context, newTimestampType internal.TimestampType, createdBy uuid.UUID) (internal.TimestampType, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestampType, func(tx internal.Tx) (internal.TimestampType, error) {
		return createTimestampType(ctx, tx, newTimestampType, createdBy)
	})
}

func (s *timestampTypeService) Update(ctx context.Context, newTimestampType internal.TimestampType, updatedBy uuid.UUID) (internal.TimestampType, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestampType, func(tx internal.Tx) (internal.TimestampType, error) {
		return updateTimestampType(ctx, tx, newTimestampType, updatedBy)
	})
}

func (s *timestampTypeService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.TimestampType, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestampType, func(tx internal.Tx) (internal.TimestampType, error) {
		existing, err := findTimestampType(ctx, tx, internal.TimestampTypesFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroTimestampType, err
		}
		return deleteCascadeTimestampType(ctx, tx, existing, deletedBy)
	})
}
