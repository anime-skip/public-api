package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

// Mutations

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID *uuid.UUID, episodeInput graphql.InputEpisode) (*graphql.Episode, error) {
	panic("mutationResolver.CreateEpisode not implemented")
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID *uuid.UUID, newEpisode graphql.InputEpisode) (*graphql.Episode, error) {
	panic("mutationResolver.UpdateEpisode not implemented")
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	panic("mutationResolver.DeleteEpisode not implemented")
}

// Queries

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*graphql.Episode, error) {
	panic("queryResolver.RecentlyAddedEpisodes not implemented")
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	panic("queryResolver.FindEpisode not implemented")
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Episode, error) {
	panic("queryResolver.FindEpisodesByShowID not implemented")
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
	panic("episodeResolver.Show not implemented")
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *graphql.Episode) ([]*graphql.Timestamp, error) {
	panic("episodeResolver.Timestamps not implemented")
}

func (r *episodeResolver) Urls(ctx context.Context, obj *graphql.Episode) ([]*graphql.EpisodeURL, error) {
	panic("episodeResolver.Urls not implemented")
}

func (r *episodeResolver) Template(ctx context.Context, obj *graphql.Episode) (*graphql.Template, error) {
	panic("episodeResolver.Template not implemented")
}
