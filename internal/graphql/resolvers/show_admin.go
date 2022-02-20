package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.Show, error) {
	panic("showAdminResolver.Show not implemented")
}

func (r *showAdminResolver) User(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
