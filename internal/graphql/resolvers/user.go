package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getUserById(ctx context.Context, id *uuid.UUID) (*graphql.User, error) {
	// Handle deletedBy case
	if id == nil {
		return nil, nil
	}

	// return user for created_by_user_id and updated_by_user_id
	user, err := r.UserService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	gqlUser := mappers.ToGraphqlUser(user)
	return &gqlUser, nil
}

// Mutations

// Queries

func (r *queryResolver) FindUser(ctx context.Context, userID *uuid.UUID) (*graphql.User, error) {
	return r.getUserById(ctx, userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*graphql.User, error) {
	panic("queryResolver.FindUserByUsername not implemented")
}

// Fields

func (r *userResolver) AdminOfShows(ctx context.Context, obj *graphql.User) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}