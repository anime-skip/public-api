package repos

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func SavePreferences(db *gorm.DB, newPreferences models.InputPreferences, existingPreferences *entities.Preferences) (*entities.Preferences, error) {
	data := mappers.PreferencesInputModelToEntity(newPreferences, existingPreferences)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update preferences for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update preferences with id='%s'", data.ID)
	}
	return data, err
}

func FindPreferencesByUserID(db *gorm.DB, userID string) (*entities.Preferences, error) {
	preferences := &entities.Preferences{}
	err := db.Unscoped().Where("user_id = ?", userID).Find(preferences).Error
	if err != nil {
		log.V("Failed query: %v", err)
		return nil, fmt.Errorf("No preferences found with user_id='%s'", userID)
	}
	return preferences, nil
}
