package repos

import (
	"anime-skip.com/backend/internal/database/entities"
	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
)

func SavePreferences(db *gorm.DB, newPreferences models.InputPreferences, existingPreferences *entities.Preferences) (*entities.Preferences, error) {
	data := mappers.PreferencesInputModelToEntity(newPreferences, existingPreferences)
	err := db.Save(data).Error
	if err != nil {
		log.E("Failed to update preferences for [%+v]: %v", data, err)
		return nil, err
	}
	return data, nil
}

func FindPreferencesByUserID(db *gorm.DB, userID string) (*entities.Preferences, error) {
	preferences := &entities.Preferences{}
	err := db.Where("user_id = ?", userID).Find(preferences).Error
	if err != nil {
		log.V("No preferences found with user_id='%s' (%v)", userID, err)
		return nil, err
	}
	return preferences, nil
}
