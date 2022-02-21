package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
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
	panic("mutationResolver.CreateEpisodeURL not implemented")
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	panic("mutationResolver.DeleteEpisodeURL not implemented")
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	panic("mutationResolver.UpdateEpisodeURL not implemented")
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
