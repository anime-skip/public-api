package postgres

import (
	"context"

	internal "anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	uuid "github.com/gofrs/uuid"
)

func findEpisodeURLs(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) ([]internal.EpisodeURL, error) {
	var scanned internal.EpisodeURL
	query := sqlbuilder.Select("episode_urls", map[string]any{
		"url":                &scanned.URL,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"timestamps_offset":  &scanned.TimestampsOffset,
		"source":             &scanned.Source,
		"episode_id":         &scanned.EpisodeID,
		"duration":           &scanned.Duration,
	})
	if filter.URL != nil {
		query.Where("url = ?", *filter.URL)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("created_at", "ASC")

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findEpisodeURLs", err)
	}
	dest := query.ScanDest()
	result := make([]internal.EpisodeURL, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findEpisodeURLs", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findEpisodeURL(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) (internal.EpisodeURL, error) {
	all, err := findEpisodeURLs(ctx, tx, filter)
	if err != nil {
		return internal.ZeroEpisodeURL, err
	} else if len(all) == 0 {
		return internal.ZeroEpisodeURL, internal.NewNotFound("EpisodeURL", "findEpisodeURL")
	}
	return all[0], nil
}

func createEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, createdBy uuid.UUID) (internal.EpisodeURL, error) {
	episodeURL.CreatedAt = *now()
	episodeURL.CreatedByUserID = &createdBy
	episodeURL.UpdatedAt = *now()
	episodeURL.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("episode_urls", map[string]any{
		"url":                episodeURL.URL,
		"created_at":         episodeURL.CreatedAt,
		"created_by_user_id": episodeURL.CreatedByUserID,
		"updated_at":         episodeURL.UpdatedAt,
		"updated_by_user_id": episodeURL.UpdatedByUserID,
		"timestamps_offset":  episodeURL.TimestampsOffset,
		"source":             episodeURL.Source,
		"episode_id":         episodeURL.EpisodeID,
		"duration":           episodeURL.Duration,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return episodeURL, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "EpisodeURL with the generated UUID already exists, try again",
			Op:      "createEpisodeURL",
			Err:     err,
		}
	} else if err != nil {
		return episodeURL, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Episode URL",
			Op:      "createEpisodeURL",
			Err:     err,
		}
	}
	return episodeURL, nil
}

func updateEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, updatedBy uuid.UUID) (internal.EpisodeURL, error) {
	episodeURL.UpdatedAt = *now()
	episodeURL.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("episode_urls", episodeURL.URL, map[string]any{
		"updated_at":         episodeURL.UpdatedAt,
		"updated_by_user_id": episodeURL.UpdatedByUserID,
		"timestamps_offset":  episodeURL.TimestampsOffset,
		"source":             episodeURL.Source,
		"episode_id":         episodeURL.EpisodeID,
		"duration":           episodeURL.Duration,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return episodeURL, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update EpisodeURL",
			Op:      "updateEpisodeURL",
			Err:     err,
		}
	}

	return episodeURL, nil
}

func deleteEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	_, err := tx.ExecContext(ctx, "DELETE FROM episode_urls WHERE url = $1", episodeURL.URL)
	return episodeURL, err
}

func deleteCascadeEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	log.V("Deleting episode url (nothing to cascade): %s", episodeURL.URL)
	return deleteEpisodeURL(ctx, tx, episodeURL)
}
