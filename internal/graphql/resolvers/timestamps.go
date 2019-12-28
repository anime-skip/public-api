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
	return nil, fmt.Errorf("not implemented")
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *timestampResolver) Type(ctx context.Context, obj *models.Timestamp) (*models.TimestampType, error) {
	return nil, fmt.Errorf("not implemented")
}
