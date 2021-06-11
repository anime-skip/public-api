package repos

import (
	"errors"

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
		return nil, err
	}
	return timestamp, nil
}

func UpdateTimestamp(db *gorm.DB, newTimestamp models.InputTimestamp, existingTimestamp *entities.Timestamp) (*entities.Timestamp, error) {
	data := mappers.TimestampInputModelToEntity(newTimestamp, existingTimestamp)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update timestamp for [%+v]: %v", data, err)
		return nil, err
	}
	return data, err
}

func DeleteTimestamp(tx *gorm.DB, timestampID string) error {
	// Delete the timestamp
	err := tx.Delete(entities.Timestamp{}, "id = ?", timestampID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete timestamp for id='%s': %v", timestampID, err)
		return err
	}
	return nil
}

func FindTimestampByID(db *gorm.DB, timestampID string) (*entities.Timestamp, error) {
	timestamp := &entities.Timestamp{}
	err := db.Where("id = ?", timestampID).Find(timestamp).Error
	if err != nil {
		log.V("No timestamp found with id='%s': %v", timestampID, err)
		return nil, err
	}
	return timestamp, nil
}

func FindTimestampsByEpisodeID(db *gorm.DB, showID string) ([]*entities.Timestamp, error) {
	timestamps := []*entities.Timestamp{}
	err := db.Where("episode_id = ?", showID).Order("at ASC").Find(&timestamps).Error
	if err != nil {
		log.V("No timestamps found with episode_id='%s': %v", showID, err)
		return nil, err
	}
	return timestamps, nil
}
