package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Helpers

// Mutations

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput graphql.InputShowAdmin) (*graphql.ShowAdmin, error) {
	panic("mutationResolver.CreateShowAdmin not implemented")
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	panic("mutationResolver.DeleteShowAdmin not implemented")
}

// Queries

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdmin not implemented")
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdminsByShowID not implemented")
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdminsByUserID not implemented")
}

// Fields

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
