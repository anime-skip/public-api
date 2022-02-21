package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

// Helpers

// Mutations

func (r *mutationResolver) SavePreferences(ctx context.Context, preferences graphql.InputPreferences) (*graphql.Preferences, error) {
	panic("mutationResolver.SavePreferences not implemented")
}

// Queries

// Fields

func (r *preferencesResolver) User(ctx context.Context, obj *graphql.Preferences) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
