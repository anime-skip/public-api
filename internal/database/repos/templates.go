package repos

import (
	"fmt"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func FindTemplateByID(db *gorm.DB, id string) (*entities.Template, error) {
	template := &entities.Template{}
	err := db.Unscoped().Where("id = ?", id).Find(template).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No template found with id='%s'", id)
	}
	return template, nil
}

func FindTemplatesByShowID(db *gorm.DB, showID string) ([]*entities.Template, error) {
	templates := []*entities.Template{}
	err := db.Where("show_id = ?", showID).Find(&templates).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No templates found with show_id='%s'", showID)
	}
	return templates, nil
}

func FindTemplateBySourceEpisodeID(db *gorm.DB, sourceEpisodeID string) (*entities.Template, error) {
	template := &entities.Template{}
	err := db.Unscoped().Where("source_episode_id = ?", sourceEpisodeID).Find(template).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No template found with source_episode_id='%s'", sourceEpisodeID)
	}
	return template, nil
}
