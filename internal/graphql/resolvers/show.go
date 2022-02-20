package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *showResolver) CreatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
