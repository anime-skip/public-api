package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Helpers

// Mutations

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID *uuid.UUID, timestampInput graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("mutationResolver.CreateTimestamp not implemented")
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID *uuid.UUID, newTimestamp graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("mutationResolver.UpdateTimestamp not implemented")
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID *uuid.UUID) (*graphql.Timestamp, error) {
	panic("mutationResolver.DeleteTimestamp not implemented")
}

func (r *mutationResolver) UpdateTimestamps(ctx context.Context, create []*graphql.InputTimestampOn, update []*graphql.InputExistingTimestamp, delete []*uuid.UUID) (*graphql.UpdatedTimestamps, error) {
	panic("mutationResolver.UpdateTimestamps not implemented")
}

// Queries

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID *uuid.UUID) (*graphql.Timestamp, error) {
	panic("queryResolver.FindTimestamp not implemented")
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.Timestamp, error) {
	panic("queryResolver.FindTimestampsByEpisodeID not implemented")
}

// Fields

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *timestampResolver) Type(ctx context.Context, obj *graphql.Timestamp) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *timestampResolver) Episode(ctx context.Context, obj *graphql.Timestamp) (*graphql.Episode, error) {
	panic("not implemented")
}
