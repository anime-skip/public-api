package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal/graphql"
)

// Helpers

// Mutations

// Queries

// Fields

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *graphql.ThirdPartyTimestamp) (*graphql.TimestampType, error) {
	return r.getTimestampTypeByID(ctx, obj.TypeID)
}
