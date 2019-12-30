package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
)

// Helpers

// Query Resolvers

type preferencesResolver struct{ *Resolver }

// Mutation Resolvers

func (r *mutationResolver) SavePreferences(ctx context.Context, newPreferences models.InputPreferences) (*models.Preferences, error) {
	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	existingPreferences, err := repos.FindPreferencesByUserID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}
	updatedPreferences, err := repos.SavePreferences(r.DB(ctx), newPreferences, existingPreferences)

	return mappers.PreferencesEntityToModel(updatedPreferences), nil
}

// Field Resolvers

func (r *preferencesResolver) User(ctx context.Context, obj *models.Preferences) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UserID)
}
