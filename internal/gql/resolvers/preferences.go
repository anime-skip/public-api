package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// Query Resolvers

type preferencesResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *preferencesResolver) User(ctx context.Context, obj *models.Preferences) (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}
