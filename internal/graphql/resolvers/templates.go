package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func templateByID(db *gorm.DB, templateID string) (*models.Template, error) {
	template, err := repos.FindTemplateByID(db, templateID)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateEntityToModel(template), nil
}

func templateBySourceEpisodeID(db *gorm.DB, sourceEpisodeID string) (*models.Template, error) {
	template, err := repos.FindTemplateBySourceEpisodeID(db, sourceEpisodeID)
	if err != nil {
		return nil, err
	}
	return mappers.TemplateEntityToModel(template), nil
}

func templatesByShowID(db *gorm.DB, showID string) ([]*models.Template, error) {
	templates, err := repos.FindTemplatesByShowID(db, showID)
	if err != nil {
		return nil, err
	}
	templateModels := make([]*models.Template, len(templates))
	for index, episode := range templates {
		templateModels[index] = mappers.TemplateEntityToModel(episode)
	}
	return templateModels, nil
}

// Query Resolvers

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *string, showID *string, showName string, season *string) (*models.Template, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplate(ctx context.Context, templateID string) (*models.Template, error) {
	panic("not implemented")
}

// Mutation Resolvers

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID string, newTemplate models.InputTemplate) (*models.Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID string) (*models.Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp models.InputTemplateTimestamp) (*models.TemplateTimestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp models.InputTemplateTimestamp) (*models.TemplateTimestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTemplateTimestamps(ctx context.Context, templateID string, add []*models.InputTemplateTimestamp, remove []*models.InputTemplateTimestamp) ([]*models.TemplateTimestamp, error) {
	panic("not implemented")
}

// Field Resolvers

type templateResolver struct{ *Resolver }

func (r *templateResolver) CreatedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *templateResolver) UpdatedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *templateResolver) DeletedBy(ctx context.Context, obj *models.Template) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *templateResolver) Show(ctx context.Context, obj *models.Template) (*models.Show, error) {
	return showByID(r.DB(ctx), obj.ShowID)
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *models.Template) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), obj.SourceEpisodeID)
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *models.Template) ([]*models.Timestamp, error) {
	panic("not implemented")
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *models.Template) ([]string, error) {
	panic("not implemented")
}