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

func createTimestamp(tx *gorm.DB, episodeID string, timestampInput models.InputTimestamp) (*models.Timestamp, error) {
	timestamp, err := repos.CreateTimestamp(tx, episodeID, timestampInput)
	if err != nil {
		return nil, err
	}
	return mappers.TimestampEntityToModel(timestamp), nil
}

func updateTimestamp(tx *gorm.DB, timestampID string, newTimestamp models.InputTimestamp) (*models.Timestamp, error) {
	existingTimestamp, err := repos.FindTimestampByID(tx, timestampID)
	if err != nil {
		return nil, err
	}
	updatedTimestamp, err := repos.UpdateTimestamp(tx, newTimestamp, existingTimestamp)
	if err != nil {
		return nil, err
	}
	return mappers.TimestampEntityToModel(updatedTimestamp), nil
}

func deleteTimestamp(tx *gorm.DB, timestampID string) (*models.Timestamp, error) {
	err := repos.DeleteTimestamp(tx, timestampID)
	if err != nil {
		return nil, err
	}
	return timestampByID(tx, timestampID)
}

// Query Resolvers

type timestampResolver struct{ *Resolver }
type thirdPartyTimestampResolver struct{ *Resolver }

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID string) (*models.Timestamp, error) {
	return timestampByID(r.DB(ctx), timestampID)
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID string) ([]*models.Timestamp, error) {
	return timestampsByEpisodeID(r.DB(ctx), episodeID)
}

// Mutation Resolvers

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID string, timestampInput models.InputTimestamp) (result *models.Timestamp, err error) {
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	result, err = createTimestamp(tx, episodeID, timestampInput)
	return result, err
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID string, newTimestamp models.InputTimestamp) (result *models.Timestamp, err error) {
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	return updateTimestamp(tx, timestampID, newTimestamp)
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID string) (result *models.Timestamp, err error) {
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	return deleteTimestamp(tx, timestampID)
}

func (r *mutationResolver) UpdateTimestamps(
	ctx context.Context,
	create []*models.InputTimestampOn,
	update []*models.InputExistingTimestamp,
	delete []string,
) (result *models.UpdatedTimestamps, err error) {
	tx, commitOrRollback := utils.StartTransaction2(r.DB(ctx), &err)
	defer commitOrRollback()

	// Don't modify the result directly in case of an unhandled error
	created := []*models.Timestamp{}
	updated := []*models.Timestamp{}
	deleted := []*models.Timestamp{}

	for _, timestamp := range create {
		newTimestamp, err := createTimestamp(tx, timestamp.EpisodeID, *timestamp.Timestamp)
		if err != nil {
			return nil, err
		}
		created = append(created, newTimestamp)
	}

	for _, timestamp := range update {
		updatedTimestamp, err := updateTimestamp(tx, timestamp.ID, *timestamp.Timestamp)
		if err != nil {
			return nil, err
		}
		updated = append(updated, updatedTimestamp)
	}

	for _, timestampID := range delete {
		deletedTimestamp, err := deleteTimestamp(tx, timestampID)
		if err != nil {
			return nil, err
		}
		deleted = append(deleted, deletedTimestamp)
	}

	result = &models.UpdatedTimestamps{
		Created: created,
		Updated: updated,
		Deleted: deleted,
	}
	return result, nil
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

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *models.ThirdPartyTimestamp) (*models.TimestampType, error) {
	return timestampTypeByID(r.DB(ctx), obj.TypeID)
}
