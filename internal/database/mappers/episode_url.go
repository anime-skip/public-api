package mappers

import (
	"strings"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/constants"
)

func _urlToSource(url string) int {
	if strings.Contains(url, "vrv") {
		return constants.EPISODE_SOURCE_VRV
	}
	if strings.Contains(url, "funimation") {
		return constants.EPISODE_SOURCE_FUNIMATION
	}
	return constants.EPISODE_SOURCE_UNKNOWN
}

func EpisodeURLInputModelToEntity(inputModel models.InputEpisodeURL, entity *entities.EpisodeURL) *entities.EpisodeURL {
	if entity == nil {
		return nil
	}

	entity.URL = inputModel.URL
	entity.Duration = inputModel.Duration
	entity.TimestampsOffset = inputModel.TimestampsOffset
	entity.Source = _urlToSource(inputModel.URL)

	return entity
}

func EpisodeURLEntityToModel(entity *entities.EpisodeURL) *models.EpisodeURL {
	return &models.EpisodeURL{
		URL:             entity.URL,
		CreatedAt:       entity.CreatedAt,
		CreatedByUserID: entity.CreatedByUserID.String(),
		UpdatedAt:       entity.UpdatedAt,
		UpdatedByUserID: entity.UpdatedByUserID.String(),

		Source:           EpisodeSourceIntToEnum(entity.Source),
		Duration:         entity.Duration,
		TimestampsOffset: entity.TimestampsOffset,
		EpisodeID:        entity.EpisodeID.String(),
	}
}
