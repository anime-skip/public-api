package resolvers

import (
	go_context "context"

	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"anime-skip.com/timestamps-service/internal/utils"
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
	auth, err := context.GetAuthenticationDetails(ctx)
	if err != nil {
		return nil, err
	}
	existingPrefs, err := r.getPreferences(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	err = utils.ApplyChanges(preferences, &existingPrefs)
	if err != nil {
		return nil, err
	}
	newPrefs := mappers.ToInternalPreferences(*existingPrefs)
	updatedPrefs, err := r.PreferencesService.Update(ctx, newPrefs)
	if err != nil {
		return nil, err
	}
	updatedGqlPrefs := mappers.ToGraphqlPreferences(updatedPrefs)
	return &updatedGqlPrefs, nil
}

// Queries

// Fields

func (r *preferencesResolver) User(ctx go_context.Context, obj *graphql.Preferences) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
