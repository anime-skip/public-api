package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
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
	panic("mutationResolver.CreateTimestampType not implemented")
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID *uuid.UUID, newTimestampType graphql.InputTimestampType) (*graphql.TimestampType, error) {
	panic("mutationResolver.UpdateTimestampType not implemented")
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*graphql.TimestampType, error) {
	panic("mutationResolver.DeleteTimestampType not implemented")
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
