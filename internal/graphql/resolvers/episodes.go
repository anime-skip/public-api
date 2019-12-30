package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type episodeResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}
