package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

func ShowInputModelToEntity(inputModel models.InputShow, entity *entities.Show) *entities.Show {
	if entity == nil {
		return nil
	}

	entity.Name = inputModel.Name
	entity.OriginalName = inputModel.OriginalName
	entity.Website = inputModel.Website
	entity.Image = inputModel.Image

	return entity
}

func ShowEntityToModel(entity *entities.Show) *models.Show {
	var deletedByUserID *string
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.Show{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,

		Name:         entity.Name,
		OriginalName: entity.OriginalName,
		Website:      entity.Website,
		Image:        entity.Image,
	}
}
