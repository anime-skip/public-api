package repos

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// FindUserByID finds a user by their ID and returns them
func FindUserByID(ctx context.Context, orm *database.ORM, userID string) (*models.User, error) {
	user := &entities.User{}
	err := findOne(orm, "User", user, "id = ?", userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToModel(user)
}

// FindUserByUsername finds a user by their username and returns them
func FindUserByUsername(ctx context.Context, orm *database.ORM, username string) (*models.User, error) {
	user := &entities.User{}
	err := findOne(orm, "User", user, "username = ?", username)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToModel(user)
}
