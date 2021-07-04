package repos

import (
	"errors"

	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func FindAPIClientByID(db *gorm.DB, clientID string) (*entities.APIClient, error) {
	apiClient := &entities.APIClient{}
	err := db.Where("id = ?", clientID).Find(apiClient).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.E("FindAPIClientByID(%s) error: %v", clientID, err)
		return nil, err
	}
	return apiClient, nil
}
