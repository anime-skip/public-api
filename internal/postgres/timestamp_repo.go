package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findTimestamps(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) ([]internal.Timestamp, error) {
	var scanned internal.Timestamp
	query := sqlbuilder.Select("timestamps", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"episode_id":         &scanned.EpisodeID,
		"at":                 &scanned.At,
		"type_id":            &scanned.TypeID,
		"source":             &scanned.Source,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.EpisodeID != nil {
		query.Where("episode_id = ?", *filter.EpisodeID)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("at", "ASC")

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findTimestamps", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Timestamp, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findTimestamps", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findTimestamp(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) (internal.Timestamp, error) {
	all, err := findTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTimestamp, internal.NewNotFound("Timestamp", "findTimestamp")
	}
	return all[0], nil
}

func createTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, createdBy uuid.UUID) (internal.Timestamp, error) {
	id, err := utils.RandomID()
	if err != nil {
		return timestamp, err
	}
	timestamp.ID = id
	timestamp.CreatedAt = *now()
	timestamp.CreatedByUserID = &createdBy
	timestamp.UpdatedAt = *now()
	timestamp.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("timestamps", map[string]any{
		"id":                 timestamp.ID,
		"created_at":         timestamp.CreatedAt,
		"created_by_user_id": timestamp.CreatedByUserID,
		"updated_at":         timestamp.UpdatedAt,
		"updated_by_user_id": timestamp.UpdatedByUserID,
		"episode_id":         timestamp.EpisodeID,
		"at":                 timestamp.At,
		"type_id":            timestamp.TypeID,
		"source":             timestamp.Source,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return timestamp, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "Timestamp with the generated UUID already exists, try again",
			Op:      "createTimestamp",
			Err:     err,
		}
	} else if err != nil {
		return timestamp, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Timestamp",
			Op:      "createTimestamp",
			Err:     err,
		}
	}
	return timestamp, nil
}

func updateTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, updatedBy uuid.UUID) (internal.Timestamp, error) {
	timestamp.UpdatedAt = *now()
	timestamp.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("timestamps", timestamp.ID, map[string]any{
		"updated_at":         timestamp.UpdatedAt,
		"updated_by_user_id": timestamp.UpdatedByUserID,
		"deleted_at":         timestamp.DeletedAt,
		"deleted_by_user_id": timestamp.DeletedByUserID,
		"episode_id":         timestamp.EpisodeID,
		"at":                 timestamp.At,
		"type_id":            timestamp.TypeID,
		"source":             timestamp.Source,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return timestamp, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Timestamp",
			Op:      "updateTimestamp",
			Err:     err,
		}
	}

	return timestamp, nil
}

func deleteTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, deletedBy uuid.UUID) (internal.Timestamp, error) {
	timestamp.DeletedByUserID = &deletedBy
	timestamp.DeletedAt = now()
	return updateTimestamp(ctx, tx, timestamp, deletedBy)
}

func deleteCascadeTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, deletedBy uuid.UUID) (internal.Timestamp, error) {
	log.V("Deleting timestamp: %v", timestamp.ID)
	deletedTimestamp, err := deleteTimestamp(ctx, tx, timestamp, deletedBy)
	if err != nil {
		return internal.Timestamp{}, err
	}

	log.V("Deleting timestamp template timestamps")
	templateTimestamps, err := findTemplateTimestamps(ctx, tx, internal.TemplateTimestampsFilter{
		TimestampID: timestamp.ID,
	})
	if err != nil {
		return internal.ZeroTimestamp, err
	}
	for _, templateTimestamp := range templateTimestamps {
		_, err := deleteCascadeTemplateTimestamp(ctx, tx, templateTimestamp)
		if err != nil {
			return internal.ZeroTimestamp, err
		}
	}

	log.V("Done deleting timestamp: %v", timestamp.ID)
	return deletedTimestamp, nil
}
