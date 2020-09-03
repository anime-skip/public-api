package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func timestampByID(db *gorm.DB, timestampID string) (*models.Timestamp, error) {
	timestamp, err := repos.FindTimestampByID(db, timestampID)
	if err != nil {
		return nil, err
	}
	return mappers.TimestampEntityToModel(timestamp), nil
}

func timestampsByEpisodeID(db *gorm.DB, episodeID string) ([]*models.Timestamp, error) {
	timestamps, err := repos.FindTimestampsByEpisodeID(db, episodeID)
	if err != nil {
		return nil, err
	}

	timestampModels := make([]*models.Timestamp, len(timestamps))
	for index, timestamp := range timestamps {
		timestampModels[index] = mappers.TimestampEntityToModel(timestamp)
	}
	return timestampModels, nil
}

// Query Resolvers

type timestampResolver struct{ *Resolver }

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID string) (*models.Timestamp, error) {
	return timestampByID(r.DB(ctx), timestampID)
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID string) ([]*models.Timestamp, error) {
	return timestampsByEpisodeID(r.DB(ctx), episodeID)
}

// Mutation Resolvers

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID string, timestampInput models.InputTimestamp) (*models.Timestamp, error) {
	timestamp, err := repos.CreateTimestamp(r.DB(ctx), episodeID, timestampInput)
	if err != nil {
		return nil, err
	}
	return mappers.TimestampEntityToModel(timestamp), nil
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID string, newTimestamp models.InputTimestamp) (*models.Timestamp, error) {
	existingTimestamp, err := repos.FindTimestampByID(r.DB(ctx), timestampID)
	if err != nil {
		return nil, err
	}
	updatedTimestamp, err := repos.UpdateTimestamp(r.DB(ctx), newTimestamp, existingTimestamp)

	return mappers.TimestampEntityToModel(updatedTimestamp), nil
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID string) (*models.Timestamp, error) {
	db := r.DB(ctx)
	err := repos.DeleteTimestamp(r.DB(ctx), false, timestampID)
	if err != nil {
		return nil, err
	}

	return timestampByID(db.Unscoped(), timestampID)
}

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
	return timestampTypeByID(r.DB(ctx), obj.TypeID)
}

func (r *timestampResolver) Episode(ctx context.Context, obj *models.Timestamp) (*models.Episode, error) {
	return episodeByID(r.DB(ctx), obj.EpisodeID)
}
