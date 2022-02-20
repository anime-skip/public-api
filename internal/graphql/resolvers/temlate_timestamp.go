package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *templateTimestampResolver) Template(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Template, error) {
	panic("templateTimestampResolver.Template not implemented")
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Timestamp, error) {
	panic("templateTimestampResolver.Timestamp not implemented")
}
