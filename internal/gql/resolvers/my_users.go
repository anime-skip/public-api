package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// Query Resolvers

type myUserResolver struct{ *Resolver }

func (r *queryResolver) MyUser(ctx context.Context, username string) (*models.MyUser, error) {
	return repos.FindMyUser(ctx, r.ORM, username)
}

// Mutation Resolvers

// Field Resolvers

func (r *myUserResolver) AdminOfShows(ctx context.Context, obj *models.MyUser) ([]*models.ShowAdmin, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *myUserResolver) Preferences(ctx context.Context, obj *models.MyUser) (*models.Preferences, error) {
	return repos.FindPreferencesByUserID(ctx, r.ORM, obj.ID)
}
