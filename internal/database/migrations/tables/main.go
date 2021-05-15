package tables

import (
	"anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func migrateTable(migrationID string, tableName string, sql string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: migrationID,
		Migrate: func(db *gorm.DB) error {
			return db.Exec(sql).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(tableName).Error
		},
	}
}

func migrateTableChange(migrationID string, upSql string, downSql string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: migrationID,
		Migrate: func(db *gorm.DB) error {
			log.V(migrationID)
			return db.Exec(upSql).Error
		},
		Rollback: func(db *gorm.DB) error {
			log.V(migrationID)
			return db.Exec(downSql).Error
		},
	}
}
