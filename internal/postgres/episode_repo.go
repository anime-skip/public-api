package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func getRecentlyAddedEpisodes(ctx context.Context, db internal.Database, params internal.GetRecentlyAddedParams) ([]internal.Episode, error) {
	panic("not implemented")
}
