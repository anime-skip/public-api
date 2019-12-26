package tables

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"strings"
)

func migrateTable(tableName string, commands ...[]string) *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "CREATE_" + strings.ToUpper(tableName) + "_TABLE",
		Migrate: func(db *gorm.DB) error {
			for _, command := range commands {
				db = db.Exec(strings.Join(command, "\n"))
			}
			return db.Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.DropTable(tableName).Error
		},
	}
}
