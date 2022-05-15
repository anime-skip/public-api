package resolvers

import (
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/context"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/mappers"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getShowById(ctx context.Context, id *uuid.UUID) (*internal.Show, error) {
	if id == nil {
		return nil, nil
	}
	show, err := r.ShowService.Get(ctx, internal.ShowsFilter{
		ID:             id,
		IncludeDeleted: true,
	})
	if err != nil {
		return nil, err
	}
	return &show, nil
}

// Mutations

func (r *mutationResolver) CreateShow(ctx context.Context, showInput internal.InputShow, becomeAdmin bool) (*internal.Show, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	newShow := internal.Show{}
	mappers.ApplyGraphqlInputShow(showInput, &newShow)

	created, err := r.ShowService.Create(ctx, newShow, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID *uuid.UUID, newShow internal.InputShow) (*internal.Show, error) {
	log.V("Updating: %v", showID)
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := r.ShowService.Get(ctx, internal.ShowsFilter{
		ID: showID,
	})
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputShow(newShow, &existing)
	log.V("Updating to %+v", existing)
	updated, err := r.ShowService.Update(ctx, existing, auth.UserID)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	return &updated, nil
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID *uuid.UUID) (*internal.Show, error) {
	auth, err := context.GetAuthClaims(ctx)
	if err != nil {
		return nil, err
	}

	deleted, err := r.ShowService.Delete(ctx, *showID, auth.UserID)
	if err != nil {
		return nil, err
	}

	return &deleted, nil
}

// Queries

func (r *queryResolver) FindShow(ctx context.Context, showID *uuid.UUID) (*internal.Show, error) {
	return r.getShowById(ctx, showID)
}

var defaultSearchShowFilter = internal.ShowsFilter{
	Pagination: &internal.Pagination{
		Offset: 0,
		Limit:  25,
	},
	Sort: "ASC",
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*internal.Show, error) {
	filter := defaultSearchShowFilter
	if search != nil {
		filter.NameContains = utils.Ptr(strings.TrimSpace(*search))
	}
	if offset != nil {
		filter.Pagination.Offset = *offset
	}
	if limit != nil {
		filter.Pagination.Limit = *limit
	}
	if sort != nil {
		filter.Sort = *sort
	}

	shows, err := r.ShowService.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return utils.PtrSlice(shows), nil
}

// Fields

func (r *showResolver) CreatedBy(ctx context.Context, obj *internal.Show) (*internal.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *internal.Show) (*internal.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *internal.Show) (*internal.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *showResolver) Admins(ctx context.Context, obj *internal.Show) ([]*internal.ShowAdmin, error) {
	return r.getShowAdminsByShowId(ctx, obj.ID)
}

func (r *showResolver) Episodes(ctx context.Context, obj *internal.Show) ([]*internal.Episode, error) {
	return r.getEpisodesByShowID(ctx, obj.ID)
}

func (r *showResolver) Templates(ctx context.Context, obj *internal.Show) ([]*internal.Template, error) {
	return r.getTemplatesByShowID(ctx, obj.ID)
}

func (r *showResolver) SeasonCount(ctx context.Context, obj *internal.Show) (int, error) {
	return r.ShowService.GetSeasonCount(ctx, *obj.ID)
}

func (r *showResolver) EpisodeCount(ctx context.Context, obj *internal.Show) (int, error) {
	episodes, err := r.EpisodeService.List(ctx, internal.EpisodesFilter{
		ShowID: obj.ID,
	})
	if err != nil {
		return 0, err
	}
	return len(episodes), nil
}
