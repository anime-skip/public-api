// Code generated by cmd/sqlgen/main.go, DO NOT EDIT.

package postgres

import (
	internal "anime-skip.com/timestamps-service/internal"
	"context"
	"fmt"
	uuid "github.com/gofrs/uuid"
	sqlx "github.com/jmoiron/sqlx"
	"time"
)

func getUserByID(ctx context.Context, db internal.Database, id uuid.UUID) (internal.User, error) {
	var user internal.User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE id=$1", id)
	return user, err
}

func getUserByUsername(ctx context.Context, db internal.Database, username string) (internal.User, error) {
	var user internal.User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL", username)
	return user, err
}

func getUnscopedUserByUsername(ctx context.Context, db internal.Database, username string) (internal.User, error) {
	var user internal.User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE username=$1", username)
	return user, err
}

func getUserByEmail(ctx context.Context, db internal.Database, email string) (internal.User, error) {
	var user internal.User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE email=$1 AND deleted_at IS NULL", email)
	return user, err
}

func getUnscopedUserByEmail(ctx context.Context, db internal.Database, email string) (internal.User, error) {
	var user internal.User
	err := db.GetContext(ctx, &user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}

func insertUserInTx(ctx context.Context, tx *sqlx.Tx, user internal.User) (internal.User, error) {
	newUser := user
	newUser.CreatedAt = time.Now()
	newUser.DeletedAt = nil
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO users(id, created_at, deleted_at, username, email, password_hash, profile_url, email_verified, role) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		newUser.ID, newUser.CreatedAt, newUser.DeletedAt, newUser.Username, newUser.Email, newUser.PasswordHash, newUser.ProfileURL, newUser.EmailVerified, newUser.Role,
	)
	if err != nil {
		return internal.User{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.User{}, err
	}
	if changedRows != 1 {
		return internal.User{}, fmt.Errorf("Inserted %d rows, not 1", changedRows)
	}
	return newUser, err
}

func insertUser(ctx context.Context, db internal.Database, user internal.User) (internal.User, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.User{}, err
	}
	defer tx.Rollback()

	newUser, err := insertUserInTx(ctx, tx, user)
	if err != nil {
		return internal.User{}, err
	}

	tx.Commit()
	return newUser, nil
}

func updateUserInTx(ctx context.Context, tx *sqlx.Tx, newUser internal.User) (internal.User, error) {
	updatedUser := newUser
	result, err := tx.ExecContext(
		ctx,
		"UPDATE users SET id=$1, created_at=$2, deleted_at=$3, username=$4, email=$5, password_hash=$6, profile_url=$7, email_verified=$8, role=$9",
		updatedUser.ID, updatedUser.CreatedAt, updatedUser.DeletedAt, updatedUser.Username, updatedUser.Email, updatedUser.PasswordHash, updatedUser.ProfileURL, updatedUser.EmailVerified, updatedUser.Role,
	)
	if err != nil {
		return internal.User{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.User{}, err
	}
	if changedRows != 1 {
		return internal.User{}, fmt.Errorf("Updated %d rows, not 1", changedRows)
	}
	return updatedUser, err
}

func updateUser(ctx context.Context, db internal.Database, user internal.User) (internal.User, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.User{}, err
	}
	defer tx.Rollback()

	newUser, err := updateUserInTx(ctx, tx, user)
	if err != nil {
		return internal.User{}, err
	}

	tx.Commit()
	return newUser, nil
}
