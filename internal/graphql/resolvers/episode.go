package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"anime-skip.com/timestamps-service/internal/log"
	"anime-skip.com/timestamps-service/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getEpisodeByID(ctx context.Context, id *uuid.UUID) (*graphql.Episode, error) {
	if id == nil {
		return nil, nil
	}
	internalEpisode, err := r.EpisodeService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	episode := mappers.ToGraphqlEpisode(internalEpisode)
	return &episode, nil
}

func (r *Resolver) getEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Episode, error) {
	internalEpisodes, err := r.EpisodeService.GetByShowID(ctx, *showID)
	if err != nil {
		return nil, err
	}
	episodes := mappers.ToGraphqlEpisodePointers(internalEpisodes)
	return episodes, nil
}

// Mutations

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID *uuid.UUID, episodeInput graphql.InputEpisode) (*graphql.Episode, error) {
	internalInput := internal.Episode{
		BaseEntity: internal.BaseEntity{
			ID: utils.RandomID(),
		},
		ShowID: *showID,
	}
	mappers.ApplyGraphqlInputEpisode(episodeInput, &internalInput)

	created, err := r.EpisodeService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlEpisode(created)
	return &result, nil
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID *uuid.UUID, newEpisode graphql.InputEpisode) (*graphql.Episode, error) {
	log.V("Updating: %v", episodeID)
	existing, err := r.EpisodeService.GetByID(ctx, *episodeID)
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputEpisode(newEpisode, &existing)
	log.V("Updating to %+v", existing)
	created, err := r.EpisodeService.Update(ctx, existing)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	result := mappers.ToGraphqlEpisode(created)
	return &result, nil
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	deleted, err := r.EpisodeService.Delete(ctx, *episodeID)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlEpisode(deleted)
	return &result, nil
}

// Queries

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*graphql.Episode, error) {
	params := internal.GetRecentlyAddedParams{
		Pagination: internal.Pagination{
			Limit:  utils.IntOr(limit, 10),
			Offset: utils.IntOr(offset, 0),
		},
	}
	internalEpisodes, err := r.EpisodeService.GetRecentlyAdded(ctx, params)
	if err != nil {
		return nil, err
	}
	episodes := mappers.ToGraphqlEpisodePointers(internalEpisodes)
	return episodes, nil
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	return r.getEpisodeByID(ctx, episodeID)
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Episode, error) {
	return r.getEpisodesByShowID(ctx, showID)
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *uuid.UUID, offset *int, limit *int, sort *string) ([]*graphql.Episode, error) {
	panic("queryResolver.SearchEpisodes not implemented")
}

func (r *queryResolver) FindEpisodeByName(ctx context.Context, name string) ([]*graphql.ThirdPartyEpisode, error) {
	panic("queryResolver.FindEpisodeByName not implemented")
}

// Fields

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *episodeResolver) Show(ctx context.Context, obj *graphql.Episode) (*graphql.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *graphql.Episode) ([]*graphql.Timestamp, error) {
	return r.getTimestampsByEpisodeID(ctx, obj.ID)
}

func (r *episodeResolver) Urls(ctx context.Context, obj *graphql.Episode) ([]*graphql.EpisodeURL, error) {
	return r.getEpisodeURLsByEpisodeID(ctx, obj.ID)
}

func (r *episodeResolver) Template(ctx context.Context, obj *graphql.Episode) (*graphql.Template, error) {
	return r.getTemplateByEpisodeID(ctx, obj.ID)
}
