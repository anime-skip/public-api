package resolvers

import (
	go_context "context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getPreferences(ctx go_context.Context, userID uuid.UUID) (*internal.Preferences, error) {
	prefs, err := r.PreferencesService.Get(ctx, internal.PreferencesFilter{
		UserID: &userID,
	})
	if err != nil {
		return nil, err
	}
	return &prefs, nil
}

// Mutations

func (r *mutationResolver) SavePreferences(ctx go_context.Context, changes map[string]any) (*internal.Preferences, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	// Apply updates to struct
	newPrefs, err := r.getPreferences(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	err = utils.ApplyChanges(changes, &newPrefs)
	if err != nil {
		return nil, err
	}

	// Update data
	updated, err := r.PreferencesService.Update(ctx, *newPrefs)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

// Queries

// Fields

func (r *preferencesResolver) User(ctx go_context.Context, obj *internal.Preferences) (*internal.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
