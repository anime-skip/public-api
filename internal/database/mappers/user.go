package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
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
		ProfileURL:   entity.ProfileURL,
		AdminOfShows: nil,
	}
}

// UserEntityToAccountModel -
func UserEntityToAccountModel(entity *entities.User) *models.Account {
	if entity == nil {
		return nil
	}
	return &models.Account{
		ID:            entity.ID.String(),
		CreatedAt:     entity.CreatedAt,
		DeletedAt:     entity.DeletedAt,
		Username:      entity.Username,
		Email:         entity.Email,
		ProfileURL:    entity.ProfileURL,
		AdminOfShows:  nil,
		EmailVerified: entity.EmailVerified,
		Role:          RoleIntToEnum(entity.Role),
		Preferences:   nil,
	}
}
