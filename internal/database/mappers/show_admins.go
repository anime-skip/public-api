package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func ShowAdminInputModelToEntity(inputModel models.InputShowAdmin, entity *entities.ShowAdmin) *entities.ShowAdmin {
	if entity == nil {
		return nil
	}

	entity.ShowID = uuid.FromStringOrNil(inputModel.ShowID)
	entity.UserID = uuid.FromStringOrNil(inputModel.UserID)

	return entity
}

func ShowAdminEntityToModel(entity *entities.ShowAdmin) *models.ShowAdmin {
	var deletedByUserID *string = nil
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.ShowAdmin{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,
		ShowID:          entity.ShowID.String(),
		UserID:          entity.UserID.String(),
	}
}