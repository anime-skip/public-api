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

func (r *episodeResolver) Show(ctx context.Context, obj *graphql.Episode) (*graphql.Show, error) {
	panic("episodeResolver.Show not implemented")
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *graphql.Episode) ([]*graphql.Timestamp, error) {
	panic("episodeResolver.Timestamps not implemented")
}

func (r *episodeResolver) Urls(ctx context.Context, obj *graphql.Episode) ([]*graphql.EpisodeURL, error) {
	panic("episodeResolver.Urls not implemented")
}

func (r *episodeResolver) Template(ctx context.Context, obj *graphql.Episode) (*graphql.Template, error) {
	panic("episodeResolver.Template not implemented")
}
