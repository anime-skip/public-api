package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

// FindUserByID finds a user by their ID and returns them
func FindUserByID(ctx context.Context, orm *database.ORM, userID string) (*models.User, error) {
	user := &entities.User{}
	err := orm.DB.Where("id = ?", userID).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
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
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with username='%s'", username)
	}
	return mappers.UserEntityToModel(user)
}

// FindMyUser will find your user based on the token provided, but for now it looks for the username argument
func FindMyUser(ctx context.Context, orm *database.ORM) (*models.MyUser, error) {
	username := "the_admin"
	user := &entities.User{}
	err := orm.DB.Where("username = ?", username).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with username='%s'", username)
	}
	return mappers.UserEntityToMyUserModel(user)
}
