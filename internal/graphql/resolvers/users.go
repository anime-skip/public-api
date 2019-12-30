package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func userByID(ctx context.Context, db *gorm.DB, userID string) (*models.User, error) {
	user, err := repos.FindUserByID(ctx, db, userID)
	return mappers.UserEntityToModel(user), err
}

func deletedUserByID(ctx context.Context, db *gorm.DB, userID *string) (*models.User, error) {
	if userID == nil {
		return nil, nil
	}
	user, err := repos.FindUserByID(ctx, db, *userID)
	return mappers.UserEntityToModel(user), err
}

// Query Resolvers

func (r *queryResolver) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	return userByID(ctx, r.DB(ctx), userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := repos.FindUserByUsername(ctx, r.DB(ctx), username)
	return mappers.UserEntityToModel(user), err
}

// Mutation Resolvers

// Field Resolvers

type userResolver struct{ *Resolver }

func (r *userResolver) AdminOfShows(ctx context.Context, obj *models.User) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(ctx, r.DB(ctx), obj.ID)
}
