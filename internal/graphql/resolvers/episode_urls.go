package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Query Resolvers

type episodeUrlResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}
