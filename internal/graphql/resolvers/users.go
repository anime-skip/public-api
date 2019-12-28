package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

func (r *queryResolver) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, userID string) (*models.User, error) {
	return repos.FindUserByUsername(ctx, r.ORM, userID)
}

// Mutation Resolvers

func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
	return false, fmt.Errorf("Not implemented")
}

// Field Resolvers

type userResolver struct{ *Resolver }

func (r *userResolver) AdminOfShows(ctx context.Context, obj *models.User) ([]*models.ShowAdmin, error) {
	return nil, fmt.Errorf("Not implemented")
}
