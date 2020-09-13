package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/gofrs/uuid"
)

func TimestampInputModelToEntity(inputModel models.InputTimestamp, entity *entities.Timestamp) *entities.Timestamp {
	if entity == nil {
		return nil
	}

	entity.At = inputModel.At
	entity.TypeID = uuid.FromStringOrNil(inputModel.TypeID)
	entity.Source = TimestampSouceEnumToInt(inputModel.Source)

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
		Source:    TimestampSouceIntToEnum(entity.Source),
		TypeID:    entity.TypeID.String(),
		EpisodeID: entity.EpisodeID.String(),
	}
}

func TimestampModelToThirdPartyTimestamp(entity *models.Timestamp) *models.ThirdPartyTimestamp {
	id := entity.ID
	return &models.ThirdPartyTimestamp{
		ID:     &id,
		At:     entity.At,
		TypeID: entity.TypeID,
	}
}
