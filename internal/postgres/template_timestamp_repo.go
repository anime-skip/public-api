package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
)

func findTemplateTimestamps(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) ([]internal.TemplateTimestamp, error) {
	var scanned internal.TemplateTimestamp
	query := sqlbuilder.Select("template_timestamps", map[string]any{
		"template_id":  &scanned.TemplateID,
		"timestamp_id": &scanned.TimestampID,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.TemplateID != nil {
		query.Where("template_id = ?", *filter.TemplateID)
	}
	if filter.TimestampID != nil {
		query.Where("timestamp_id = ?", *filter.TimestampID)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findTemplateTimestamps", err)
	}
	dest := query.ScanDest()
	result := make([]internal.TemplateTimestamp, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findTemplateTimestamps", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findTemplateTimestamp(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) (internal.TemplateTimestamp, error) {
	all, err := findTemplateTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTemplateTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTemplateTimestamp, internal.NewNotFound("TemplateTimestamp", "findTemplateTimestamp")
	}
	return all[0], nil
}

func createTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {

	sql, args := sqlbuilder.Insert("template_timestamps", map[string]any{
		"template_id":  templateTimestamp.TemplateID,
		"timestamp_id": templateTimestamp.TimestampID,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return templateTimestamp, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "TemplateTimestamp with the generated UUID already exists, try again",
			Op:      "createTemplateTimestamp",
			Err:     err,
		}
	} else if err != nil {
		return templateTimestamp, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Template Timestamp",
			Op:      "createTemplateTimestamp",
			Err:     err,
		}
	}
	return templateTimestamp, nil
}

func deleteTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	_, err := tx.ExecContext(ctx, "DELETE FROM template_timestamps WHERE template_id = $1 AND timestamp_id = $2", templateTimestamp.TemplateID, templateTimestamp.TimestampID)
	return templateTimestamp, err
}

func deleteCascadeTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	log.V("Deleting template timestamp (nothing to cascade): template=%v, timestamp=%v", templateTimestamp.TemplateID, templateTimestamp.TimestampID)
	return deleteTemplateTimestamp(ctx, tx, templateTimestamp)
}
