package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
