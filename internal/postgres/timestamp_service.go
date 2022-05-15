package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type timestampService struct {
	db internal.Database
}

func NewTimestampService(db internal.Database) internal.TimestampService {
	return &timestampService{db}
}

func (s *timestampService) Get(ctx context.Context, filter internal.TimestampsFilter) (internal.Timestamp, error) {
	return inTx(ctx, s.db, false, internal.ZeroTimestamp, func(tx internal.Tx) (internal.Timestamp, error) {
		return findTimestamp(ctx, tx, filter)
	})
}

func (s *timestampService) List(ctx context.Context, filter internal.TimestampsFilter) ([]internal.Timestamp, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.Timestamp, error) {
		return findTimestamps(ctx, tx, filter)
	})
}

func (s *timestampService) Create(ctx context.Context, newTimestamp internal.Timestamp, createdBy uuid.UUID) (internal.Timestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestamp, func(tx internal.Tx) (internal.Timestamp, error) {
		return createTimestamp(ctx, tx, newTimestamp, createdBy)
	})
}

func (s *timestampService) Update(ctx context.Context, newTimestamp internal.Timestamp, updatedBy uuid.UUID) (internal.Timestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestamp, func(tx internal.Tx) (internal.Timestamp, error) {
		return updateTimestamp(ctx, tx, newTimestamp, updatedBy)
	})
}

func (s *timestampService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.Timestamp, error) {
	return inTx(ctx, s.db, true, internal.ZeroTimestamp, func(tx internal.Tx) (internal.Timestamp, error) {
		existing, err := findTimestamp(ctx, tx, internal.TimestampsFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroTimestamp, err
		}
		return deleteCascadeTimestamp(ctx, tx, existing, deletedBy)
	})
}
