package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils/context_utils"
)

// Helpers

// Query Resolvers

type preferencesResolver struct{ *Resolver }

// Mutation Resolvers

func (r *mutationResolver) SavePreferences(ctx context.Context, newPreferences models.InputPreferences) (*models.Preferences, error) {
	userID, err := context_utils.UserID(ctx)
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
