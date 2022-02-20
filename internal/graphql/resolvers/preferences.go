package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *preferencesResolver) User(ctx context.Context, obj *graphql.Preferences) (*graphql.User, error) {
	return r.getUserById(ctx, obj.UserID)
}
