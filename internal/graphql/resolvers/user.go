package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getUserById(ctx context.Context, id *uuid.UUID) (*graphql.User, error) {
	if id == nil {
		return nil, nil
	}
	internalUser, err := r.UserService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	user := mappers.ToGraphqlUser(internalUser)
	return &user, nil
}

// Mutations

// Queries

func (r *queryResolver) FindUser(ctx context.Context, userID *uuid.UUID) (*graphql.User, error) {
	return r.getUserById(ctx, userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*graphql.User, error) {
	user, err := r.UserService.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	gqlUser := mappers.ToGraphqlUser(user)
	return &gqlUser, nil
}

// Fields

func (r *userResolver) AdminOfShows(ctx context.Context, obj *graphql.User) ([]*graphql.ShowAdmin, error) {
	return r.getShowAdminsByUserId(ctx, obj.ID)
}
