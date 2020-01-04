package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func timestampTypeByID(db *gorm.DB, timestampTypeID string) (*models.TimestampType, error) {
	timestampType, err := repos.FindTimestampTypeByID(db, timestampTypeID)
	if err != nil {
		return nil, err
	}
	return mappers.TimestampTypeEntityToModel(timestampType), nil
}

func allTimestampTypes(db *gorm.DB) ([]*models.TimestampType, error) {
	timestampTypes, err := repos.FindAllTimestampTypes(db)
	if err != nil {
		return nil, err
	}

	timestampTypeModels := make([]*models.TimestampType, len(timestampTypes))
	for index, timestampType := range timestampTypes {
		timestampTypeModels[index] = mappers.TimestampTypeEntityToModel(timestampType)
	}
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
	return mappers.TimestampTypeEntityToModel(timestampType), nil
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID string, newTimestampType models.InputTimestampType) (*models.TimestampType, error) {
	existingTimestampType, err := repos.FindTimestampTypeByID(r.DB(ctx), timestampTypeID)
	if err != nil {
		return nil, err
	}
	updatedTimestampType, err := repos.UpdateTimestampType(r.DB(ctx), newTimestampType, existingTimestampType)

	return mappers.TimestampTypeEntityToModel(updatedTimestampType), nil
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID string) (*models.TimestampType, error) {
	db := r.DB(ctx)
	err := repos.DeleteTimestampType(r.DB(ctx), false, timestampTypeID)
	if err != nil {
		return nil, err
	}

	return timestampTypeByID(db.Unscoped(), timestampTypeID)
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
