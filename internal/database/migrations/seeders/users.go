package seeders

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/pkg/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var adminUser = &entities.User{
	// ID:            adminUUID,
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
		log.V("AdminID", adminUUID)
		return db.Create(adminUser).Error
	},
	Rollback: func(db *gorm.DB) error {
		return db.Delete(adminUser).Error
	},
}
