package mappers

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
)

var animeSkipSource = models.TimestampSourceAnimeSkip

func EpisodeInputModelToEntity(inputModel models.InputEpisode, entity *entities.Episode) *entities.Episode {
	if entity == nil {
		return nil
	}

	entity.Name = inputModel.Name
	entity.Season = inputModel.Season
	entity.Number = inputModel.Number
	entity.AbsoluteNumber = inputModel.AbsoluteNumber

	return entity
}

func EpisodeEntityToModel(entity *entities.Episode) *models.Episode {
	var deletedByUserID *string
	if entity.DeletedByUserID != nil {
		str := entity.DeletedByUserID.String()
		deletedByUserID = &str
	}
	return &models.Episode{
		ID:              entity.ID.String(),
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),
		DeletedAt:       entity.DeletedAt,
		DeletedByUserID: deletedByUserID,

		Name:           entity.Name,
		Season:         entity.Season,
		Number:         entity.Number,
		AbsoluteNumber: entity.AbsoluteNumber,
		ShowID:         entity.ShowID.String(),
	}
}

// This does not map the timestamps, it relies on a custom resolver to do that
func EpisodeEntityToThirdPartyEpisodeModel(entity *entities.Episode) *models.ThirdPartyEpisode {
	id := entity.ID.String()
	return &models.ThirdPartyEpisode{
		ID:             &id,
		Name:           entity.Name,
		AbsoluteNumber: entity.AbsoluteNumber,
		Number:         entity.Number,
		Season:         entity.Season,
		Source:         &animeSkipSource,
		ShowID:         entity.ShowID.String(),
		Timestamps:     nil,
	}
}
