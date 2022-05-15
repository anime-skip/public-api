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

func (r *Resolver) getEpisodeByID(ctx context.Context, id *uuid.UUID) (*internal.Episode, error) {
	if id == nil {
		return nil, nil
	}
	episode, err := r.EpisodeService.Get(ctx, internal.EpisodesFilter{
		ID:             id,
		IncludeDeleted: true,
	})
	if err != nil {
		return nil, err
	}
	return &episode, nil
}

func (r *Resolver) getEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.Episode, error) {
	episodes, err := r.EpisodeService.List(ctx, internal.EpisodesFilter{
		ShowID: showID,
	})
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(episodes), nil
}

// Mutations

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID *uuid.UUID, episodeInput internal.InputEpisode) (*internal.Episode, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	internalInput := internal.Episode{
		ShowID: showID,
	}
	mappers.ApplyGraphqlInputEpisode(episodeInput, &internalInput)

	created, err := r.EpisodeService.Create(ctx, internalInput, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID *uuid.UUID, newEpisode internal.InputEpisode) (*internal.Episode, error) {
	log.V("Updating: %v", episodeID)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := r.EpisodeService.Get(ctx, internal.EpisodesFilter{
		ID: episodeID,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputEpisode(newEpisode, &existing)
	log.V("Updating to %+v", existing)
	updated, err := r.EpisodeService.Update(ctx, existing, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &updated, nil
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID *uuid.UUID) (*internal.Episode, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.EpisodeService.Delete(ctx, *episodeID, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*internal.Episode, error) {
	filter := internal.RecentlyAddedEpisodesFilter{
		Pagination: internal.Pagination{
			Limit:  utils.ValueOr(limit, 10),
			Offset: utils.ValueOr(offset, 0),
		},
	}
	episodes, err := r.EpisodeService.ListRecentlyAdded(ctx, filter)
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(episodes), nil
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID *uuid.UUID) (*internal.Episode, error) {
	return r.getEpisodeByID(ctx, episodeID)
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*internal.Episode, error) {
	return r.getEpisodesByShowID(ctx, showID)
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *uuid.UUID, offset *int, limit *int, sort *string) ([]*internal.Episode, error) {
	panic(errors.NewPanicedError("queryResolver.SearchEpisodes not implemented"))
}

func (r *queryResolver) FindEpisodeByName(ctx context.Context, name string) ([]*internal.ThirdPartyEpisode, error) {
	episodes, err := r.ThirdPartyService.FindEpisodeByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return utils.PtrSlice(episodes), nil
}

// Fields

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *internal.Episode) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *internal.Episode) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *internal.Episode) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *episodeResolver) Show(ctx context.Context, obj *internal.Episode) (*internal.Show, error) {
	return r.getShowById(ctx, obj.ShowID)
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *internal.Episode) ([]*internal.Timestamp, error) {
	return r.getTimestampsByEpisodeID(ctx, obj.ID)
}

func (r *episodeResolver) Urls(ctx context.Context, obj *internal.Episode) ([]*internal.EpisodeURL, error) {
	return r.getEpisodeURLsByEpisodeID(ctx, obj.ID)
}

func (r *episodeResolver) Template(ctx context.Context, obj *internal.Episode) (*internal.Template, error) {
	return r.getTemplateByEpisodeID(ctx, obj.ID)
}
