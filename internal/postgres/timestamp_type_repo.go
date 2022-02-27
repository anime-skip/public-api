package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

func deleteCascadeTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType) (internal.TimestampType, error) {
	// Since timestamp types are soft delete, we don't need to update existing timestamps to a
	// different type, they'll stay the same type, that type just won't be returned by
	// allTimestampTypes anymore
	log.V("Deleting timestamp type (nothing to cascade): %v", timestampType.ID)
	return deleteTimestampTypeInTx(ctx, tx, timestampType)
}
