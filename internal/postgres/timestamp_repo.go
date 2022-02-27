package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func deleteCascadeTimestamp(ctx context.Context, tx internal.Tx, template internal.Timestamp) (internal.Episode, error) {
	panic("deleteCascadeTimestamp not implemented")
}
