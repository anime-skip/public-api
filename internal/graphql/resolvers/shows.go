package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Helpers

func showByID(ctx context.Context, orm *database.ORM, showID string) (*models.Show, error) {
	show, err := repos.FindShowByID(ctx, orm, showID)
	return mappers.ShowEntityToModel(show), err
}

// Query Resolvers

type showResolver struct{ *Resolver }

func (r *queryResolver) FindShow(ctx context.Context, showID string) (*models.Show, error) {
	return showByID(ctx, r.ORM(ctx), showID)
}

func (r *queryResolver) FindShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*models.Show, error) {
	shows, err := repos.FindShows(ctx, r.ORM(ctx), *search, *offset, *limit, *sort)
	if err != nil {
		return nil, err
	}
	showModels := make([]*models.Show, len(shows))
	for index, entity := range shows {
		showModels[index] = mappers.ShowEntityToModel(entity)
	}
	return showModels, nil
}

// Mutation Resolvers

func (r *mutationResolver) CreateShow(ctx context.Context, showInput models.InputShow) (*models.Show, error) {
	show, err := repos.CreateShow(ctx, r.ORM(ctx), showInput)
	if err != nil {
		return nil, err
	}
	return mappers.ShowEntityToModel(show), nil
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID string, newShow models.InputShow) (*models.Show, error) {
	existingShow, err := repos.FindShowByID(ctx, r.ORM(ctx), showID)
	if err != nil {
		return nil, err
	}
	updatedShow, err := repos.UpdateShow(ctx, r.ORM(ctx), newShow, existingShow)

	return mappers.ShowEntityToModel(updatedShow), nil
}

// Field Resolvers

func (r *showResolver) CreatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return deletedUserByID(ctx, r.ORM(ctx), obj.DeletedByUserID)
}

func (r *showResolver) Admins(ctx context.Context, obj *models.Show) ([]*models.ShowAdmin, error) {
	return showAdminsByShowID(ctx, r.ORM(ctx), obj.ID)
}

func (r *showResolver) Episodes(ctx context.Context, obj *models.Show) ([]*models.Episode, error) {
	return nil, fmt.Errorf("not implemented")
}
