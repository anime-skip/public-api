package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
