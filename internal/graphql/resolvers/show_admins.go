package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
)

// Helpers

func showAdminByID(ctx context.Context, orm *database.ORM, showAdminID string) (*models.ShowAdmin, error) {
	user, err := repos.FindShowAdminByID(ctx, orm, showAdminID)
	return mappers.ShowAdminEntityToModel(user), err
}

func showAdminsByUserID(ctx context.Context, orm *database.ORM, userID string) ([]*models.ShowAdmin, error) {
	admins, err := repos.FindShowAdminsByUserID(ctx, orm, userID)
	if err != nil {
		return nil, err
	}

	adminModels := make([]*models.ShowAdmin, len(admins))
	for index, entity := range admins {
		adminModels[index] = mappers.ShowAdminEntityToModel(entity)
	}
	return adminModels, nil
}

func showAdminsByShowID(ctx context.Context, orm *database.ORM, showID string) ([]*models.ShowAdmin, error) {
	admins, err := repos.FindShowAdminsByShowID(ctx, orm, showID)
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

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID string) ([]*models.ShowAdmin, error) {
	return showAdminsByShowID(ctx, r.ORM(ctx), showID)
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID string) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(ctx, r.ORM(ctx), userID)
}

// Mutation Resolvers

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput models.InputShowAdmin) (*models.ShowAdmin, error) {
	showAdmin, err := repos.CreateShowAdmin(ctx, r.ORM(ctx), showAdminInput)
	if err != nil {
		return nil, err
	}
	return mappers.ShowAdminEntityToModel(showAdmin), nil
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID string) (*models.ShowAdmin, error) {
	showAdmin, err := repos.FindShowAdminByID(ctx, r.ORM(ctx), showAdminID)
	if err != nil {
		return nil, err
	}
	showAdminModel := mappers.ShowAdminEntityToModel(showAdmin)

	err = repos.DeleteShowAdmin(ctx, r.ORM(ctx), showAdmin)
	if err != nil {
		return nil, err
	}

	return showAdminModel, nil
}

// Field Resolvers

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.CreatedByUserID)
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UpdatedByUserID)
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return deletedUserByID(ctx, r.ORM(ctx), obj.DeletedByUserID)
}

func (r *showAdminResolver) Show(ctx context.Context, obj *models.ShowAdmin) (*models.Show, error) {
	return showByID(ctx, r.ORM(ctx), obj.ShowID)
}

func (r *showAdminResolver) User(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
	return userByID(ctx, r.ORM(ctx), obj.UserID)
}
