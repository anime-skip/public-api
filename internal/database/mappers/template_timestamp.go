package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func TemplateTimestampInputModelToEntity(inputModel models.InputTemplateTimestamp, entity *entities.TemplateTimestamp) *entities.TemplateTimestamp {
	if entity == nil {
		return nil
	}

	entity.TemplateID = uuid.FromStringOrNil(inputModel.TemplateID)
	entity.TimestampID = uuid.FromStringOrNil(inputModel.TimestampID)

	return entity
}

func TemplateTimestampEntityToModel(entity *entities.TemplateTimestamp) *models.TemplateTimestamp {
	return &models.TemplateTimestamp{
		TemplateID:  entity.TemplateID.String(),
		TimestampID: entity.TimestampID.String(),
	}
}
