package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Helpers

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
	panic("queryResolver.FindTimestampType not implemented")
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*graphql.TimestampType, error) {
	panic("queryResolver.AllTimestampTypes not implemented")
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
