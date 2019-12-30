package resolvers

import (
	"context"

	"github.com/aklinker1/anime-skip-backend/internal/database/mappers"
	"github.com/aklinker1/anime-skip-backend/internal/database/repos"
	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
)

// Helpers

// Query Resolvers

type myUserResolver struct{ *Resolver }

func (r *queryResolver) MyUser(ctx context.Context) (*models.MyUser, error) {
	userID, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := repos.FindUserByID(r.DB(ctx), userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserEntityToMyUserModel(user), nil
}

// Mutation Resolvers

// Field Resolvers

func (r *myUserResolver) AdminOfShows(ctx context.Context, obj *models.MyUser) ([]*models.ShowAdmin, error) {
	return showAdminsByUserID(r.DB(ctx), obj.ID)
}

func (r *myUserResolver) Preferences(ctx context.Context, obj *models.MyUser) (*models.Preferences, error) {
	preferences, err := repos.FindPreferencesByUserID(r.DB(ctx), obj.ID)
	return mappers.PreferencesEntityToModel(preferences), err
}
