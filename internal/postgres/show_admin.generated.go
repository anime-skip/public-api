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

func getShowAdminByID(ctx context.Context, db internal.Database, id uuid.UUID) (internal.ShowAdmin, error) {
	var showAdmin internal.ShowAdmin
	err := db.GetContext(ctx, &showAdmin, "SELECT * FROM show_admins WHERE id=$1", id)
	return showAdmin, err
}

func getShowAdminsByShowID(ctx context.Context, db internal.Database, showID uuid.UUID) ([]internal.ShowAdmin, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM show_admins WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	showAdmins := []internal.ShowAdmin{}
	for rows.Next() {
		var showAdmin internal.ShowAdmin
		err = rows.StructScan(&showAdmin)
		if err != nil {
			return nil, err
		}
		showAdmins = append(showAdmins, showAdmin)
	}
	return showAdmins, nil
}

func getUnscopedShowAdminsByShowID(ctx context.Context, db internal.Database, showID uuid.UUID) ([]internal.ShowAdmin, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM show_admins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	showAdmins := []internal.ShowAdmin{}
	for rows.Next() {
		var showAdmin internal.ShowAdmin
		err = rows.StructScan(&showAdmin)
		if err != nil {
			return nil, err
		}
		showAdmins = append(showAdmins, showAdmin)
	}
	return showAdmins, nil
}

func getShowAdminsByUserID(ctx context.Context, db internal.Database, userID uuid.UUID) ([]internal.ShowAdmin, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM show_admins WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	showAdmins := []internal.ShowAdmin{}
	for rows.Next() {
		var showAdmin internal.ShowAdmin
		err = rows.StructScan(&showAdmin)
		if err != nil {
			return nil, err
		}
		showAdmins = append(showAdmins, showAdmin)
	}
	return showAdmins, nil
}

func getUnscopedShowAdminsByUserID(ctx context.Context, db internal.Database, userID uuid.UUID) ([]internal.ShowAdmin, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM show_admins")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	showAdmins := []internal.ShowAdmin{}
	for rows.Next() {
		var showAdmin internal.ShowAdmin
		err = rows.StructScan(&showAdmin)
		if err != nil {
			return nil, err
		}
		showAdmins = append(showAdmins, showAdmin)
	}
	return showAdmins, nil
}

func insertShowAdminInTx(ctx context.Context, tx *sqlx.Tx, showAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	newShowAdmin := showAdmin
	auth, err := context1.GetAuthenticationDetails(ctx)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	newShowAdmin.CreatedAt = time.Now()
	newShowAdmin.CreatedByUserID = auth.UserID
	newShowAdmin.UpdatedAt = time.Now()
	newShowAdmin.UpdatedByUserID = auth.UserID
	newShowAdmin.DeletedAt = nil
	newShowAdmin.DeletedByUserID = nil
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO show_admins(id, created_at, created_by_user_id, updated_at, updated_by_user_id, deleted_at, deleted_by_user_id, show_id, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		newShowAdmin.ID, newShowAdmin.CreatedAt, newShowAdmin.CreatedByUserID, newShowAdmin.UpdatedAt, newShowAdmin.UpdatedByUserID, newShowAdmin.DeletedAt, newShowAdmin.DeletedByUserID, newShowAdmin.ShowID, newShowAdmin.UserID,
	)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	if changedRows != 1 {
		return internal.ShowAdmin{}, fmt.Errorf("Inserted more than 1 row (%d)", changedRows)
	}
	return newShowAdmin, err
}

func insertShowAdmin(ctx context.Context, db internal.Database, showAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	defer tx.Rollback()

	result, err := insertShowAdminInTx(ctx, tx, showAdmin)
	if err != nil {
		return internal.ShowAdmin{}, err
	}

	tx.Commit()
	return result, nil
}

func updateShowAdminInTx(ctx context.Context, tx *sqlx.Tx, newShowAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	updatedShowAdmin := newShowAdmin
	auth, err := context1.GetAuthenticationDetails(ctx)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	updatedShowAdmin.UpdatedAt = time.Now()
	updatedShowAdmin.UpdatedByUserID = auth.UserID
	result, err := tx.ExecContext(
		ctx,
		"UPDATE show_admins SET id=$1, created_at=$2, created_by_user_id=$3, updated_at=$4, updated_by_user_id=$5, deleted_at=$6, deleted_by_user_id=$7, show_id=$8, user_id=$9",
		updatedShowAdmin.ID, updatedShowAdmin.CreatedAt, updatedShowAdmin.CreatedByUserID, updatedShowAdmin.UpdatedAt, updatedShowAdmin.UpdatedByUserID, updatedShowAdmin.DeletedAt, updatedShowAdmin.DeletedByUserID, updatedShowAdmin.ShowID, updatedShowAdmin.UserID,
	)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	if changedRows != 1 {
		return internal.ShowAdmin{}, fmt.Errorf("Updated more than 1 row (%d)", changedRows)
	}
	return updatedShowAdmin, err
}

func updateShowAdmin(ctx context.Context, db internal.Database, showAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	defer tx.Rollback()

	result, err := updateShowAdminInTx(ctx, tx, showAdmin)
	if err != nil {
		return internal.ShowAdmin{}, err
	}

	tx.Commit()
	return result, nil
}

func deleteShowAdminInTx(ctx context.Context, tx *sqlx.Tx, newShowAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	deletedShowAdmin := newShowAdmin
	auth, err := context1.GetAuthenticationDetails(ctx)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	deletedShowAdmin.UpdatedAt = time.Now()
	deletedShowAdmin.UpdatedByUserID = auth.UserID
	now := time.Now()
	deletedShowAdmin.DeletedAt = &now
	deletedShowAdmin.DeletedByUserID = &auth.UserID
	result, err := tx.ExecContext(ctx, "DELETE FROM show_admins WHERE id=$1", deletedShowAdmin.ID)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	if changedRows != 1 {
		return internal.ShowAdmin{}, fmt.Errorf("Deleted more than 1 row (%d)", changedRows)
	}
	return deletedShowAdmin, err
}

func deleteShowAdmin(ctx context.Context, db internal.Database, showAdmin internal.ShowAdmin) (internal.ShowAdmin, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.ShowAdmin{}, err
	}
	defer tx.Rollback()

	result, err := deleteShowAdminInTx(ctx, tx, showAdmin)
	if err != nil {
		return internal.ShowAdmin{}, err
	}

	tx.Commit()
	return result, nil
}