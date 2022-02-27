package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func deleteCascadeAPIClient(ctx context.Context, tx internal.Tx, episode internal.APIClient) (internal.APIClient, error) {
	panic("deleteCascadeAPIClient not implemented")
}
