package resolvers

import (
	ctx "context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/context"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
)

func (r *queryResolver) Account(ctx ctx.Context) (*graphql.Account, error) {
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
	account := mappers.InternalUserToAccount(user)
	return &account, nil
}
