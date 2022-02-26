package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	uuid "github.com/gofrs/uuid"
)

func getEpisodeSeasonCountByShowID(ctx context.Context, db internal.Database, id uuid.UUID) (int, error) {
	row, err := db.QueryContext(ctx, "SELECT DISTINCT count(season) FROM episodes WHERE show_id=$1", id)
	if err != nil {
		return 0, err
	}

	var count int
	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func getRecentlyAddedEpisodes(ctx context.Context, db internal.Database, params internal.GetRecentlyAddedParams) ([]internal.Episode, error) {
	panic("not implemented")
}
