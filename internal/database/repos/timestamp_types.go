package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateTimestampType(db *gorm.DB, timestampTypeInput models.InputTimestampType) (*entities.TimestampType, error) {
	timestampType := mappers.TimestampTypeInputModelToEntity(timestampTypeInput, &entities.TimestampType{})
	err := db.Model(&timestampType).Create(timestampType).Error
	if err != nil {
		log.E("Failed to create timestamp type with [%+v]: %v", timestampTypeInput, err)
		return nil, err
	}
	return timestampType, nil
}

func UpdateTimestampType(db *gorm.DB, newTimestampType models.InputTimestampType, existingTimestampType *entities.TimestampType) (*entities.TimestampType, error) {
	data := mappers.TimestampTypeInputModelToEntity(newTimestampType, existingTimestampType)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update timestamp type for [%+v]: %v", data, err)
		return nil, err
	}
	return data, err
}

func DeleteTimestampType(tx *gorm.DB, timestampTypeID string) error {
	// Delete the timestampType
	err := tx.Delete(entities.TimestampType{}, "id = ?", timestampTypeID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete timestamp type for id='%s': %v", timestampTypeID, err)
		return err
	}
	return nil
}

func FindTimestampTypeByID(db *gorm.DB, timestampTypeID string) (*entities.TimestampType, error) {
	timestampType := &entities.TimestampType{}
	err := db.Where("id = ?", timestampTypeID).Find(timestampType).Error
	if err != nil {
		log.V("No timestamp type found with id='%s': %v", timestampTypeID, err)
		return nil, err
	}
	return timestampType, nil
}

func FindAllTimestampTypes(db *gorm.DB) ([]*entities.TimestampType, error) {
	timestampTypes := []*entities.TimestampType{}
	err := db.Find(&timestampTypes).Error
	if err != nil {
		log.V("Failed loading timestamp types: %v", err)
		return nil, err
	}
	return timestampTypes, nil
}
