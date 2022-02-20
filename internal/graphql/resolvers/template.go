package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

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
	panic("templareResolver.Show not implemented")
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *graphql.Template) (*graphql.Episode, error) {
	panic("templareResolver.SourceEpisode not implemented")
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *graphql.Template) ([]*graphql.Timestamp, error) {
	panic("templareResolver.Timestamps not implemented")
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *graphql.Template) ([]*uuid.UUID, error) {
	panic("templareResolver.TimestampIds not implemented")
}
