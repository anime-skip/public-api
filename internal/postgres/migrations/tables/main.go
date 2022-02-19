package tables

import (
	"fmt"

	"anime-skip.com/timestamps-service/internal/postgres/migrations/sqlx_migration"
	"github.com/jmoiron/sqlx"
)

func createTable(migrationID string, tableName string, sql string) *sqlx_migration.Migration {
	return &sqlx_migration.Migration{
		ID: migrationID,
		Up: func(tx *sqlx.Tx) error {
			_, err := tx.Exec(sql)
			return err
		},
		Down: func(tx *sqlx.Tx) error {
			sql := fmt.Sprintf("DROP TABLE %s", tableName)
			_, err := tx.Exec(sql)
			return err
		},
	}
}

func sqlMigration(migrationID string, upSql string, downSql string) *sqlx_migration.Migration {
	return &sqlx_migration.Migration{
		ID: migrationID,
		Up: func(tx *sqlx.Tx) error {
			_, err := tx.Exec(upSql)
			return err
		},
		Down: func(tx *sqlx.Tx) error {
			_, err := tx.Exec(downSql)
			return err
		},
	}
}
