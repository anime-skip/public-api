package resolvers

import (
	go_context "context"

	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
)

func (r *queryResolver) Account(ctx go_context.Context) (*graphql.Account, error) {
	auth, err := context.GetAuthenticationDetails(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.UserService.GetByID(ctx, auth.UserID)
	if err != nil {
		return nil, err
	}
	account := mappers.ToGraphqlAccount(user)
	return &account, nil
}

func (r *accountResolver) Preferences(ctx go_context.Context, obj *graphql.Account) (*graphql.Preferences, error) {
	prefs, err := r.PreferencesService.GetByUserID(ctx, *obj.ID)
	if err != nil {
		return nil, err
	}
	gqlPrefs := mappers.ToGraphqlPreferences(prefs)
	return &gqlPrefs, nil
}

func (r *accountResolver) AdminOfShows(ctx go_context.Context, obj *graphql.Account) ([]*graphql.ShowAdmin, error) {
	panic("accountResolver.AdminOfShows not implemented")
}
