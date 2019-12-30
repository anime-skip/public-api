package resolvers

import (
	"context"
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	"github.com/jinzhu/gorm"
)

// Helpers

func showByID(db *gorm.DB, showID string) (*models.Show, error) {
	show, err := repos.FindShowByID(db, showID)
	if err != nil {
		return nil, err
	}
	return mappers.ShowEntityToModel(show), nil
}

// Query Resolvers

type showResolver struct{ *Resolver }

func (r *queryResolver) FindShow(ctx context.Context, showID string) (*models.Show, error) {
	return showByID(r.DB(ctx), showID)
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*models.Show, error) {
	shows, err := repos.SearchShows(r.DB(ctx), *search, *offset, *limit, *sort)
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

func (r *mutationResolver) CreateShow(ctx context.Context, showInput models.InputShow, becomeAdmin bool) (showModel *models.Show, err error) {
	tx := r.DB(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to create show: %+v", r)
			tx.Rollback()
		}
	}()

	show, err := repos.CreateShow(tx, showInput)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	showModel = mappers.ShowEntityToModel(show)
	if !becomeAdmin {
		tx.Commit()
		return showModel, nil
	}

	// Add the Admin relation for this user
	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	newAdmin := models.InputShowAdmin{
		ShowID: show.ID.String(),
		UserID: userID,
	}
	_, err = repos.CreateShowAdmin(tx, newAdmin)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return showModel, nil
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID string, newShow models.InputShow) (*models.Show, error) {
	existingShow, err := repos.FindShowByID(r.DB(ctx), showID)
	if err != nil {
		return nil, err
	}
	updatedShow, err := repos.UpdateShow(r.DB(ctx), newShow, existingShow)

	return mappers.ShowEntityToModel(updatedShow), nil
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID string) (*models.Show, error) {
	db := r.DB(ctx)
	show, err := repos.FindShowByID(db, showID)
	if err != nil {
		return nil, err
	}

	err = repos.DeleteShow(db, show)
	if err != nil {
		return nil, err
	}
	return mappers.ShowEntityToModel(show), nil
}

// Field Resolvers

func (r *showResolver) CreatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *showResolver) Admins(ctx context.Context, obj *models.Show) ([]*models.ShowAdmin, error) {
	return showAdminsByShowID(r.DB(ctx), obj.ID)
}

func (r *showResolver) Episodes(ctx context.Context, obj *models.Show) ([]*models.Episode, error) {
	return nil, fmt.Errorf("not implemented")
}
