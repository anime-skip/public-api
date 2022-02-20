package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.Episode, error) {
	panic("not implemented")
}
