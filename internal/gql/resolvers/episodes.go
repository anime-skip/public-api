package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/gql/models"
)

// Query Resolvers

type episodeResolver struct{ *Resolver }

// Mutation Resolvers

// Field Resolvers

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.CreatedByUserID)
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return repos.FindUserByID(ctx, r.ORM, obj.UpdatedByUserID)
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
	return repos.FindUserByIDPtr(ctx, r.ORM, obj.DeletedByUserID)
}
