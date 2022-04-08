package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/graphql"
	"anime-skip.com/public-api/internal/graphql/mappers"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
	"github.com/gofrs/uuid"
)

// Helpers

func (r *Resolver) getShowById(ctx context.Context, id *uuid.UUID) (*graphql.Show, error) {
	if id == nil {
		return nil, nil
	}
	internalShow, err := r.ShowService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	show := mappers.ToGraphqlShow(internalShow)
	return &show, nil
}

// Mutations

func (r *mutationResolver) CreateShow(ctx context.Context, showInput graphql.InputShow, becomeAdmin bool) (*graphql.Show, error) {
	internalInput := internal.Show{
		BaseEntity: internal.BaseEntity{
			ID: utils.RandomID(),
		},
	}
	mappers.ApplyGraphqlInputShow(showInput, &internalInput)

	created, err := r.ShowService.Create(ctx, internalInput)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlShow(created)
	return &result, nil
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID *uuid.UUID, newShow graphql.InputShow) (*graphql.Show, error) {
	log.V("Updating: %v", showID)
	existing, err := r.ShowService.GetByID(ctx, *showID)
	if err != nil {
		return nil, err
	}
	mappers.ApplyGraphqlInputShow(newShow, &existing)
	log.V("Updating to %+v", existing)
	created, err := r.ShowService.Update(ctx, existing)
	if err != nil {
		log.V("Failed to update: %v", err)
		return nil, err
	}

	result := mappers.ToGraphqlShow(created)
	return &result, nil
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	deleted, err := r.ShowService.Delete(ctx, *showID)
	if err != nil {
		return nil, err
	}

	result := mappers.ToGraphqlShow(deleted)
	return &result, nil
}

// Queries

func (r *queryResolver) FindShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	return r.getShowById(ctx, showID)
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*graphql.Show, error) {
	panic(errors.NewPanicedError("queryResolver.SearchShows not implemented"))
}

// Fields

func (r *showResolver) CreatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	return r.getUserById(ctx, obj.DeletedByUserID)
}

func (r *showResolver) Admins(ctx context.Context, obj *graphql.Show) ([]*graphql.ShowAdmin, error) {
	return r.getShowAdminsByShowId(ctx, obj.ID)
}

func (r *showResolver) Episodes(ctx context.Context, obj *graphql.Show) ([]*graphql.Episode, error) {
	return r.getEpisodesByShowID(ctx, obj.ID)
}

func (r *showResolver) Templates(ctx context.Context, obj *graphql.Show) ([]*graphql.Template, error) {
	return r.getTemplatesByShowID(ctx, obj.ID)
}

func (r *showResolver) SeasonCount(ctx context.Context, obj *graphql.Show) (int, error) {
	return r.ShowService.GetSeasonCount(ctx, *obj.ID)
}

func (r *showResolver) EpisodeCount(ctx context.Context, obj *graphql.Show) (int, error) {
	episodes, err := r.EpisodeService.GetByShowID(ctx, *obj.ID)
	if err != nil {
		return 0, err
	}
	return len(episodes), nil
}
