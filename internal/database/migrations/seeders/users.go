package seeders

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var adminUser = &entities.User{
	ID:            adminUUID,
	CreatedAt:     now,
	DeletedAt:     nil,
	Username:      "the_admin",
	Email:         "aaronklinker1@gmail.com",
	PasswordHash:  "password",
	ProfileURL:    "https://ca.slack-edge.com/T02F01E85-UD1534SV6-df241d756573-512",
	EmailVerified: true,
	Role:          0,
}

// SeedAdminUser inserts the admin user
var SeedAdminUser *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_ADMIN_USER",
	Migrate: func(db *gorm.DB) error {
		command := fmt.Sprintf(
			"INSERT INTO public.users(id, created_at, username, email, password_hash, profile_url, email_verified, role) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %t, %d);",
			adminUser.ID.String(),
			"2019-12-25 17:35:37.485712-06",
			adminUser.Username,
			adminUser.Email,
			adminUser.PasswordHash,
			adminUser.ProfileURL,
			adminUser.EmailVerified,
			adminUser.Role,
		)
		return db.Exec(command).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(adminUser).Error
	},
}
