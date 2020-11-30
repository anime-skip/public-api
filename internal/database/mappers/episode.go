package mappers

import (
	"strings"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
)

var animeSkipSource = models.TimestampSourceAnimeSkip

func EpisodeInputModelToEntity(inputModel models.InputEpisode, entity *entities.Episode) *entities.Episode {
	if entity == nil {
		return nil
	}

	// Replace empty values with nil
	Name := inputModel.Name
	if Name != nil && strings.TrimSpace(*Name) == "" {
		Name = nil
	}
	Season := inputModel.Season
	if Season != nil && strings.TrimSpace(*Season) == "" {
		Season = nil
	}
	Number := inputModel.Number
	if Number != nil && strings.TrimSpace(*Number) == "" {
		Number = nil
	}
	AbsoluteNumber := inputModel.AbsoluteNumber
	if AbsoluteNumber != nil && strings.TrimSpace(*AbsoluteNumber) == "" {
		AbsoluteNumber = nil
	}

	entity.Name = Name
	entity.Season = Season
	entity.Number = Number
	entity.AbsoluteNumber = AbsoluteNumber
	entity.BaseDuration = inputModel.BaseDuration

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
		BaseDuration:   entity.BaseDuration,
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
		BaseDuration:   entity.BaseDuration,
		Source:         &animeSkipSource,
		ShowID:         entity.ShowID.String(),
		Timestamps:     nil,
	}
}
