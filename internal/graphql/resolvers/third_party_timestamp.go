package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/graphql"
)

// Helpers

// Mutations

// Queries

// Fields

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *graphql.ThirdPartyTimestamp) (*graphql.TimestampType, error) {
	panic(errors.NewPanicedError("thirdPartyTimestampResolver.Type not implemented"))
}
