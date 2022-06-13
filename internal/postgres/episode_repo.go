package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findEpisodes(ctx context.Context, tx internal.Tx, filter internal.EpisodesFilter) ([]internal.Episode, error) {
	var scanned internal.Episode
	query := sqlbuilder.Select("episodes", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"name":               &scanned.Name,
		"show_id":            &scanned.ShowID,
		"season":             &scanned.Season,
		"number":             &scanned.Number,
		"absolute_number":    &scanned.AbsoluteNumber,
		"base_duration":      &scanned.BaseDuration,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.ShowID != nil {
		query.Where("show_id = ?", *filter.ShowID)
	}
	if filter.AbsoluteNumber != nil {
		query.Where("absolute_number = ?", *filter.AbsoluteNumber)
	}
	if filter.Number != nil {
		query.Where("number = ?", *filter.Number)
	}
	if filter.Season != nil {
		query.Where("season = ?", *filter.Season)
	}
	if filter.Name != nil {
		query.Where("name = ?", *filter.Name)
	}
	if filter.NameContains != nil {
		query.Where("name ILIKE ?", "%"+*filter.NameContains+"%")
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("name", filter.Sort)

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findEpisodes", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Episode, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findEpisodes", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findEpisode(ctx context.Context, tx internal.Tx, filter internal.EpisodesFilter) (internal.Episode, error) {
	all, err := findEpisodes(ctx, tx, filter)
	if err != nil {
		return internal.ZeroEpisode, err
	} else if len(all) == 0 {
		return internal.ZeroEpisode, internal.NewNotFound("Episode", "findEpisode")
	}
	return all[0], nil
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
	query := `
		SELECT
			id,
			created_at,
			created_by_user_id,
			updated_at,
			updated_by_user_id,
			deleted_at,
			deleted_by_user_id,
			name,
			show_id,
			season,
			"number",
			absolute_number
		FROM episodes
		WHERE id IN (
			SELECT episode_id
			FROM (SELECT DISTINCT ON (episode_id) * FROM timestamps) as episode_ids
			ORDER BY created_at DESC NULLS LAST
			LIMIT $1
			OFFSET $2
		)
		ORDER BY created_at DESC NULLS LAST;
	`
	rows, err := tx.QueryContext(ctx, query, filter.Limit, filter.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			&e.Name,
			&e.ShowID,
			&e.Season,
			&e.Number,
			&e.AbsoluteNumber,
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

func countEpisodeSeasonsByShowID(ctx context.Context, tx internal.Tx, showID uuid.UUID) (int, error) {
	return count(
		ctx,
		tx,
		"SELECT COUNT(*) FROM (SELECT DISTINCT season FROM episodes WHERE show_id = $1 AND deleted_at IS NULL) AS temp",
		showID,
	)
}

func createEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, createdBy uuid.UUID) (internal.Episode, error) {
	id, err := utils.RandomID()
	if err != nil {
		return episode, err
	}
	episode.ID = id
	episode.CreatedAt = *now()
	episode.CreatedByUserID = &createdBy
	episode.UpdatedAt = *now()
	episode.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("episodes", map[string]any{
		"id":                 episode.ID,
		"created_at":         episode.CreatedAt,
		"created_by_user_id": episode.CreatedByUserID,
		"updated_at":         episode.UpdatedAt,
		"updated_by_user_id": episode.UpdatedByUserID,
		"name":               episode.Name,
		"show_id":            episode.ShowID,
		"season":             episode.Season,
		"\"number\"":         episode.Number,
		"absolute_number":    episode.AbsoluteNumber,
		"base_duration":      episode.BaseDuration,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return episode, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "Episode with the generated UUID already exists, try again",
			Op:      "createEpisode",
			Err:     err,
		}
	} else if err != nil {
		return episode, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Episode",
			Op:      "createEpisode",
			Err:     err,
		}
	}
	return episode, nil
}

func updateEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, updatedBy uuid.UUID) (internal.Episode, error) {
	episode.UpdatedAt = *now()
	episode.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("episodes", episode.ID, map[string]any{
		"updated_at":         episode.UpdatedAt,
		"updated_by_user_id": episode.UpdatedByUserID,
		"deleted_at":         episode.DeletedAt,
		"deleted_by_user_id": episode.DeletedByUserID,
		"name":               episode.Name,
		"show_id":            episode.ShowID,
		"season":             episode.Season,
		"\"number\"":         episode.Number,
		"absolute_number":    episode.AbsoluteNumber,
		"base_duration":      episode.BaseDuration,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return episode, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Episode",
			Op:      "updateEpisode",
			Err:     err,
		}
	}

	return episode, nil
}

func deleteEpisode(ctx context.Context, tx internal.Tx, episode internal.Episode, deletedBy uuid.UUID) (internal.Episode, error) {
	episode.DeletedByUserID = &deletedBy
	episode.DeletedAt = now()
	return updateEpisode(ctx, tx, episode, deletedBy)
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
