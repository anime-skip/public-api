package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type preferencesResolver struct{ *Resolver }

// Mutation Resolvers

func (r *mutationResolver) SavePreferences(ctx context.Context, id string, newPreferences models.InputPreferences) (*models.Preferences, error) {
	userID := "00000000-0000-0000-0000-000000000000"

	existingPreferences, err := repos.FindPreferencesByUserID(ctx, r.ORM(ctx), userID)
	if err != nil {
		return nil, err
	}
	updatedPreferences, err := repos.SavePreferences(ctx, r.ORM(ctx), newPreferences, existingPreferences)

	return mappers.PreferencesEntityToModel(updatedPreferences), nil
}

// Field Resolvers

func (r *preferencesResolver) User(ctx context.Context, obj *models.Preferences) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UserID)
}
