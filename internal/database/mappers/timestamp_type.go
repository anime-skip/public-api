package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
)

func TimestampTypeInputModelToEntity(inputModel models.InputTimestampType, entity *entities.TimestampType) *entities.TimestampType {
	if entity == nil {
		return nil
	}

	entity.Name = inputModel.Name
	entity.Description = inputModel.Description

	return entity
}

func TimestampTypeEntityToModel(entity *entities.TimestampType) *models.TimestampType {
	var deletedByUserID *string
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.TimestampType{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,

		Name:        entity.Name,
		Description: entity.Description,
	}
}
