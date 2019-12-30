package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type timestampResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return userByID(ctx, r.DB(ctx), obj.CreatedByUserID)
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return userByID(ctx, r.DB(ctx), obj.UpdatedByUserID)
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return deletedUserByID(ctx, r.DB(ctx), obj.DeletedByUserID)
}

func (r *timestampResolver) Type(ctx context.Context, obj *models.Timestamp) (*models.TimestampType, error) {
	return nil, fmt.Errorf("not implemented")
}
