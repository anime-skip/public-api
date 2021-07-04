package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateTemplateTimestamp(db *gorm.DB, templateTimestampInput models.InputTemplateTimestamp) (*entities.TemplateTimestamp, error) {
	templateTimestamp := mappers.TemplateTimestampInputModelToEntity(templateTimestampInput, &entities.TemplateTimestamp{})
	err := db.Model(&templateTimestamp).Create(templateTimestamp).Error
	if err != nil {
		log.E("Failed to create template timestamp with [%+v]: %v", templateTimestampInput, err)
		return nil, err
	}
	return templateTimestamp, nil
}

func DeleteTemplateTimestamp(tx *gorm.DB, templateID string, timestampID string) (err error) {
	// Delete the templateTimestamp
	err = tx.Delete(entities.TemplateTimestamp{}, "template_id = ? AND timestamp_id = ?", templateID, timestampID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.V("Failed to delete templateTimestamp for template_id='%s' and timestamp_id='%s': %v", templateID, timestampID, err)
		return err
	}

	return nil
}

func FindTemplateTimestampByIDs(db *gorm.DB, templateID string, timestampID string) (*entities.TemplateTimestamp, error) {
	templateTimestamp := &entities.TemplateTimestamp{}
	err := db.Where("template_id = ? AND timestamp_id = ?", templateID, timestampID).Find(templateTimestamp).Error
	if err != nil {
		log.V("No template timestamp found with template_id='%s' and timestamp_id='%s': %v", templateID, timestampID, err)
		return nil, err
	}
	return templateTimestamp, nil
}

func FindTemplateTimestampsByTemplateID(db *gorm.DB, templateID string) ([]*entities.TemplateTimestamp, error) {
	templateTimestamps := []*entities.TemplateTimestamp{}
	err := db.Where("template_id = ?", templateID).Find(&templateTimestamps).Error
	if err != nil {
		log.V("No templateTimestamps found with template_id='%s': %v", templateID, err)
		return nil, err
	}
	return templateTimestamps, nil
}

func FindTemplateTimestampByTimestampID(db *gorm.DB, timestampID string) (*entities.TemplateTimestamp, error) {
	templateTimestamp := &entities.TemplateTimestamp{}
	err := db.Where("timestamp_id = ?", timestampID).First(&templateTimestamp).Error
	if err != nil {
		log.V("No templateTimestamp was found with timestamp_id='%s': %v", timestampID, err)
		return nil, err
	}
	return templateTimestamp, nil
}
