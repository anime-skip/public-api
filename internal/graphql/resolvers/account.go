package resolvers

import (
	go_context "context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
)

func (r *queryResolver) Account(ctx go_context.Context) (*graphql.Account, error) {
	auth, err := context.GetAuthenticationDetails(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.UserService.GetUserByID(ctx, internal.GetUserByIDParams{
		UserID: auth.UserID,
	})
	if err != nil {
		return nil, err
	}
	account := mappers.ToGraphqlAccount(user)
	return &account, nil
}

func (r *accountResolver) Preferences(ctx go_context.Context, obj *graphql.Account) (*graphql.Preferences, error) {
	prefs, err := r.PreferencesService.GetPreferencesByUserID(ctx, internal.GetPreferencesByUserIDParams{
		UserID: *obj.ID,
	})
	if err != nil {
		return nil, err
	}
	gqlPrefs := mappers.ToGraphqlPreferences(prefs)
	return &gqlPrefs, nil
}

func (r *accountResolver) AdminOfShows(ctx go_context.Context, obj *graphql.Account) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}
