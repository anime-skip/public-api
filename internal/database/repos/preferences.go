package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

// SavePreferences updates the preferences for the given preferences id
func SavePreferences(ctx context.Context, orm *database.ORM, newPreferences models.InputPreferences, existingPreferences *entities.Preferences) (*entities.Preferences, error) {
	data := mappers.PreferencesInputModelToEntity(newPreferences, existingPreferences)
	err := orm.DB.Model(data).Update(*data).Error
	if err != nil {
		log.E("Failed to update preferences for [%+v]: %v", data, err)
		return nil, fmt.Errorf("Failed to update preferences with id='%s'", data.ID)
	}
	return data, err
}

// FindPreferencesByUserID finds a set of preferences by the user they belong to
func FindPreferencesByUserID(ctx context.Context, orm *database.ORM, userID string) (*entities.Preferences, error) {
	preferences := &entities.Preferences{}
	err := orm.DB.Where("user_id = ?", userID).Find(preferences).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No preferences found with user_id='%s'", userID)
	}
	return preferences, nil
}
