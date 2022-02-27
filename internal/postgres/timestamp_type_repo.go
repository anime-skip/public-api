package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func deleteCascadeTimestampType(ctx context.Context, tx internal.Tx, template internal.TimestampType) (internal.Episode, error) {
	panic("deleteCascadeTimestampType not implemented")
}
