package postgres

import (
	"context"
	"fmt"
	"strings"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findAPIClients(ctx context.Context, tx internal.Tx, filter internal.APIClientsFilter) ([]internal.APIClient, error) {
	// Build query
	where := []string{"1 = 1"}
	args := []any{}
	if !utils.ValueOr(filter.IncludeDeleted, false) {
		where = append(where, "deleted_at IS NULL")
	}
	if filter.UserID != nil {
		args = append(args, *filter.UserID)
		condition := fmt.Sprintf("user_id = $%d", len(args))
		where = append(where, condition)
	}
	if filter.ID != nil {
		args = append(args, *filter.ID)
		condition := fmt.Sprintf("id = $%d", len(args))
		where = append(where, condition)
	}
	conditions := strings.Join(where, " AND ")
	var limitOffset string
	if p := filter.Pagination; p != nil {
		args = append(args, p.Limit, p.Offset)
		limitOffset = fmt.Sprintf("LIMIT $%d OFFSET $%d", len(args)-1, len(args))
	}
	query := fmt.Sprintf(
		`
			SELECT 
				id
				created_at
				created_by_user_id
				updated_at
				updated_by_user_id
				deleted_at
				deleted_by_user_id
				user_id
				app_name
				description
				allowed_origins
				rate_limit_rpm
			FROM api_clients
			WHERE %s
			%s
		`,
		conditions,
		limitOffset,
	)

	// Execute
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to query API clients",
			Op:      "findAPIClients",
			Err:     err,
		}
	}
	defer rows.Close()

	result := make([]internal.APIClient, 0)
	for rows.Next() {
		var client internal.APIClient
		err = rows.Scan(
			&client.ID,
			&client.CreatedAt,
			&client.CreatedByUserID,
			&client.UpdatedAt,
			&client.UpdatedByUserID,
			&client.DeletedAt,
			&client.DeletedByUserID,
			&client.UserID,
			&client.AppName,
			&client.Description,
			&client.AllowedOrigins,
			&client.RateLimitRPM,
		)
		if err != nil {
			return nil, &internal.Error{
				Code:    internal.EINTERNAL,
				Message: "Failed to load API clients",
				Op:      "findAPIClients",
				Err:     err,
			}
		}
		result = append(result, client)
	}
	return result, rows.Err()
}

func findAPIClient(ctx context.Context, tx internal.Tx, filter internal.APIClientsFilter) (internal.APIClient, error) {
	all, err := findAPIClients(ctx, tx, filter)
	if err != nil {
		return internal.ZeroAPIClient, err
	} else if len(all) == 0 {
		return internal.ZeroAPIClient, internal.NewNotFound("API client", "findAPIClient")
	}
	return all[0], nil
}

func createAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient, createdBy uuid.UUID) (internal.APIClient, error) {
	var id uuid.UUID
	err := utils.RandomID(&id)
	if err != nil {
		return apiClient, err
	}
	apiClient.ID = id.String()

	apiClient.CreatedAt = *now()
	apiClient.CreatedByUserID = createdBy
	apiClient.UpdatedAt = *now()
	apiClient.UpdatedByUserID = createdBy

	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO api_clients
			(
				id,
				created_at,
				created_by_user_id,
				updated_at,
				updated_by_user_id,
				user_id,
				app_name,
				description,
				allowed_origins,
				rate_limit_rpm
			)
			VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
			)
		`,
		apiClient.ID,
		apiClient.CreatedAt,
		apiClient.CreatedByUserID,
		apiClient.UpdatedAt,
		apiClient.UpdatedByUserID,
		apiClient.UserID,
		apiClient.AppName,
		apiClient.Description,
		apiClient.AllowedOrigins,
		apiClient.RateLimitRPM,
	)

	if err != nil {
		return apiClient, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create API client",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}
	return apiClient, nil
}

func updateAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient, updatedBy uuid.UUID) (internal.APIClient, error) {
	apiClient.UpdatedAt = *now()
	apiClient.UpdatedByUserID = updatedBy

	_, err := tx.ExecContext(
		ctx,
		`
			UPDATE api_clients
			SET
				updated_at = $1,
				updated_by_user_id = $2,
				deleted_at = $3,
				deleted_by_user_id = $4,
				user_id = $5,
				app_name = $6,
				description = $7,
				allowed_origins = $8,
				rate_limit_rpm = $9
			WHERE
				id = $10
		`,
		apiClient.UpdatedAt,
		apiClient.UpdatedByUserID,
		apiClient.DeletedAt,
		apiClient.UpdatedByUserID,
		apiClient.UserID,
		apiClient.AppName,
		apiClient.Description,
		apiClient.AllowedOrigins,
		apiClient.RateLimitRPM,
		apiClient.ID,
	)

	if err != nil {
		return apiClient, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update API client",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}
	return apiClient, nil
}

func deleteAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient, deletedBy uuid.UUID) (internal.APIClient, error) {
	apiClient.DeletedByUserID = &deletedBy
	apiClient.DeletedAt = now()
	return updateAPIClient(ctx, tx, apiClient, deletedBy)
}

func deleteCascadeAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient, deletedBy uuid.UUID) (internal.APIClient, error) {
	log.V("Deleting api client: %v", apiClient.ID)
	return deleteAPIClient(ctx, tx, apiClient, deletedBy)
}
