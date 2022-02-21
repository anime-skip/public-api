package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Helpers

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
	panic("queryResolver.FindEpisodeURL not implemented")
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.EpisodeURL, error) {
	panic("queryResolver.FindEpisodeUrlsByEpisodeID not implemented")
}

// Fields

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.Episode, error) {
	panic("episodeUrlResolver.Episode not implemented")
}
