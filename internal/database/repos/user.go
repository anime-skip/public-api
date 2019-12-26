package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// FindUserByID finds a user by their ID and returns them
func FindUserByID(ctx context.Context, orm *database.ORM, userID string) (*models.User, error) {
	user := &entities.User{}
	err := orm.DB.Where("id = ?", userID).Find(user).Error
	if err != nil {
		return nil, fmt.Errorf("No user found with id='%s'", userID)
	}
	return mappers.UserEntityToModel(user)
}

// FindUserByIDPtr finds a user by their ID and returns them. It will return nil if the pointer is nil
func FindUserByIDPtr(ctx context.Context, orm *database.ORM, userID *string) (*models.User, error) {
	if userID == nil {
		return nil, nil
	}
	return FindUserByID(ctx, orm, *userID)
}

// FindUserByUsername finds a user by their username and returns them
func FindUserByUsername(ctx context.Context, orm *database.ORM, username string) (*models.User, error) {
	user := &entities.User{}
	err := orm.DB.Where("username = ?", username).Find(user).Error
	if err != nil {
		return nil, fmt.Errorf("No user found with username='%s'", username)
	}
	return mappers.UserEntityToModel(user)
}