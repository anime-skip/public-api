package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

// Helpers

// Mutations

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("mutationResolver.AddTimestampToTemplate not implemented")
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("mutationResolver.RemoveTimestampFromTemplate not implemented")
}

// Queries

// Fields

func (r *templateTimestampResolver) Template(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Template, error) {
	panic("templateTimestampResolver.Template not implemented")
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Timestamp, error) {
	panic("templateTimestampResolver.Timestamp not implemented")
}
