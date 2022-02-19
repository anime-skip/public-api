package sqlx_migration

import "github.com/jmoiron/sqlx"

type Migration struct {
	ID   string
	Up   func(tx *sqlx.Tx) error
	Down func(tx *sqlx.Tx) error
}

type ExistingMigration struct {
	ID string
}
