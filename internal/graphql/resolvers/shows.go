package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type showResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *showResolver) CreatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return deletedUserByID(ctx, r.ORM(ctx), obj.DeletedByUserID)
}

func (r *showResolver) Admins(ctx context.Context, obj *models.Show) ([]*models.ShowAdmin, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *showResolver) Episodes(ctx context.Context, obj *models.Show) ([]*models.Episode, error) {
	return nil, fmt.Errorf("not implemented")
}
