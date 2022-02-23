package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTimestampByID(ctx context.Context, id *uuid.UUID) (*graphql.Timestamp, error) {
	if id == nil {
		return nil, nil
	}
	internalTimestamp, err := r.TimestampService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	timestamp := mappers.ToGraphqlTimestamp(internalTimestamp)
	return &timestamp, nil
}

func (r *Resolver) getTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.Timestamp, error) {
	internalTimestamps, err := r.TimestampService.GetByEpisodeID(ctx, *episodeID)
	if err != nil {
		return nil, err
	}
	episodes := mappers.ToGraphqlTimestampPointers(internalTimestamps)
	return episodes, nil
}

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
	return r.getTimestampByID(ctx, timestampID)
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.Timestamp, error) {
	return r.getTimestampsByEpisodeID(ctx, episodeID)
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
	return r.getTimestampTypeByID(ctx, obj.TypeID)
}

func (r *timestampResolver) Episode(ctx context.Context, obj *graphql.Timestamp) (*graphql.Episode, error) {
	return r.getEpisodeByID(ctx, obj.EpisodeID)
}
