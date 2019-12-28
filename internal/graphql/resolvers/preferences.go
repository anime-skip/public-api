package resolvers

import (
	"fmt"
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type preferencesResolver struct{ *Resolver }

// Mutation Resolvers

func (r *mutationResolver) SavePreferences(ctx context.Context, id string, preferences models.InputPreferences) (*models.Preferences, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Field Resolvers

func (r *preferencesResolver) User(ctx context.Context, obj *models.Preferences) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.UserID)
}
