package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/mappers"
)

// Helpers

// Mutations

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp internal.InputTemplateTimestamp) (*internal.TemplateTimestamp, error) {
	internalInput := internal.TemplateTimestamp{}
	mappers.ApplyGraphqlInputTemplateTimestamp(templateTimestamp, &internalInput)

	created, err := r.TemplateTimestampService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp internal.InputTemplateTimestamp) (*internal.TemplateTimestamp, error) {
	deleted, err := r.TemplateTimestampService.Delete(ctx, templateTimestamp)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

// Fields

func (r *templateTimestampResolver) Template(ctx context.Context, obj *internal.TemplateTimestamp) (*internal.Template, error) {
	return r.getTemplateByID(ctx, obj.TemplateID)
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *internal.TemplateTimestamp) (*internal.Timestamp, error) {
	return r.getTimestampByID(ctx, obj.TimestampID)
}
