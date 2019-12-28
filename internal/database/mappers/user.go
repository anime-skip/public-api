package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// UserModelToEntity -
func UserModelToEntity(model *models.User) *entities.User {
	return nil
}

// UserEntityToModel -
func UserEntityToModel(entity *entities.User) *models.User {
	if entity == nil {
		return nil
	}
	return &models.User{
		ID:           entity.ID.String(),
		CreatedAt:    entity.CreatedAt,
		DeletedAt:    entity.DeletedAt,
		Username:     entity.Username,
		Email:        entity.Email,
		ProfileURL:   entity.ProfileURL,
		AdminOfShows: nil,
	}
}

// UserEntityToMyUserModel -
func UserEntityToMyUserModel(entity *entities.User) *models.MyUser {
	if entity == nil {
		return nil
	}
	return &models.MyUser{
		ID:            entity.ID.String(),
		CreatedAt:     entity.CreatedAt,
		DeletedAt:     entity.DeletedAt,
		Username:      entity.Username,
		Email:         entity.Email,
		ProfileURL:    entity.ProfileURL,
		AdminOfShows:  nil,
		EmailVerified: entity.EmailVerified,
		Role:          models.RoleAdmin, // TODO - entity.Role,
		Preferences:   nil,
	}
}
