package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

// FindUserByID finds a user by their ID and returns them
func FindUserByID(ctx context.Context, orm *database.ORM, userID string) (*entities.User, error) {
	user := &entities.User{}
	err := orm.DB.Where("id = ?", userID).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with id='%s'", userID)
	}
	return user, nil
}

// FindUserByUsername finds a user by their username and returns them
func FindUserByUsername(ctx context.Context, orm *database.ORM, username string) (*entities.User, error) {
	user := &entities.User{}
	err := orm.DB.Where("username = ?", username).Find(user).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No user found with username='%s'", username)
	}
	return user, nil
}
