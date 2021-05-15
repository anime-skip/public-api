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

func CreateTimestamp(db *gorm.DB, episodeID string, timestampInput models.InputTimestamp) (*entities.Timestamp, error) {
	timestamp := mappers.TimestampInputModelToEntity(timestampInput, &entities.Timestamp{
		EpisodeID: uuid.FromStringOrNil(episodeID),
	})
	err := db.Model(&timestamp).Create(timestamp).Error
	if err != nil {
		log.E("Failed to create timestamp with [%+v]: %v", timestampInput, err)
		return nil, fmt.Errorf("Failed to create timestamp: %v", err)
	}
	return timestamp, nil
}

func UpdateTimestamp(db *gorm.DB, newTimestamp models.InputTimestamp, existingTimestamp *entities.Timestamp) (*entities.Timestamp, error) {
	data := mappers.TimestampInputModelToEntity(newTimestamp, existingTimestamp)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update timestamp for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update timestamp with id='%s'", data.ID)
	}
	return data, err
}

func DeleteTimestamp(tx *gorm.DB, timestampID string) error {
	// Delete the timestamp
	err := tx.Delete(entities.Timestamp{}, "id = ?", timestampID).Error
	if err != nil {
		log.E("Failed to delete timestamp for id='%s': %v", timestampID, err)
		return fmt.Errorf("Failed to delete timestamp with id='%s'", timestampID)
	}
	return nil
}

func FindTimestampByID(db *gorm.DB, timestampID string) (*entities.Timestamp, error) {
	timestamp := &entities.Timestamp{}
	err := db.Unscoped().Where("id = ?", timestampID).Find(timestamp).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No timestamp found with id='%s'", timestampID)
	}
	return timestamp, nil
}

func FindTimestampsByEpisodeID(db *gorm.DB, showID string) ([]*entities.Timestamp, error) {
	timestamps := []*entities.Timestamp{}
	err := db.Where("episode_id = ?", showID).Order("at ASC").Find(&timestamps).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No timestamps found with episode_id='%s'", showID)
	}
	return timestamps, nil
}
