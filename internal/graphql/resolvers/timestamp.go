package resolvers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getTimestampByID(ctx context.Context, id *uuid.UUID) (*internal.Timestamp, error) {
	if id == nil {
		return nil, nil
	}
	timestamp, err := r.TimestampService.Get(ctx, internal.TimestampsFilter{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &timestamp, nil
}

func (r *Resolver) getTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*internal.Timestamp, error) {
	timestamps, err := r.TimestampService.List(ctx, internal.TimestampsFilter{
		EpisodeID: episodeID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(timestamps), nil
}

// Mutations

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID *uuid.UUID, input internal.InputTimestamp) (*internal.Timestamp, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newTimestamp := internal.Timestamp{
		EpisodeID: episodeID,
	}
	mappers.ApplyGraphqlInputTimestamp(input, &newTimestamp)

	created, err := r.TimestampService.Create(ctx, newTimestamp, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, id *uuid.UUID, changes internal.InputTimestamp) (*internal.Timestamp, error) {
	log.V("Updating: %v", id)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newTimestamp, err := r.TimestampService.Get(ctx, internal.TimestampsFilter{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputTimestamp(changes, &newTimestamp)
	log.V("Updating to %+v", newTimestamp)
	updated, err := r.TimestampService.Update(ctx, newTimestamp, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &updated, nil
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, id *uuid.UUID) (*internal.Timestamp, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.TimestampService.Delete(ctx, *id, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

func (r *mutationResolver) UpdateTimestamps(ctx context.Context, create []*internal.InputTimestampOn, update []*internal.InputExistingTimestamp, delete []*uuid.UUID) (*internal.UpdatedTimestamps, error) {
	panic(errors.NewPanicedError("mutationResolver.UpdateTimestamps not implemented"))
}

// Queries

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID *uuid.UUID) (*internal.Timestamp, error) {
	return r.getTimestampByID(ctx, timestampID)
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*internal.Timestamp, error) {
	return r.getTimestampsByEpisodeID(ctx, episodeID)
}

// Fields

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *internal.Timestamp) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *internal.Timestamp) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *internal.Timestamp) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *timestampResolver) Type(ctx context.Context, obj *internal.Timestamp) (*internal.TimestampType, error) {
	return r.getTimestampTypeByID(ctx, obj.TypeID)
}

func (r *timestampResolver) Episode(ctx context.Context, obj *internal.Timestamp) (*internal.Episode, error) {
	return r.getEpisodeByID(ctx, obj.EpisodeID)
}
