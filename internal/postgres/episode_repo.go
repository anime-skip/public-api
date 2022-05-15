package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	uuid "github.com/gofrs/uuid"
)

func getEpisodeSeasonCountByShowID(ctx context.Context, tx internal.Tx, id uuid.UUID) (int, error) {
	return count(
		ctx,
		tx,
		"SELECT DISTINCT count(season) FROM episodes WHERE show_id=$1 AND deleted_at IS NULL",
		id,
	)
}

var episodeColumns = strings.Join([]string{
	"id",
	"created_at",
	"created_by_user_id",
	"updated_at",
	"updated_by_user_id",
	"deleted_at",
	"deleted_by_user_id",
	"season",
	"\"number\"",
	"absolute_number",
	"name",
	"show_id",
}, ", ")

func scanEpisodeRows(rows *sql.Rows) ([]internal.Episode, error) {
	result := make([]internal.Episode, 0)
	for rows.Next() {
		var e internal.Episode
		err := rows.Scan(
			&e.ID,
			&e.CreatedAt,
			&e.CreatedByUserID,
			&e.UpdatedAt,
			&e.UpdatedByUserID,
			&e.DeletedAt,
			&e.DeletedByUserID,
			&e.Season,
			&e.Number,
			&e.AbsoluteNumber,
			&e.Name,
			&e.ShowID,
		)
		if err != nil {
			return nil, &internal.Error{
				Code:    internal.EINTERNAL,
				Message: "Failed to load API clients",
				Op:      "findAPIClients",
				Err:     err,
			}
		}
		result = append(result, e)
	}
	return result, rows.Err()
}

func findRecentlyAddedEpisodes(ctx context.Context, tx internal.Tx, filter internal.RecentlyAddedEpisodesFilter) ([]internal.Episode, error) {
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
	query := fmt.Sprintf(
		`
			SELECT %s FROM episodes
			WHERE id IN (
				SELECT episode_id
				FROM (SELECT DISTINCT ON (episode_id) * FROM timestamps) as episode_ids
				ORDER BY created_at DESC NULLS LAST
				LIMIT $1
				OFFSET $2
			)
			ORDER BY created_at DESC NULLS LAST;
		`,
		episodeColumns,
	)
	rows, err := tx.QueryContext(ctx, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanEpisodeRows(rows)
}

func deleteCascadeEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, deletedBy uuid.UUID) (internal.Episode, error) {
	log.V("Deleting episode: %v", episode.ID)
	deletedEpisode, err := deleteEpisode(ctx, tx, episode, deletedBy)
	if err != nil {
		return internal.Episode{}, err
	}

	log.V("Deleting episode templates")
	templates, err := findTemplates(ctx, tx, internal.TemplatesFilter{
		SourceEpisodeID: episode.ID,
	})
	if err != nil {
		return internal.Episode{}, err
	}
	for _, template := range templates {
		_, err = deleteCascadeTemplate(ctx, tx, template, deletedBy)
		if !internal.IsNotFound(err) {
			return internal.Episode{}, err
		}
	}

	log.V("Deleting episode timestamps")
	timestamps, err := findTimestamps(ctx, tx, internal.TimestampsFilter{
		EpisodeID: episode.ID,
	})
	if err != nil {
		return internal.Episode{}, err
	}
	for _, timestamp := range timestamps {
		_, err = deleteCascadeTimestamp(ctx, tx, timestamp, deletedBy)
		if !internal.IsNotFound(err) {
			return internal.Episode{}, err
		}
	}

	log.V("Deleting episode urls")
	urls, err := findEpisodeURLs(ctx, tx, internal.EpisodeURLsFilter{
		EpisodeID: episode.ID,
	})
	if err != nil {
		return internal.Episode{}, err
	}
	for _, url := range urls {
		_, err := deleteCascadeEpisodeURL(ctx, tx, url)
		if err != nil {
			return internal.Episode{}, err
		}
	}

	log.V("Done deleting episode: %v", episode.ID)
	return deletedEpisode, nil
}
