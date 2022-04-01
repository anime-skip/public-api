package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"anime-skip.com/public-api/internal/log"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getEpisodeURLByURL(ctx context.Context, url string) (*graphql.EpisodeURL, error) {
	internalEpisodeURL, err := r.EpisodeURLService.GetByURL(ctx, url)
	if err != nil {
		return nil, err
	}
	episodeURL := mappers.ToGraphqlEpisodeURL(internalEpisodeURL)
	return &episodeURL, nil
}

func (r *Resolver) getEpisodeURLsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.EpisodeURL, error) {
	internalEpisodeURLs, err := r.EpisodeURLService.GetByEpisodeId(ctx, *episodeID)
	if err != nil {
		return nil, err
	}
	episodeURLs := mappers.ToGraphqlEpisodeURLPointers(internalEpisodeURLs)
	return episodeURLs, nil
}

// Mutations

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID *uuid.UUID, episodeURLInput graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	internalInput := internal.EpisodeURL{
		EpisodeID: *episodeID,
	}
	mappers.ApplyGraphqlInputEpisodeURL(episodeURLInput, &internalInput)

	created, err := r.EpisodeURLService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlEpisodeURL(created)
	return &result, nil
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	deleted, err := r.EpisodeURLService.Delete(ctx, episodeURL)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlEpisodeURL(deleted)
	return &result, nil
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	log.V("Updating: %v", episodeURL)
	existing, err := r.EpisodeURLService.GetByURL(ctx, episodeURL)
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputEpisodeURL(newEpisodeURL, &existing)
	log.V("Updating to %+v", existing)
	created, err := r.EpisodeURLService.Update(ctx, existing)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	result := mappers.ToGraphqlEpisodeURL(created)
	return &result, nil
}

// Queries

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	return r.getEpisodeURLByURL(ctx, episodeURL)
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.EpisodeURL, error) {
	return r.getEpisodeURLsByEpisodeID(ctx, episodeID)
}

// Fields

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.Episode, error) {
	return r.getEpisodeByID(ctx, obj.EpisodeID)
}
