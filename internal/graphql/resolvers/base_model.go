package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// getUserById can be used by any created at, updated at, deleted at field resolver as a one liner
func (r *Resolver) getUserById(ctx context.Context, id *uuid.UUID) (*graphql.User, error) {
	user, err := r.UserService.GetUserByID(ctx, internal.GetUserByIDParams{
		UserID: *id,
	})
	if err != nil {
		return nil, err
	}
	gqlUser := mappers.ToGraphqlUser(user)
	return &gqlUser, nil
}
