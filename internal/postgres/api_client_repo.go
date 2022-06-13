package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findAPIClients(ctx context.Context, tx internal.Tx, filter internal.APIClientsFilter) ([]internal.APIClient, error) {
	var scanned internal.APIClient
	query := sqlbuilder.Select("api_clients", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"user_id":            &scanned.UserID,
		"app_name":           &scanned.AppName,
		"description":        &scanned.Description,
		// "allowed_origins":    &scanned.AllowedOrigins,
		"rate_limit_rpm": &scanned.RateLimitRpm,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.UserID != nil {
		query.Where("user_id = ?", *filter.UserID)
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.NameContains != nil {
		query.Where("app_name ILIKE ?", "%"+*filter.NameContains+"%")
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("created_at", filter.Sort)

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findAPIClients", err)
	}
	dest := query.ScanDest()
	result := make([]internal.APIClient, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findAPIClients", err)
		}
		result = append(result, scanned)
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
	apiClient.ID = utils.RandomString(32)
	apiClient.CreatedAt = *now()
	apiClient.CreatedByUserID = &createdBy
	apiClient.UpdatedAt = *now()
	apiClient.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("api_clients", map[string]any{
		"id":                 apiClient.ID,
		"created_at":         apiClient.CreatedAt,
		"created_by_user_id": apiClient.CreatedByUserID,
		"updated_at":         apiClient.UpdatedAt,
		"updated_by_user_id": apiClient.UpdatedByUserID,
		"user_id":            apiClient.UserID,
		"app_name":           apiClient.AppName,
		"description":        apiClient.Description,
		"rate_limit_rpm":     apiClient.RateLimitRpm,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return apiClient, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "API client with the generated ID already exists, try again",
			Op:      "createAPIClient",
			Err:     err,
		}
	} else if err != nil {
		return apiClient, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create API client",
			Op:      "createAPIClient",
			Err:     err,
		}
	}
	return apiClient, nil
}

func updateAPIClient(ctx context.Context, tx internal.Tx, apiClient internal.APIClient, updatedBy uuid.UUID) (internal.APIClient, error) {
	apiClient.UpdatedAt = *now()
	apiClient.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("api_clients", apiClient.ID, map[string]any{
		"updated_at":         apiClient.UpdatedAt,
		"updated_by_user_id": apiClient.UpdatedByUserID,
		"deleted_at":         apiClient.DeletedAt,
		"deleted_by_user_id": apiClient.DeletedByUserID,
		"user_id":            apiClient.UserID,
		"app_name":           apiClient.AppName,
		"description":        apiClient.Description,
		"rate_limit_rpm":     apiClient.RateLimitRpm,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return apiClient, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update API client",
			Op:      "updateAPIClient",
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
