package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *timestampTypeResolver) CreatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampTypeResolver) UpdatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampTypeResolver) DeletedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
