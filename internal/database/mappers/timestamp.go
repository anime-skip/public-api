package mappers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func TimestampInputModelToEntity(inputModel models.InputTimestamp, entity *entities.Timestamp) *entities.Timestamp {
	if entity == nil {
		return nil
	}

	entity.At = inputModel.At
	entity.TypeID = uuid.FromStringOrNil(inputModel.TypeID)

	return entity
}

func TimestampEntityToModel(entity *entities.Timestamp) *models.Timestamp {
	var deletedByUserID *string
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.Timestamp{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,

		At:        entity.At,
		TypeID:    entity.TypeID.String(),
		EpisodeID: entity.EpisodeID.String(),
	}
}
