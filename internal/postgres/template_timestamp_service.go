package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
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

func (s *templateTimestampService) Delete(ctx context.Context, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	panic("templateService.Delete not implemented")
}

func (s *templateTimestampService) GetByTimestampID(ctx context.Context, timestampID uuid.UUID) (internal.TemplateTimestamp, error) {
	return getTemplateTimestampByTimestampID(ctx, s.db, timestampID)
}

func (s *templateTimestampService) GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]internal.TemplateTimestamp, error) {
	return getTemplateTimestampsByTemplateID(ctx, s.db, templateID)
}
