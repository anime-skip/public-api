package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Helpers

// Mutations

func (r *mutationResolver) CreateShow(ctx context.Context, showInput graphql.InputShow, becomeAdmin bool) (*graphql.Show, error) {
	panic("mutationResolver.CreateShow not implemented")
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID *uuid.UUID, newShow graphql.InputShow) (*graphql.Show, error) {
	panic("mutationResolver.UpdateShow not implemented")
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	panic("mutationResolver.DeleteShow not implemented")
}

// Queries

func (r *queryResolver) FindShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	panic("queryResolver.FindShow not implemented")
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*graphql.Show, error) {
	panic("queryResolver.SearchShows not implemented")
}

// Fields

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
	panic("showResolver.Admins not implemented")
}

func (r *showResolver) Episodes(ctx context.Context, obj *graphql.Show) ([]*graphql.Episode, error) {
	panic("showResolver.Episodes not implemented")
}

func (r *showResolver) Templates(ctx context.Context, obj *graphql.Show) ([]*graphql.Template, error) {
	panic("showResolver.Templates not implemented")
}

func (r *showResolver) SeasonCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("showResolver.SeasonCount not implemented")
}

func (r *showResolver) EpisodeCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("showResolver.EpisodeCount not implemented")
}
