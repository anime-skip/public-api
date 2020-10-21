package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	. "anime-skip.com/backend/internal/services"
	"anime-skip.com/backend/internal/utils/log"
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
type thirdPartyEpisodeResolver struct{ *Resolver }

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*models.Episode, error) {
	episodes, err := repos.RecentlyAddedEpisodes(r.DB(ctx), *limit, *offset)
	if err != nil {
		return nil, err
	}
	episodeModels := make([]*models.Episode, len(episodes))
	for index, episode := range episodes {
		episodeModels[index] = mappers.EpisodeEntityToModel(episode)
	}
	return episodeModels, nil
}

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

func (r *queryResolver) FindEpisodeByName(ctx context.Context, episodeName string) ([]*models.ThirdPartyEpisode, error) {
	standardizedName := BetterVRV.StandardizeEpisodeName(episodeName)
	animeSkipEpisodes, err := repos.FindEpisodesByExactName(r.DB(ctx), standardizedName)
	if err != nil {
		log.E("Failed to lookup episodes from database: %v", err)
	}
	mappedAnimeSkipEpisodes := []*models.ThirdPartyEpisode{}
	for _, episode := range animeSkipEpisodes {
		mappedAnimeSkipEpisodes = append(
			mappedAnimeSkipEpisodes,
			mappers.EpisodeEntityToThirdPartyEpisodeModel(episode),
		)
	}

	thirdPartyEpisodes, err := BetterVRV.FetchEpisodesByName(episodeName)
	if err != nil {
		log.E("Failed to fetch episodes from BetterVRV: %v", err)
		return nil, err
	}

	return append(mappedAnimeSkipEpisodes, thirdPartyEpisodes...), nil
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

func (r *thirdPartyEpisodeResolver) Timestamps(ctx context.Context, obj *models.ThirdPartyEpisode) ([]*models.ThirdPartyTimestamp, error) {
	if obj.ID == nil {
		return obj.Timestamps, nil
	}
	timestamps, err := timestampsByEpisodeID(r.DB(ctx), *obj.ID)
	if err != nil {
		return nil, err
	}
	mappedTimestamps := []*models.ThirdPartyTimestamp{}
	for _, timestamp := range timestamps {
		mappedTimestamps = append(mappedTimestamps, mappers.TimestampModelToThirdPartyTimestamp(timestamp))
	}
	return mappedTimestamps, nil
}
