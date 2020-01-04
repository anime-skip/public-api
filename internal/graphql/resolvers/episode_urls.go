package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func episodeURLByURL(db *gorm.DB, url string) (*models.EpisodeURL, error) {
	episodeURL, err := repos.FindEpisodeURLByURL(db, url)
	if err != nil {
		return nil, err
	}
	return mappers.EpisodeURLEntityToModel(episodeURL), nil
}

func episodeURLsByEpisodeID(db *gorm.DB, episodeID string) ([]*models.EpisodeURL, error) {
	episodeURLs, err := repos.FindEpisodeURLsByEpisodeID(db, episodeID)
	if err != nil {
		return nil, err
	}

	episodeURLModels := make([]*models.EpisodeURL, len(episodeURLs))
	for index, episodeURL := range episodeURLs {
		episodeURLModels[index] = mappers.EpisodeURLEntityToModel(episodeURL)
	}
	return episodeURLModels, nil
}

// Query Resolvers

type episodeUrlResolver struct{ *Resolver }

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*models.EpisodeURL, error) {
	return episodeURLByURL(r.DB(ctx), episodeURL)
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID string) ([]*models.EpisodeURL, error) {
	return episodeURLsByEpisodeID(r.DB(ctx), episodeID)
}

// Mutation Resolvers

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID string, episodeURLInput models.InputEpisodeURL) (*models.EpisodeURL, error) {
	episodeURL, err := repos.CreateEpisodeURL(r.DB(ctx), episodeID, episodeURLInput)
	if err != nil {
		return nil, err
	}
	return mappers.EpisodeURLEntityToModel(episodeURL), nil
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*models.EpisodeURL, error) {
	db := r.DB(ctx)
	err := repos.DeleteEpisodeURL(r.DB(ctx), false, episodeURL)
	if err != nil {
		return nil, err
	}

	return episodeURLByURL(db.Unscoped(), episodeURL)
}

// Field Resolvers

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) DeletedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *models.EpisodeURL) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), obj.EpisodeID)
}
