package seeders

import (
	"time"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/postgres/migrations/sqlx_migration"
	"github.com/jmoiron/sqlx"
)

var adminUUID = "00000000-0000-0000-0000-000000000000"
var adminUser = &internal.User{
	ID:            adminUUID,
	CreatedAt:     time.Now(),
	DeletedAt:     nil,
	Username:      "the_admin",
	Email:         "admin@anime-skip.com",
	PasswordHash:  "password", // Can't sign in when this is a not encrypted, so it's safe to have here
	ProfileURL:    "https://ca.slack-edge.com/T02F01E85-UD1534SV6-df241d756573-512",
	EmailVerified: true,
	Role:          0,
}

// SeedAdminUser inserts the admin user
var SeedAdminUser = &sqlx_migration.Migration{
	ID: "SEED_ADMIN_USER",
	Up: func(tx *sqlx.Tx) error {
		err := insertUser(tx, *adminUser, "2019-12-25 17:35:37.485712-06")
		if err != nil {
			return err
		}
		return insertPreferences(tx, adminUser.ID, "2019-12-25 17:35:37.485712-06")
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec("DELETE FROM users WHERE id=$1", adminUser.ID)
		return err
	},
}
