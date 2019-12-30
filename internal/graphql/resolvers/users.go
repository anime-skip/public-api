package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Helpers

func userByID(ctx context.Context, orm *database.ORM, userID string) (*models.User, error) {
	user, err := repos.FindUserByID(ctx, orm, userID)
	return mappers.UserEntityToModel(user), err
}

func deletedUserByID(ctx context.Context, orm *database.ORM, userID *string) (*models.User, error) {
	if userID == nil {
		return nil, nil
	}
	user, err := repos.FindUserByID(ctx, orm, *userID)
	return mappers.UserEntityToModel(user), err
}

// Query Resolvers

func (r *queryResolver) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := repos.FindUserByUsername(ctx, r.ORM(ctx), username)
	return mappers.UserEntityToModel(user), err
}

// Mutation Resolvers

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	return false, fmt.Errorf("Not implemented")
}

// Field Resolvers

type userResolver struct{ *Resolver }

func (r *userResolver) AdminOfShows(ctx context.Context, obj *models.User) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(ctx, r.ORM(ctx), obj.ID)
}
