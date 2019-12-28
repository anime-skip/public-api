package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type showAdminResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UpdatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return deletedUserByID(ctx, r.ORM(ctx), obj.DeletedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *models.ShowAdmin) (*models.Show, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *showAdminResolver) User(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UserID)
}
