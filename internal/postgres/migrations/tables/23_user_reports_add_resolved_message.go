package tables

import (
	"anime-skip.com/public-api/internal/postgres/migrations/sqlx_migration"
)

func UserReportsAddResolvedMessage() *sqlx_migration.Migration {
	return sqlMigration(
		"MODIFY_USER_REPORTS_TABLE__add_resolved_message",
		`ALTER TABLE user_reports ADD COLUMN resolved_message varchar(500);`,
		`ALTER TABLE user_reports DROP COLUMN resolved_message;`,
	)
}
