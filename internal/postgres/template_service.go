package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"github.com/gofrs/uuid"
)

type templateService struct {
	db internal.Database
}

func NewTemplateService(db internal.Database) internal.TemplateService {
	return &templateService{db}
}

func (s *templateService) Get(ctx context.Context, filter internal.TemplatesFilter) (internal.Template, error) {
	return inTx(ctx, s.db, false, internal.ZeroTemplate, func(tx internal.Tx) (internal.Template, error) {
		return findTemplate(ctx, tx, filter)
	})
}

func (s *templateService) List(ctx context.Context, filter internal.TemplatesFilter) ([]internal.Template, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.Template, error) {
		return findTemplates(ctx, tx, filter)
	})
}

func (s *templateService) Create(ctx context.Context, newTemplate internal.Template, createdBy uuid.UUID) (internal.Template, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplate, func(tx internal.Tx) (internal.Template, error) {
		return createTemplate(ctx, tx, newTemplate, createdBy)
	})
}

func (s *templateService) Update(ctx context.Context, newTemplate internal.Template, updatedBy uuid.UUID) (internal.Template, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplate, func(tx internal.Tx) (internal.Template, error) {
		return updateTemplate(ctx, tx, newTemplate, updatedBy)
	})
}

func (s *templateService) Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (internal.Template, error) {
	return inTx(ctx, s.db, true, internal.ZeroTemplate, func(tx internal.Tx) (internal.Template, error) {
		existing, err := findTemplate(ctx, tx, internal.TemplatesFilter{
			ID: &id,
		})
		if err != nil {
			return internal.ZeroTemplate, err
		}
		return deleteCascadeTemplate(ctx, tx, existing, deletedBy)
	})
}
