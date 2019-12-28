package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type showAdminResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.CreatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.CreatedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *models.ShowAdmin) (*models.Show, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *showAdminResolver) User(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.CreatedByUserID)
}
