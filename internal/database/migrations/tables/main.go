package tables

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"strings"
)

func migrateTable(tableName string, command []string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "CREATE_" + strings.ToUpper(tableName) + "_TABLE",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(strings.Join(command, "\n")).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(tableName).Error
		},
	}
}
