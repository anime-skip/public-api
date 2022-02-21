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
	rows, err := s.db.QueryxContext(ctx, "SELECT * FROM timestamp_types WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	types := []internal.TimestampType{}
	for rows.Next() {
		var timestampType internal.TimestampType
		err = rows.StructScan(&timestampType)
		if err != nil {
			return nil, err
		}
		types = append(types, timestampType)
	}
	return types, nil
}
