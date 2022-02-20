package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
	"anime-skip.com/timestamps-service/internal/graphql/mappers"
	"github.com/gofrs/uuid"
)

// getUserById can be used by any created_by_user_id, updated_by_user_id, deleted_by_user_id field resolver as a one liner
func (r *Resolver) getUserById(ctx context.Context, id *uuid.UUID) (*graphql.User, error) {
	// Handle deletedBy case
	if id == nil {
		return nil, nil
	}

	// return user for created_by_user_id and updated_by_user_id
	user, err := r.UserService.GetByID(ctx, *id)
	if err != nil {
		return nil, err
	}
	gqlUser := mappers.ToGraphqlUser(user)
	return &gqlUser, nil
}
