package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTimestampTypeByID(ctx context.Context, id *uuid.UUID) (*internal.TimestampType, error) {
	if id == nil {
		return nil, nil
	}
	timestampType, err := r.TimestampTypeService.Get(ctx, internal.TimestampTypesFilter{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &timestampType, nil
}

// Mutations

func (r *mutationResolver) CreateTimestampType(ctx context.Context, input internal.InputTimestampType) (*internal.TimestampType, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newTimestampType := internal.TimestampType{}
	mappers.ApplyGraphqlInputTimestampType(input, &newTimestampType)

	created, err := r.TimestampTypeService.Create(ctx, newTimestampType, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, id *uuid.UUID, changes internal.InputTimestampType) (*internal.TimestampType, error) {
	log.V("Updating: %v", id)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newTimestampType, err := r.TimestampTypeService.Get(ctx, internal.TimestampTypesFilter{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputTimestampType(changes, &newTimestampType)
	log.V("Updating to %+v", newTimestampType)
	updated, err := r.TimestampTypeService.Update(ctx, newTimestampType, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &updated, nil
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, id *uuid.UUID) (*internal.TimestampType, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.TimestampTypeService.Delete(ctx, *id, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*internal.TimestampType, error) {
	return r.getTimestampTypeByID(ctx, timestampTypeID)
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*internal.TimestampType, error) {
	timestampTypes, err := r.TimestampTypeService.List(ctx, internal.TimestampTypesFilter{})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(timestampTypes), nil
}

// Fields

func (r *timestampTypeResolver) CreatedBy(ctx context.Context, obj *internal.TimestampType) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampTypeResolver) UpdatedBy(ctx context.Context, obj *internal.TimestampType) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampTypeResolver) DeletedBy(ctx context.Context, obj *internal.TimestampType) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}
