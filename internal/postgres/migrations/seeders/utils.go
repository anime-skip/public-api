package seeders

import (
	"time"

	"anime-skip.com/timestamps-service/internal"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

var now = time.Now()

func basicEntity(id string) internal.BaseEntity {
	return internal.BaseEntity{
		ID:              uuid.FromStringOrNil(id),
		CreatedAt:       now,
		CreatedByUserID: adminUUID,
		UpdatedAt:       now,
		UpdatedByUserID: adminUUID,
		DeletedAt:       nil,
		DeletedByUserID: nil,
	}
}

func insertUser(tx *sqlx.Tx, user internal.User, createdAt string) error {
	_, err := tx.Exec(
		`INSERT INTO users(
			id,  created_at,  username,  email,  password_hash,  profile_url,  email_verified,  role
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		)`,
		user.ID,
		createdAt,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.ProfileURL,
		user.EmailVerified,
		user.Role,
	)
	return err
}

func insertPreferences(tx *sqlx.Tx, userID uuid.UUID, createdAt string) error {
	_, err := tx.Exec(
		"INSERT INTO preferences(id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		userID,
		userID,
		createdAt,
		createdAt,
	)
	return err
}

func insertTimestampType(tx *sqlx.Tx, timestampType internal.TimestampType) error {
	_, err := tx.Exec(
		`INSERT INTO timestamp_types(
			id,
			created_at,
			created_by_user_id,
			updated_at,
			updated_by_user_id,
			deleted_at,
			deleted_by_user_id,
			name,
			description
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)`,
		timestampType.ID,
		timestampType.CreatedAt,
		timestampType.CreatedByUserID,
		timestampType.UpdatedAt,
		timestampType.UpdatedByUserID,
		timestampType.DeletedAt,
		timestampType.DeletedByUserID,
		timestampType.Name,
		timestampType.Description,
	)
	return err
}

func deleteTimestampType(tx *sqlx.Tx, timestampTypeID uuid.UUID) error {
	_, err := tx.Exec("DELETE FROM timestamp_types WHERE id=$1", timestampTypeID)
	return err
}

func insertAPIClient(tx *sqlx.Tx, client internal.APIClient) error {
	_, err := tx.Exec(
		`INSERT INTO api_clients(
			id,
			created_at,
			created_by_user_id,
			updated_at,
			updated_by_user_id,
			deleted_at,
			deleted_by_user_id,
			user_id,
			app_name,
			description,
			allowed_origins,
			rate_limit_rpm
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
		)`,
		client.ID,
		client.CreatedAt,
		client.CreatedByUserID,
		client.UpdatedAt,
		client.UpdatedByUserID,
		client.DeletedAt,
		client.DeletedByUserID,
		client.UserID,
		client.AppName,
		client.Description,
		client.AllowedOrigins,
		client.RateLimitRPM,
	)
	return err
}

func deleteAPIClient(tx *sqlx.Tx, clientID uuid.UUID) error {
	_, err := tx.Exec("DELETE FROM api_clients WHERE id=$1", clientID)
	return err
}
