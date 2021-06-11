package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"anime-skip.com/backend/internal/utils"
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
	var err error
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	entity, err := repos.FindEpisodeURLByURL(tx, episodeURL)
	if err != nil {
		return nil, err
	}

	err = repos.DeleteEpisodeURL(tx, episodeURL)
	if err != nil {
		return nil, err
	}

	return mappers.EpisodeURLEntityToModel(entity), nil
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL models.InputEpisodeURL) (*models.EpisodeURL, error) {
	existingEpisodeURL, err := repos.FindEpisodeURLByURL(r.DB(ctx), episodeURL)
	if err != nil {
		return nil, err
	}
	updatedEpisodeURL, err := repos.UpdateEpisodeURL(r.DB(ctx), newEpisodeURL, existingEpisodeURL)
	if err != nil {
		return nil, err
	}
	return mappers.EpisodeURLEntityToModel(updatedEpisodeURL), nil
}

// Field Resolvers

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *models.EpisodeURL) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), obj.EpisodeID)
}
