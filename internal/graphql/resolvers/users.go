package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func userByID(db *gorm.DB, userID string) (*models.User, error) {
	user, err := repos.FindUserByID(db, userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToModel(user), nil
}

func deletedUserByID(db *gorm.DB, userID *string) (*models.User, error) {
	if userID == nil {
		return nil, nil
	}
	user, err := repos.FindUserByID(db, *userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToModel(user), nil
}

// Query Resolvers

func (r *queryResolver) FindUser(ctx context.Context, userID string) (*models.User, error) {
	return userByID(r.DB(ctx), userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := repos.FindUserByUsername(r.DB(ctx), username)
	return mappers.UserEntityToModel(user), err
}

// Mutation Resolvers

// Field Resolvers

type userResolver struct{ *Resolver }

func (r *userResolver) AdminOfShows(ctx context.Context, obj *models.User) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(r.DB(ctx), obj.ID)
}
