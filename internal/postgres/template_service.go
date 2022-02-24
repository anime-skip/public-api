package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"github.com/gofrs/uuid"
)

type templateService struct {
	db internal.Database
}

func NewTemplateService(db internal.Database) internal.TemplateService {
	return &templateService{db}
}

func (s *templateService) GetByID(ctx context.Context, id uuid.UUID) (internal.Template, error) {
	return getTemplateByID(ctx, s.db, id)
}

func (s *templateService) GetByShowID(ctx context.Context, showID uuid.UUID) ([]internal.Template, error) {
	return getTemplatesByShowID(ctx, s.db, showID)
}

func (s *templateService) GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) (internal.Template, error) {
	panic("templateService.GetByEpisodeID not implemented")
}

func (s *templateService) Create(ctx context.Context, newTemplate internal.Template) (internal.Template, error) {
	return insertTemplate(ctx, s.db, newTemplate)
}

func (s *templateService) Update(ctx context.Context, newTemplate internal.Template) (internal.Template, error) {
	return updateTemplate(ctx, s.db, newTemplate)
}

func (s *templateService) Delete(ctx context.Context, template internal.Template) (internal.Template, error) {
	panic("templateService.Delete not implemented")
}
