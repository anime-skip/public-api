package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/mappers"
)

// Helpers

// Mutations

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	internalInput := internal.TemplateTimestamp{}
	mappers.ApplyGraphqlInputTemplateTimestamp(templateTimestamp, &internalInput)

	created, err := r.TemplateTimestampService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTemplateTimestamp(created)
	return &result, nil
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	internalTemplateTimestamp := internal.TemplateTimestamp{}
	mappers.ApplyGraphqlInputTemplateTimestamp(templateTimestamp, &internalTemplateTimestamp)

	deleted, err := r.TemplateTimestampService.Delete(ctx, internalTemplateTimestamp)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTemplateTimestamp(deleted)
	return &result, nil
}

// Queries

// Fields

func (r *templateTimestampResolver) Template(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Template, error) {
	return r.getTemplateByID(ctx, obj.TemplateID)
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Timestamp, error) {
	return r.getTimestampByID(ctx, obj.TimestampID)
}
