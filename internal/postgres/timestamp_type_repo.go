package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findTimestampTypes(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) ([]internal.TimestampType, error) {
	var scanned internal.TimestampType
	query := sqlbuilder.Select("timestamp_types", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"name":               &scanned.Name,
		"description":        &scanned.Description,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findTimestampTypes", err)
	}
	dest := query.ScanDest()
	result := make([]internal.TimestampType, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findTimestampTypes", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findTimestampType(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) (internal.TimestampType, error) {
	all, err := findTimestampTypes(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestampType, err
	} else if len(all) == 0 {
		return internal.ZeroTimestampType, internal.NewNotFound("TimestampType", "findTimestampType")
	}
	return all[0], nil
}

func createTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, createdBy uuid.UUID) (internal.TimestampType, error) {
	id, err := utils.RandomID()
	if err != nil {
		return timestampType, err
	}
	timestampType.ID = id
	timestampType.CreatedAt = *now()
	timestampType.CreatedByUserID = &createdBy
	timestampType.UpdatedAt = *now()
	timestampType.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("timestamp_types", map[string]any{
		"id":                 timestampType.ID,
		"created_at":         timestampType.CreatedAt,
		"created_by_user_id": timestampType.CreatedByUserID,
		"updated_at":         timestampType.UpdatedAt,
		"updated_by_user_id": timestampType.UpdatedByUserID,
		"deleted_at":         timestampType.DeletedAt,
		"deleted_by_user_id": timestampType.DeletedByUserID,
		"name":               timestampType.Name,
		"description":        timestampType.Description,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return timestampType, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "TimestampType with the generated UUID already exists, try again",
			Op:      "createTimestampType",
			Err:     err,
		}
	} else if err != nil {
		return timestampType, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Timestamp Type",
			Op:      "createTimestampType",
			Err:     err,
		}
	}
	return timestampType, nil
}

func updateTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, updatedBy uuid.UUID) (internal.TimestampType, error) {
	timestampType.UpdatedAt = *now()
	timestampType.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("timestamp_types", timestampType.ID, map[string]any{
		"updated_at":         timestampType.UpdatedAt,
		"updated_by_user_id": timestampType.UpdatedByUserID,
		"deleted_at":         timestampType.DeletedAt,
		"deleted_by_user_id": timestampType.DeletedByUserID,
		"name":               timestampType.Name,
		"description":        timestampType.Description,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return timestampType, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Timestamp Type",
			Op:      "updateTimestampType",
			Err:     err,
		}
	}

	return timestampType, nil
}

func deleteTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, deletedBy uuid.UUID) (internal.TimestampType, error) {
	timestampType.DeletedByUserID = &deletedBy
	timestampType.DeletedAt = now()
	return updateTimestampType(ctx, tx, timestampType, deletedBy)
}

func deleteCascadeTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, deletedBy uuid.UUID) (internal.TimestampType, error) {
	// Since timestamp types are soft delete, we don't need to update existing timestamps to a
	// different type, they'll stay the same type, that type just won't be returned by
	// allTimestampTypes anymore
	log.V("Deleting timestamp type (nothing to cascade): %v", timestampType.ID)
	return deleteTimestampType(ctx, tx, timestampType, deletedBy)
}
