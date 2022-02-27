package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func deleteCascadeShow(ctx context.Context, tx internal.Tx, episode internal.Show) (internal.Show, error) {
	panic("deleteCascadeShow not implemented")
}
