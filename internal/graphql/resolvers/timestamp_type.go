package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTimestampTypeByID(ctx context.Context, id *uuid.UUID) (*graphql.TimestampType, error) {
	if id == nil {
		return nil, nil
	}
	internalTimestampType, err := r.TimestampTypeService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	timestampType := mappers.ToGraphqlTimestampType(internalTimestampType)
	return &timestampType, nil
}

// Mutations

func (r *mutationResolver) CreateTimestampType(ctx context.Context, timestampTypeInput graphql.InputTimestampType) (*graphql.TimestampType, error) {
	internalInput := internal.TimestampType{
		BaseEntity: internal.BaseEntity{
			ID: utils.RandomID(),
		},
	}
	mappers.ApplyGraphqlInputTimestampType(timestampTypeInput, &internalInput)

	created, err := r.TimestampTypeService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTimestampType(created)
	return &result, nil
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID *uuid.UUID, newTimestampType graphql.InputTimestampType) (*graphql.TimestampType, error) {
	log.V("Updating: %v", timestampTypeID)
	existing, err := r.TimestampTypeService.GetByID(ctx, *timestampTypeID)
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputTimestampType(newTimestampType, &existing)
	log.V("Updating to %+v", existing)
	created, err := r.TimestampTypeService.Update(ctx, existing)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	result := mappers.ToGraphqlTimestampType(created)
	return &result, nil
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*graphql.TimestampType, error) {
	deleted, err := r.TimestampTypeService.Delete(ctx, *timestampTypeID)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlTimestampType(deleted)
	return &result, nil
}

// Queries

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*graphql.TimestampType, error) {
	return r.getTimestampTypeByID(ctx, timestampTypeID)
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*graphql.TimestampType, error) {
	internalTypes, err := r.TimestampTypeService.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	timestampType := mappers.ToGraphqlTimestampTypePointers(internalTypes)
	return timestampType, nil
}

// Fields

func (r *timestampTypeResolver) CreatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampTypeResolver) UpdatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampTypeResolver) DeletedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
