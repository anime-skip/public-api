package repos

import (
	"fmt"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateTimestampType(db *gorm.DB, timestampTypeInput models.InputTimestampType) (*entities.TimestampType, error) {
	timestampType := mappers.TimestampTypeInputModelToEntity(timestampTypeInput, &entities.TimestampType{})
	err := db.Model(&timestampType).Create(timestampType).Error
	if err != nil {
		log.E("Failed to create timestamp type with [%+v]: %v", timestampTypeInput, err)
		return nil, fmt.Errorf("Failed to create timestamp type: %v", err)
	}
	return timestampType, nil
}

func UpdateTimestampType(db *gorm.DB, newTimestampType models.InputTimestampType, existingTimestampType *entities.TimestampType) (*entities.TimestampType, error) {
	data := mappers.TimestampTypeInputModelToEntity(newTimestampType, existingTimestampType)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update timestamp type for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update timestamp type with id='%s'", data.ID)
	}
	return data, err
}

func DeleteTimestampType(db *gorm.DB, inTransaction bool, timestampTypeID string) (err error) {
	tx := utils.StartTransaction(db, inTransaction)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to delete timestamp type and it's admins: %+v", r)
			tx.Rollback()
		}
	}()

	// Delete the timestampType
	err = tx.Delete(entities.TimestampType{}, "id = ?", timestampTypeID).Error
	if err != nil {
		log.E("Failed to delete timestamp type for id='%s': %v", timestampTypeID, err)
		tx.Rollback()
		return fmt.Errorf("Failed to delete timestamp type with id='%s'", timestampTypeID)
	}

	utils.CommitTransaction(tx, inTransaction)
	return nil
}

func FindTimestampTypeByID(db *gorm.DB, timestampTypeID string) (*entities.TimestampType, error) {
	timestampType := &entities.TimestampType{}
	err := db.Unscoped().Where("id = ?", timestampTypeID).Find(timestampType).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No timestamp type found with id='%s'", timestampTypeID)
	}
	return timestampType, nil
}

func FindAllTimestampTypes(db *gorm.DB) ([]*entities.TimestampType, error) {
	timestampTypes := []*entities.TimestampType{}
	err := db.Find(&timestampTypes).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No timestamp types found")
	}
	return timestampTypes, nil
}
