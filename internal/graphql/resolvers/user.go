package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getUserById(ctx context.Context, id *uuid.UUID) (*internal.User, error) {
	if id == nil {
		return nil, nil
	}
	fullUser, err := r.UserService.Get(ctx, internal.UsersFilter{
		ID:             id,
		IncludeDeleted: true,
	})
	if err != nil {
		return nil, err
	}
	user := mappers.ToUser(fullUser)
	return &user, nil
}

// Mutations

// Queries

func (r *queryResolver) FindUser(ctx context.Context, userID *uuid.UUID) (*internal.User, error) {
	return r.getUserById(ctx, userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*internal.User, error) {
	fullUser, err := r.UserService.Get(ctx, internal.UsersFilter{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	user := mappers.ToUser(fullUser)
	return &user, nil
}

// Fields

func (r *userResolver) AdminOfShows(ctx context.Context, obj *internal.User) ([]*internal.ShowAdmin, error) {
	return r.getShowAdminsByUserId(ctx, obj.ID)
}
