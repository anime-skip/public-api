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

func (r *showResolver) Admins(ctx context.Context, obj *graphql.Show) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *showResolver) Episodes(ctx context.Context, obj *graphql.Show) ([]*graphql.Episode, error) {
	panic("not implemented")
}

func (r *showResolver) Templates(ctx context.Context, obj *graphql.Show) ([]*graphql.Template, error) {
	panic("not implemented")
}

func (r *showResolver) SeasonCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("not implemented")
}

func (r *showResolver) EpisodeCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("not implemented")
}
