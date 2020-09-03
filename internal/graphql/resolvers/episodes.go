package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func episodeByID(db *gorm.DB, episodeID string) (*models.Episode, error) {
	episode, err := repos.FindEpisodeByID(db, episodeID)
	if err != nil {
		return nil, err
	}
	return mappers.EpisodeEntityToModel(episode), nil
}

func episodesByShowID(db *gorm.DB, showID string) ([]*models.Episode, error) {
	episodes, err := repos.FindEpisodesByShowID(db, showID)
	if err != nil {
		return nil, err
	}

	episodeModels := make([]*models.Episode, len(episodes))
	for index, episode := range episodes {
		episodeModels[index] = mappers.EpisodeEntityToModel(episode)
	}
	return episodeModels, nil
}

// Query Resolvers

type episodeResolver struct{ *Resolver }

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID string) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), episodeID)
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID string) ([]*models.Episode, error) {
	return episodesByShowID(r.DB(ctx), showID)
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *string, offset *int, limit *int, sort *string) ([]*models.Episode, error) {
	episodes, err := repos.SearchEpisodes(r.DB(ctx), *search, showID, *offset, *limit, *sort)
	if err != nil {
		return nil, err
	}
	episodeModels := make([]*models.Episode, len(episodes))
	for index, entity := range episodes {
		episodeModels[index] = mappers.EpisodeEntityToModel(entity)
	}
	return episodeModels, nil
}

// Mutation Resolvers

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID string, episodeInput models.InputEpisode) (*models.Episode, error) {
	episode, err := repos.CreateEpisode(r.DB(ctx), showID, episodeInput)
	if err != nil {
		return nil, err
	}
	return mappers.EpisodeEntityToModel(episode), nil
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID string, newEpisode models.InputEpisode) (*models.Episode, error) {
	existingEpisode, err := repos.FindEpisodeByID(r.DB(ctx), episodeID)
	if err != nil {
		return nil, err
	}
	updatedEpisode, err := repos.UpdateEpisode(r.DB(ctx), newEpisode, existingEpisode)

	return mappers.EpisodeEntityToModel(updatedEpisode), nil
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID string) (*models.Episode, error) {
	db := r.DB(ctx)
	err := repos.DeleteEpisode(r.DB(ctx), false, episodeID)
	if err != nil {
		return nil, err
	}

	return episodeByID(db.Unscoped(), episodeID)
}

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

func (r *episodeResolver) Show(ctx context.Context, obj *models.Episode) (*models.Show, error) {
	return showByID(r.DB(ctx), obj.ShowID)
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *models.Episode) ([]*models.Timestamp, error) {
	return timestampsByEpisodeID(r.DB(ctx), obj.ID)
}

func (r *episodeResolver) Urls(ctx context.Context, obj *models.Episode) ([]*models.EpisodeURL, error) {
	return episodeURLsByEpisodeID(r.DB(ctx), obj.ID)
}
