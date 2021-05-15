package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func TemplateInputModelToEntity(inputModel models.InputTemplate, entity *entities.Template) *entities.Template {
	if entity == nil {
		return nil
	}

	entity.ShowID = uuid.FromStringOrNil(inputModel.ShowID)
	entity.Type = TemplateTypeEnumToInt(inputModel.Type)
	entity.Seasons = inputModel.Seasons
	entity.SourceEpisodeID = uuid.FromStringOrNil(inputModel.SourceEpisodeID)

	return entity
}

func TemplateEntityToModel(entity *entities.Template) *models.Template {
	var deletedByUserID *string
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.Template{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,

		ShowID:          entity.ShowID.String(),
		Type:            TemplateTypeIntToEnum(entity.Type),
		Seasons:         entity.Seasons,
		SourceEpisodeID: entity.SourceEpisodeID.String(),
	}
}
