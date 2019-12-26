package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// UserModelToEntity -
func UserModelToEntity(model *models.User) (*entities.User, error) {
	return nil, nil
}

// UserEntityToModel -
func UserEntityToModel(entity *entities.User) (*models.User, error) {
	model := &models.User{
		ID:           entity.ID.String(),
		CreatedAt:    entity.CreatedAt,
		DeletedAt:    entity.DeletedAt,
		Username:     entity.Username,
		Email:        entity.Email,
		ProfileURL:   entity.ProfileURL,
		AdminOfShows: nil,
	}
	return model, nil
}
