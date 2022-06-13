package sqlx_migration

import "anime-skip.com/public-api/internal"

type Migration struct {
	ID   string
	Up   func(tx internal.Tx) error
	Down func(tx internal.Tx) error
}

type ExistingMigration struct {
	ID string
}
