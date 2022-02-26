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
	// What is this query?
	// 1. Grab timestamps with distinct episode ids
	// 2. Grab just the episode_id the timestamp belongs to and sort them by newest first (this is
	//    where we apply pagination)
	// 3. Select all the episodes with those ids, making sure to sort them again since the
	//    timestamps' created_at can be in a different order than the episodes' created_at
	//
	// This isn't perfect (episode could be missing or slightly out of order), but it beats what the
	// the old query - 13ms vs 10s
	// https://github.com/anime-skip/backend/blob/33fd2b842bc847bed67c9b1f9283e78b710cd1f5/internal/database/repos/episodes.go#L137-L153
	query := `
		SELECT * FROM episodes
		WHERE id IN (
			SELECT episode_id
			FROM (SELECT DISTINCT ON (episode_id) * FROM timestamps) as episode_ids
			ORDER BY created_at DESC NULLS LAST
			LIMIT $1
			OFFSET $2
		)
		ORDER BY created_at DESC NULLS LAST;
	`
	rows, err := db.QueryxContext(ctx, query, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	episodes := []internal.Episode{}
	for rows.Next() {
		var episode internal.Episode
		err = rows.StructScan(&episode)
		if err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}

	return episodes, nil
}
