// Code generated by cmd/sqlgen/main.go, DO NOT EDIT.

package postgres

import (
	internal "anime-skip.com/timestamps-service/internal"
	context1 "anime-skip.com/timestamps-service/internal/context"
	errors1 "anime-skip.com/timestamps-service/internal/errors"
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/gofrs/uuid"
	"time"
)

func getAPIClientByID(ctx context.Context, db internal.Database, id string) (internal.APIClient, error) {
	var apiClient internal.APIClient
	err := db.GetContext(ctx, &apiClient, "SELECT * FROM api_clients WHERE id=$1", id)
	if errors.Is(err, sql.ErrNoRows) {
		return internal.APIClient{}, errors1.NewRecordNotFound(fmt.Sprintf("APIClient.id=%s", id))
	}
	return apiClient, err
}

func getAPIClientsByUserID(ctx context.Context, db internal.Database, userID uuid.UUID) ([]internal.APIClient, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM api_clients WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apiClients := []internal.APIClient{}
	for rows.Next() {
		var apiClient internal.APIClient
		err = rows.StructScan(&apiClient)
		if err != nil {
			return nil, err
		}
		apiClients = append(apiClients, apiClient)
	}
	return apiClients, nil
}

func getUnscopedAPIClientsByUserID(ctx context.Context, db internal.Database, userID uuid.UUID) ([]internal.APIClient, error) {
	rows, err := db.QueryxContext(ctx, "SELECT * FROM api_clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	apiClients := []internal.APIClient{}
	for rows.Next() {
		var apiClient internal.APIClient
		err = rows.StructScan(&apiClient)
		if err != nil {
			return nil, err
		}
		apiClients = append(apiClients, apiClient)
	}
	return apiClients, nil
}

func insertAPIClientInTx(ctx context.Context, tx internal.Tx, apiClient internal.APIClient) (internal.APIClient, error) {
	newAPIClient := apiClient
	claims, err := context1.GetAuthClaims(ctx)
	if err != nil {
		return internal.APIClient{}, err
	}
	newAPIClient.CreatedAt = time.Now()
	newAPIClient.CreatedByUserID = claims.UserID
	newAPIClient.UpdatedAt = time.Now()
	newAPIClient.UpdatedByUserID = claims.UserID
	newAPIClient.DeletedAt = nil
	newAPIClient.DeletedByUserID = nil
	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO api_clients(id, created_at, created_by_user_id, updated_at, updated_by_user_id, deleted_at, deleted_by_user_id, user_id, app_name, description, allowed_origins, rate_limit_rpm) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		newAPIClient.ID, newAPIClient.CreatedAt, newAPIClient.CreatedByUserID, newAPIClient.UpdatedAt, newAPIClient.UpdatedByUserID, newAPIClient.DeletedAt, newAPIClient.DeletedByUserID, newAPIClient.UserID, newAPIClient.AppName, newAPIClient.Description, newAPIClient.AllowedOrigins, newAPIClient.RateLimitRPM,
	)
	if err != nil {
		return internal.APIClient{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.APIClient{}, err
	}
	if changedRows != 1 {
		return internal.APIClient{}, fmt.Errorf("Inserted more than 1 row (%d)", changedRows)
	}
	return newAPIClient, err
}

func insertAPIClient(ctx context.Context, db internal.Database, apiClient internal.APIClient) (internal.APIClient, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.APIClient{}, err
	}
	defer tx.Rollback()

	result, err := insertAPIClientInTx(ctx, tx, apiClient)
	if err != nil {
		return internal.APIClient{}, err
	}

	tx.Commit()
	return result, nil
}

func updateAPIClientInTx(ctx context.Context, tx internal.Tx, newAPIClient internal.APIClient) (internal.APIClient, error) {
	updatedAPIClient := newAPIClient
	claims, err := context1.GetAuthClaims(ctx)
	if err != nil {
		return internal.APIClient{}, err
	}
	updatedAPIClient.UpdatedAt = time.Now()
	updatedAPIClient.UpdatedByUserID = claims.UserID
	result, err := tx.ExecContext(
		ctx,
		"UPDATE api_clients SET id=$1, created_at=$2, created_by_user_id=$3, updated_at=$4, updated_by_user_id=$5, deleted_at=$6, deleted_by_user_id=$7, user_id=$8, app_name=$9, description=$10, allowed_origins=$11, rate_limit_rpm=$12",
		updatedAPIClient.ID, updatedAPIClient.CreatedAt, updatedAPIClient.CreatedByUserID, updatedAPIClient.UpdatedAt, updatedAPIClient.UpdatedByUserID, updatedAPIClient.DeletedAt, updatedAPIClient.DeletedByUserID, updatedAPIClient.UserID, updatedAPIClient.AppName, updatedAPIClient.Description, updatedAPIClient.AllowedOrigins, updatedAPIClient.RateLimitRPM,
	)
	if err != nil {
		return internal.APIClient{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.APIClient{}, err
	}
	if changedRows != 1 {
		return internal.APIClient{}, fmt.Errorf("Updated more than 1 row (%d)", changedRows)
	}
	return updatedAPIClient, err
}

func updateAPIClient(ctx context.Context, db internal.Database, apiClient internal.APIClient) (internal.APIClient, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.APIClient{}, err
	}
	defer tx.Rollback()

	result, err := updateAPIClientInTx(ctx, tx, apiClient)
	if err != nil {
		return internal.APIClient{}, err
	}

	tx.Commit()
	return result, nil
}

func deleteAPIClientInTx(ctx context.Context, tx internal.Tx, newAPIClient internal.APIClient) (internal.APIClient, error) {
	deletedAPIClient := newAPIClient
	claims, err := context1.GetAuthClaims(ctx)
	if err != nil {
		return internal.APIClient{}, err
	}
	deletedAPIClient.UpdatedAt = time.Now()
	deletedAPIClient.UpdatedByUserID = claims.UserID
	now := time.Now()
	deletedAPIClient.DeletedAt = &now
	deletedAPIClient.DeletedByUserID = &claims.UserID
	result, err := tx.ExecContext(ctx, "DELETE FROM api_clients WHERE id=$1", deletedAPIClient.ID)
	if err != nil {
		return internal.APIClient{}, err
	}
	changedRows, err := result.RowsAffected()
	if err != nil {
		return internal.APIClient{}, err
	}
	if changedRows != 1 {
		return internal.APIClient{}, fmt.Errorf("Deleted more than 1 row (%d)", changedRows)
	}
	return deletedAPIClient, err
}

func deleteAPIClient(ctx context.Context, db internal.Database, apiClient internal.APIClient) (internal.APIClient, error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return internal.APIClient{}, err
	}
	defer tx.Rollback()

	result, err := deleteAPIClientInTx(ctx, tx, apiClient)
	if err != nil {
		return internal.APIClient{}, err
	}

	tx.Commit()
	return result, nil
}
