package resolvers

import (
	"context"

	"anime-skip.com/public-api/internal"
)

// Helpers

// Mutations

// Queries

// Fields

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *internal.ThirdPartyTimestamp) (*internal.TimestampType, error) {
	return r.getTimestampTypeByID(ctx, obj.TypeID)
}
