package repos

import (
	"fmt"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

func CreateEpisodeURL(db *gorm.DB, episodeID string, episodeURLInput models.InputEpisodeURL) (*entities.EpisodeURL, error) {
	episodeURL := mappers.EpisodeURLInputModelToEntity(episodeURLInput, &entities.EpisodeURL{
		EpisodeID: uuid.FromStringOrNil(episodeID),
	})
	err := db.Create(&episodeURL).Error
	if err != nil {
		log.E("Failed to create episode url with [%+v]: %v", episodeURLInput, err)
		return nil, fmt.Errorf("Failed to create episode url: %v", err)
	}
	return episodeURL, nil
}

func UpdateEpisodeURL(db *gorm.DB, newEpisodeURL models.InputEpisodeURL, existingEpisodeURL *entities.EpisodeURL) (*entities.EpisodeURL, error) {
	data := mappers.EpisodeURLInputModelToEntity(newEpisodeURL, existingEpisodeURL)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update episode url for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update episode url with url='%s'", data.URL)
	}
	return data, err
}

func DeleteEpisodeURL(tx *gorm.DB, episodeURLID string) error {
	// Delete the episodeURL
	err := tx.Delete(entities.EpisodeURL{}, "url = ?", episodeURLID).Error
	if err != nil {
		log.E("Failed to delete episode url for url='%s': %v", episodeURLID, err)
		return fmt.Errorf("Failed to delete episode url with url='%s'", episodeURLID)
	}

	return nil
}

func FindEpisodeURLByURL(db *gorm.DB, url string) (*entities.EpisodeURL, error) {
	episodeURL := &entities.EpisodeURL{}
	err := db.Unscoped().Where("url = ?", url).Find(episodeURL).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No episode url found with url='%s'", url)
	}
	return episodeURL, nil
}

func FindEpisodeURLsByEpisodeID(db *gorm.DB, showID string) ([]*entities.EpisodeURL, error) {
	episodeURLs := []*entities.EpisodeURL{}
	err := db.Where("episode_id = ?", showID).Find(&episodeURLs).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No episode urls found with episode_id='%s'", showID)
	}
	return episodeURLs, nil
}
