package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/jinzhu/gorm"
)

// Helpers

func showAdminByID(db *gorm.DB, showAdminID string) (*models.ShowAdmin, error) {
	user, err := repos.FindShowAdminByID(db, showAdminID)
	if err != nil {
		return nil, err
	}
	return mappers.ShowAdminEntityToModel(user), err
}

func showAdminsByUserID(db *gorm.DB, userID string) ([]*models.ShowAdmin, error) {
	admins, err := repos.FindShowAdminsByUserID(db, userID)
	if err != nil {
		return nil, err
	}

	adminModels := make([]*models.ShowAdmin, len(admins))
	for index, entity := range admins {
		adminModels[index] = mappers.ShowAdminEntityToModel(entity)
	}
	return adminModels, nil
}

func showAdminsByShowID(db *gorm.DB, showID string) ([]*models.ShowAdmin, error) {
	admins, err := repos.FindShowAdminsByShowID(db, showID)
	if err != nil {
		return nil, err
	}

	adminModels := make([]*models.ShowAdmin, len(admins))
	for index, entity := range admins {
		adminModels[index] = mappers.ShowAdminEntityToModel(entity)
	}
	return adminModels, nil
}

// Query Resolvers

type showAdminResolver struct{ *Resolver }

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID string) (*models.ShowAdmin, error) {
	return showAdminByID(r.DB(ctx), showAdminID)
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID string) ([]*models.ShowAdmin, error) {
	return showAdminsByShowID(r.DB(ctx), showID)
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID string) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(r.DB(ctx), userID)
}

// Mutation Resolvers

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput models.InputShowAdmin) (*models.ShowAdmin, error) {
	showAdmin, err := repos.CreateShowAdmin(r.DB(ctx), showAdminInput)
	if err != nil {
		return nil, err
	}
	return mappers.ShowAdminEntityToModel(showAdmin), nil
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID string) (*models.ShowAdmin, error) {
	showAdmin, err := repos.FindShowAdminByID(r.DB(ctx), showAdminID)
	if err != nil {
		return nil, err
	}

	err = repos.DeleteShowAdmin(r.DB(ctx), showAdmin)
	if err != nil {
		return nil, err
	}

	return mappers.ShowAdminEntityToModel(showAdmin), nil
}

// Field Resolvers

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(r.DB(ctx), obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UpdatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return deletedUserByID(r.DB(ctx), obj.DeletedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *models.ShowAdmin) (*models.Show, error) {
	return showByID(r.DB(ctx), obj.ShowID)
}

func (r *showAdminResolver) User(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(r.DB(ctx), obj.UserID)
}
