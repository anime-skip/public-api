package resolvers

import (
	"context"

	"anime-skip.com/timestamps-service/internal/graphql"
)

func (r *userResolver) AdminOfShows(ctx context.Context, obj *graphql.User) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}
