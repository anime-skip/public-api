package resolvers

import (
	go_context "context"

	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getPreferences(ctx go_context.Context, userID uuid.UUID) (*graphql.Preferences, error) {
	prefs, err := r.PreferencesService.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	gqlPrefs := mappers.ToGraphqlPreferences(prefs)
	return &gqlPrefs, nil
}

// Mutations

func (r *mutationResolver) SavePreferences(ctx go_context.Context, preferences map[string]interface{}) (*graphql.Preferences, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	// Apply updates to struct
	existingPrefs, err := r.getPreferences(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	err = utils.ApplyChanges(preferences, &existingPrefs)
	if err != nil {
		return nil, err
	}
	newInternalPrefs := mappers.ToInternalPreferences(*existingPrefs)

	// Update data
	updatedInternalPrefs, err := r.PreferencesService.Update(ctx, newInternalPrefs)
	if err != nil {
		return nil, err
	}

	updatedPrefs := mappers.ToGraphqlPreferences(updatedInternalPrefs)
	return &updatedPrefs, nil
}

// Queries

// Fields

func (r *preferencesResolver) User(ctx go_context.Context, obj *graphql.Preferences) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
