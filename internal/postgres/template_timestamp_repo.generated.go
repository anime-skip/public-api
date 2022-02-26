// Code generated by cmd/sqlgen/main.go, DO NOT EDIT.

package postgres

import (
	internal "anime-skip.com/timestamps-service/internal"
	"context"
	"fmt"
	uuid "github.com/gofrs/uuid"
)

func getTemplateTimestampsByTemplateID(ctx context.Context, db internal.Database, templateID uuid.UUID) ([]internal.TemplateTimestamp, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM template_timestamps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	templateTimestamps := []internal.TemplateTimestamp{}
	for rows.Next() {
		var templateTimestamp internal.TemplateTimestamp
		err = rows.StructScan(&templateTimestamp)
		if err != nil {
			return nil, err
		}
		templateTimestamps = append(templateTimestamps, templateTimestamp)
	}
	return templateTimestamps, nil
}

func getTemplateTimestampsByTimestampID(ctx context.Context, db internal.Database, timestampID uuid.UUID) ([]internal.TemplateTimestamp, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM template_timestamps")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	templateTimestamps := []internal.TemplateTimestamp{}
	for rows.Next() {
		var templateTimestamp internal.TemplateTimestamp
		err = rows.StructScan(&templateTimestamp)
		if err != nil {
			return nil, err
		}
		templateTimestamps = append(templateTimestamps, templateTimestamp)
	}
	return templateTimestamps, nil
}

func insertTemplateTimestampInTx(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	newTemplateTimestamp := templateTimestamp
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO template_timestamps(template_id, timestamp_id) VALUES ($1, $2)",
		newTemplateTimestamp.TemplateID, newTemplateTimestamp.TimestampID,
	)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	if changedRows != 1 {
		return internal.TemplateTimestamp{}, fmt.Errorf("Inserted more than 1 row (%d)", changedRows)
	}
	return newTemplateTimestamp, err
}

func insertTemplateTimestamp(ctx context.Context, db internal.Database, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	defer tx.Rollback()

	result, err := insertTemplateTimestampInTx(ctx, tx, templateTimestamp)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}

	tx.Commit()
	return result, nil
}

func deleteTemplateTimestampInTx(ctx context.Context, tx internal.Tx, newTemplateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	deletedTemplateTimestamp := newTemplateTimestamp
	result, err := tx.ExecContext(ctx, "DELETE FROM template_timestamps WHERE template_id=$1 AND timestamp_id=$2", deletedTemplateTimestamp.TemplateID, deletedTemplateTimestamp.TimestampID)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	if changedRows != 1 {
		return internal.TemplateTimestamp{}, fmt.Errorf("Deleted more than 1 row (%d)", changedRows)
	}
	return deletedTemplateTimestamp, err
}

func deleteTemplateTimestamp(ctx context.Context, db internal.Database, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}
	defer tx.Rollback()

	result, err := deleteTemplateTimestampInTx(ctx, tx, templateTimestamp)
	if err != nil {
		return internal.TemplateTimestamp{}, err
	}

	tx.Commit()
	return result, nil
}
