package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func CreateTemplate(db *gorm.DB, templateInput models.InputTemplate) (*entities.Template, error) {
	template := mappers.TemplateInputModelToEntity(templateInput, &entities.Template{})
	err := db.Model(&template).Create(template).Error
	if err != nil {
		log.E("Failed to create template with [%+v]: %v", templateInput, err)
		return nil, err
	}
	return template, nil
}

func UpdateTemplate(db *gorm.DB, newTemplate models.InputTemplate, existingTemplate *entities.Template) (*entities.Template, error) {
	data := mappers.TemplateInputModelToEntity(newTemplate, existingTemplate)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update template for [%+v]: %v", data, err)
		return nil, err
	}
	return data, nil
}

func DeleteTemplate(tx *gorm.DB, templateID string) (err error) {
	// Delete the template
	err = tx.Delete(entities.Template{}, "id = ?", templateID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.E("Failed to delete template for id='%s': %v", templateID, err)
		return err
	}

	// Delete the TemplateTimestamps for that template
	templateTimestamps, err := FindTemplateTimestampsByTemplateID(tx, templateID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	for _, templateTimestamp := range templateTimestamps {
		err = DeleteTemplateTimestamp(tx, templateID, templateTimestamp.TimestampID.String())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return nil
}

func FindTemplateByID(db *gorm.DB, id string) (*entities.Template, error) {
	template := &entities.Template{}
	err := db.Where("id = ?", id).Find(template).Error
	if err != nil {
		return nil, err
	}
	return template, nil
}

func FindTemplatesByShowID(db *gorm.DB, showID string) ([]*entities.Template, error) {
	templates := []*entities.Template{}
	err := db.Where("show_id = ?", showID).Find(&templates).Error
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func FindTemplateBySourceEpisodeID(db *gorm.DB, sourceEpisodeID string) (*entities.Template, error) {
	template := &entities.Template{}
	err := db.Where("source_episode_id = ?", sourceEpisodeID).Find(template).Error
	if err != nil {
		return nil, err
	}
	return template, nil
}
