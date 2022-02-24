package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

type timestampService struct {
	db internal.Database
}

func NewTimestampService(db internal.Database) internal.TimestampService {
	return &timestampService{db}
}

func (s *timestampService) GetByID(ctx context.Context, id uuid.UUID) (internal.Timestamp, error) {
	return getTimestampByID(ctx, s.db, id)
}

func (s *timestampService) GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) ([]internal.Timestamp, error) {
	return getTimestampsByEpisodeID(ctx, s.db, episodeID)
}

func (s *timestampService) Create(ctx context.Context, newTimestamp internal.Timestamp) (internal.Timestamp, error) {
	return insertTimestamp(ctx, s.db, newTimestamp)
}

func (s *timestampService) Update(ctx context.Context, newTimestamp internal.Timestamp) (internal.Timestamp, error) {
	return updateTimestamp(ctx, s.db, newTimestamp)
}

func (s *timestampService) Delete(ctx context.Context, timestamp internal.Timestamp) (internal.Timestamp, error) {
	panic("timestampService.Delete not implemented")
}
