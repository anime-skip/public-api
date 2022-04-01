package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
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
	internalInput := internal.Template{
		BaseEntity: internal.BaseEntity{
			ID: utils.RandomID(),
		},
	}
	mappers.ApplyGraphqlInputTemplate(newTemplate, &internalInput)

	created, err := r.TemplateService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTemplate(created)
	return &result, nil
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID *uuid.UUID, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	log.V("Updating: %v", templateID)
	existing, err := r.TemplateService.GetByID(ctx, *templateID)
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputTemplate(newTemplate, &existing)
	log.V("Updating to %+v", existing)
	created, err := r.TemplateService.Update(ctx, existing)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	result := mappers.ToGraphqlTemplate(created)
	return &result, nil
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID *uuid.UUID) (*graphql.Template, error) {
	deleted, err := r.TemplateService.Delete(ctx, *templateID)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTemplate(deleted)
	return &result, nil
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
	templateTimestamps, err := r.TemplateTimestampService.GetByTemplateID(ctx, *obj.ID)
	if err != nil {
		return nil, err
	}

	timestamps := []internal.Timestamp{}
	for _, templateTimestamp := range templateTimestamps {
		timestamp, err := r.TimestampService.GetByID(ctx, templateTimestamp.TimestampID)
		if err != nil {
			return nil, err
		}
		timestamps = append(timestamps, timestamp)
	}
	return mappers.ToGraphqlTimestampPointers(timestamps), nil
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *graphql.Template) ([]*uuid.UUID, error) {
	templateTimestamps, err := r.TemplateTimestampService.GetByTemplateID(ctx, *obj.ID)
	if err != nil {
		return nil, err
	}

	ids := []*uuid.UUID{}
	for _, timestamp := range templateTimestamps {
		ids = append(ids, &timestamp.TimestampID)
	}
	return ids, nil
}
