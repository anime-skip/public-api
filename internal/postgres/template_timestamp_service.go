package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

type templateTimestampService struct {
	db internal.Database
}

func NewTemplateTimestampService(db internal.Database) internal.TemplateTimestampService {
	return &templateTimestampService{db}
}

func (s *templateTimestampService) Create(ctx context.Context, newTemplateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return insertTemplateTimestamp(ctx, s.db, newTemplateTimestamp)
}

func (s *templateTimestampService) Delete(ctx context.Context, existing internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	defer tx.Rollback()

	deleted, err := deleteCascadeTemplateTimestamp(ctx, tx, existing)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	tx.Commit()
	return deleted, nil
}

func (s *templateTimestampService) GetByTimestampID(ctx context.Context, timestampID uuid.UUID) (internal.TemplateTimestamp, error) {
	return getTemplateTimestampByTimestampID(ctx, s.db, timestampID)
}

func (s *templateTimestampService) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]internal.TemplateTimestamp, error) {
	return getTemplateTimestampsByTemplateID(ctx, s.db, templateID)
}
