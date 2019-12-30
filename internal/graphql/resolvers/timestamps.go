package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils/log"
)

// Helpers

// Query Resolvers

type timestampResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *timestampResolver) Type(ctx context.Context, obj *models.Timestamp) (*models.TimestampType, error) {
	log.W("TODO - add timestamp type field resolver for timestamp model")
	return nil, nil
}
