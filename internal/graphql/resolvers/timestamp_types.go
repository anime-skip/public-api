package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database/mappers"
	"anime-skip.com/backend/internal/database/repos"
	"anime-skip.com/backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

var _timestampTypeCache = map[string]*models.TimestampType{}
var _timestampTypeCacheHasAll = false

// Helpers

func timestampTypeByID(db *gorm.DB, timestampTypeID string) (*models.TimestampType, error) {
	if timestampType, isCached := _timestampTypeCache[timestampTypeID]; isCached {
		return timestampType, nil
	}
	timestampType, err := repos.FindTimestampTypeByID(db, timestampTypeID)
	if err != nil {
		return nil, err
	}

	model := mappers.TimestampTypeEntityToModel(timestampType)
	_timestampTypeCache[model.ID] = model

	return model, nil
}

func allTimestampTypes(db *gorm.DB) ([]*models.TimestampType, error) {
	if _timestampTypeCacheHasAll {
		timestampTypeModels := []*models.TimestampType{}
		for _, model := range _timestampTypeCache {
			timestampTypeModels = append(timestampTypeModels, model)
		}
		return timestampTypeModels, nil
	}
	timestampTypes, err := repos.FindAllTimestampTypes(db)
	if err != nil {
		return nil, err
	}

	timestampTypeModels := make([]*models.TimestampType, len(timestampTypes))
	for index, timestampType := range timestampTypes {
		model := mappers.TimestampTypeEntityToModel(timestampType)
		timestampTypeModels[index] = model
		_timestampTypeCache[model.ID] = model
	}
	_timestampTypeCacheHasAll = true
	return timestampTypeModels, nil
}

// Query Resolvers

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID string) (*models.TimestampType, error) {
	return timestampTypeByID(r.DB(ctx), timestampTypeID)
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*models.TimestampType, error) {
	return allTimestampTypes(r.DB(ctx))
}

// Mutation Resolvers

func (r *mutationResolver) CreateTimestampType(ctx context.Context, timestampTypeInput models.InputTimestampType) (*models.TimestampType, error) {
	timestampType, err := repos.CreateTimestampType(r.DB(ctx), timestampTypeInput)
	if err != nil {
		return nil, err
	}

	model := mappers.TimestampTypeEntityToModel(timestampType)
	_timestampTypeCache[model.ID] = model

	return model, nil
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID string, newTimestampType models.InputTimestampType) (*models.TimestampType, error) {
	existingTimestampType, err := repos.FindTimestampTypeByID(r.DB(ctx), timestampTypeID)
	if err != nil {
		return nil, err
	}
	updatedTimestampType, err := repos.UpdateTimestampType(r.DB(ctx), newTimestampType, existingTimestampType)

	model := mappers.TimestampTypeEntityToModel(updatedTimestampType)
	_timestampTypeCache[model.ID] = model

	return model, nil
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID string) (*models.TimestampType, error) {
	db := r.DB(ctx)
	err := repos.DeleteTimestampType(r.DB(ctx), false, timestampTypeID)
	if err != nil {
		return nil, err
	}

	model, err := timestampTypeByID(db.Unscoped(), timestampTypeID)
	if err != nil {
		return nil, err
	}
	delete(_timestampTypeCache, model.ID)

	return model, nil
}

// Field Resolvers

type timestampTypeResolver struct{ *Resolver }

func (r *timestampTypeResolver) CreatedBy(ctx context.Context, obj *models.TimestampType) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *timestampTypeResolver) UpdatedBy(ctx context.Context, obj *models.TimestampType) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *timestampTypeResolver) DeletedBy(ctx context.Context, obj *models.TimestampType) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}
