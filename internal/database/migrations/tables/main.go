package tables

import (
	"strings"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func migrateTable(migrationID string, tableName string, command []string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: migrationID,
		Migrate: func(db *gorm.DB) error {
			return db.Exec(strings.Join(command, "\n")).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(tableName).Error
		},
	}
}

func migrateTableChange(migrationID string, upCommand []string, downCommand []string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: migrationID,
		Migrate: func(db *gorm.DB) error {
			return db.Exec(strings.Join(upCommand, "\n")).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(strings.Join(downCommand, "\n")).Error
		},
	}
}
