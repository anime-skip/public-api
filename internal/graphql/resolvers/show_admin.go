package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getShowAdminsByShowId(ctx context.Context, showID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	internalAdmins, err := r.ShowAdminService.GetByShowID(ctx, *showID)
	if err != nil {
		return nil, err
	}
	admins := mappers.ToGraphqlShowAdminPointers(internalAdmins)
	return admins, nil
}

func (r *Resolver) getShowAdminsByUserId(ctx context.Context, userID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	internalAdmins, err := r.ShowAdminService.GetByUserID(ctx, *userID)
	if err != nil {
		return nil, err
	}
	admins := mappers.ToGraphqlShowAdminPointers(internalAdmins)
	return admins, nil
}

// Mutations

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput graphql.InputShowAdmin) (*graphql.ShowAdmin, error) {
	panic(errors.NewPanicedError("TODO - show admins are disabled"))
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	deleted, err := r.ShowAdminService.Delete(ctx, *showAdminID)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlShowAdmin(deleted)
	return &result, nil
}

// Queries

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	internalAdmin, err := r.ShowAdminService.GetByID(ctx, *showAdminID)
	if err != nil {
		return nil, err
	}
	admin := mappers.ToGraphqlShowAdmin(internalAdmin)
	return &admin, nil
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	return r.getShowAdminsByShowId(ctx, showID)
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	return r.getShowAdminsByUserId(ctx, userID)
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
	return r.getShowById(ctx, obj.ShowID)
}

func (r *showAdminResolver) User(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
