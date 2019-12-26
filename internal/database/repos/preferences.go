package repos

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
)

// FindPreferencesByUserID finds a set of preferences by the user they belong to
func FindPreferencesByUserID(ctx context.Context, orm *database.ORM, userID string) (*models.Preferences, error) {
	preferences := &entities.Preferences{}
	err := orm.DB.Where("user_id = ?", userID).Find(preferences).Error
	if err != nil {
		log.E("Failed query: %v", err)
		return nil, fmt.Errorf("No preferences found with user_id='%s'", userID)
	}
	return mappers.PreferencesEntityToModel(preferences)
}
