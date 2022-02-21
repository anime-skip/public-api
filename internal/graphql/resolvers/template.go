package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTemplateByID(ctx context.Context, id *uuid.UUID) (*graphql.Template, error) {
	if id == nil {
		return nil, nil
	}
	internalTemplate, err := r.TemplateService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	template := mappers.ToGraphqlTemplate(internalTemplate)
	return &template, nil
}

func (r *Resolver) getTemplateByEpisodeID(ctx context.Context, episodeID *uuid.UUID) (*graphql.Template, error) {
	internalTemplate, err := r.TemplateService.GetByEpisodeID(ctx, *episodeID)
	if err != nil {
		return nil, err
	}
	template := mappers.ToGraphqlTemplate(internalTemplate)
	return &template, nil
}

func (r *Resolver) getTemplatesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Template, error) {
	internalTemplates, err := r.TemplateService.GetByShowID(ctx, *showID)
	if err != nil {
		return nil, err
	}
	templates := mappers.ToGraphqlTemplatePointers(internalTemplates)
	return templates, nil
}

// Mutations

func (r *mutationResolver) CreateTemplate(ctx context.Context, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("mutationResolver.CreateTemplate not implemented")
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID *uuid.UUID, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("mutationResolver.UpdateTemplate not implemented")
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID *uuid.UUID) (*graphql.Template, error) {
	panic("mutationResolver.DeleteTemplate not implemented")
}

// Queries

func (r *queryResolver) FindTemplate(ctx context.Context, templateID *uuid.UUID) (*graphql.Template, error) {
	return r.getTemplateByID(ctx, templateID)
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Template, error) {
	return r.getTemplatesByShowID(ctx, showID)
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *uuid.UUID, showName *string, season *string) (*graphql.Template, error) {
	panic("queryResolver.FindTemplateByDetails not implemented")
}

// Fields

func (r *templateResolver) CreatedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *templateResolver) UpdatedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *templateResolver) DeletedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *templateResolver) Show(ctx context.Context, obj *graphql.Template) (*graphql.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *graphql.Template) (*graphql.Episode, error) {
	return r.getEpisodeByID(ctx, obj.SourceEpisodeID)
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *graphql.Template) ([]*graphql.Timestamp, error) {
	panic("templareResolver.Timestamps not implemented")
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *graphql.Template) ([]*uuid.UUID, error) {
	panic("templareResolver.TimestampIds not implemented")
}
