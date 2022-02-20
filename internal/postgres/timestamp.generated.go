// Code generated by cmd/sqlgen/main.go, DO NOT EDIT.

package postgres

import (
	internal "anime-skip.com/timestamps-service/internal"
	context1 "anime-skip.com/timestamps-service/internal/context"
	"context"
	"fmt"
	uuid "github.com/gofrs/uuid"
	sqlx "github.com/jmoiron/sqlx"
	"time"
)

func getTimestampByID(ctx context.Context, db internal.Database, id uuid.UUID) (internal.Timestamp, error) {
	var timestamp internal.Timestamp
	err := db.GetContext(ctx, &timestamp, "SELECT * FROM timestamps WHERE id=$1", id)
	return timestamp, err
}

func insertTimestampInTx(ctx context.Context, tx *sqlx.Tx, timestamp internal.Timestamp) (internal.Timestamp, error) {
	newTimestamp := timestamp
	auth, err := context1.GetAuthenticationDetails(ctx)
	if err != nil {
		return internal.Timestamp{}, err
	}
	newTimestamp.CreatedAt = time.Now()
	newTimestamp.CreatedByUserID = auth.UserID
	newTimestamp.UpdatedAt = time.Now()
	newTimestamp.UpdatedByUserID = auth.UserID
	newTimestamp.DeletedAt = nil
	newTimestamp.DeletedByUserID = nil
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO timestamps(id, created_at, created_by_user_id, updated_at, updated_by_user_id, deleted_at, deleted_by_user_id, at, source, type_id, episode_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		newTimestamp.ID, newTimestamp.CreatedAt, newTimestamp.CreatedByUserID, newTimestamp.UpdatedAt, newTimestamp.UpdatedByUserID, newTimestamp.DeletedAt, newTimestamp.DeletedByUserID, newTimestamp.At, newTimestamp.Source, newTimestamp.TypeID, newTimestamp.EpisodeID,
	)
	if err != nil {
		return internal.Timestamp{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.Timestamp{}, err
	}
	if changedRows != 1 {
		return internal.Timestamp{}, fmt.Errorf("Inserted %d rows, not 1", changedRows)
	}
	return newTimestamp, err
}

func insertTimestamp(ctx context.Context, db internal.Database, timestamp internal.Timestamp) (internal.Timestamp, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.Timestamp{}, err
	}
	defer tx.Rollback()

	newTimestamp, err := insertTimestampInTx(ctx, tx, timestamp)
	if err != nil {
		return internal.Timestamp{}, err
	}

	tx.Commit()
	return newTimestamp, nil
}

func updateTimestampInTx(ctx context.Context, tx *sqlx.Tx, newTimestamp internal.Timestamp) (internal.Timestamp, error) {
	updatedTimestamp := newTimestamp
	auth, err := context1.GetAuthenticationDetails(ctx)
	if err != nil {
		return internal.Timestamp{}, err
	}
	updatedTimestamp.UpdatedAt = time.Now()
	updatedTimestamp.UpdatedByUserID = auth.UserID
	result, err := tx.ExecContext(
		ctx,
		"UPDATE timestamps SET id=$1, created_at=$2, created_by_user_id=$3, updated_at=$4, updated_by_user_id=$5, deleted_at=$6, deleted_by_user_id=$7, at=$8, source=$9, type_id=$10, episode_id=$11",
		updatedTimestamp.ID, updatedTimestamp.CreatedAt, updatedTimestamp.CreatedByUserID, updatedTimestamp.UpdatedAt, updatedTimestamp.UpdatedByUserID, updatedTimestamp.DeletedAt, updatedTimestamp.DeletedByUserID, updatedTimestamp.At, updatedTimestamp.Source, updatedTimestamp.TypeID, updatedTimestamp.EpisodeID,
	)
	if err != nil {
		return internal.Timestamp{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.Timestamp{}, err
	}
	if changedRows != 1 {
		return internal.Timestamp{}, fmt.Errorf("Updated %d rows, not 1", changedRows)
	}
	return updatedTimestamp, err
}

func updateTimestamp(ctx context.Context, db internal.Database, timestamp internal.Timestamp) (internal.Timestamp, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.Timestamp{}, err
	}
	defer tx.Rollback()

	newTimestamp, err := updateTimestampInTx(ctx, tx, timestamp)
	if err != nil {
		return internal.Timestamp{}, err
	}

	tx.Commit()
	return newTimestamp, nil
}
