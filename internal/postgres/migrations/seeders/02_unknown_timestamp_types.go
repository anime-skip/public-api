package seeders

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/config"
	"anime-skip.com/timestamps-service/internal/postgres/migrations/sqlx_migration"
	"github.com/jmoiron/sqlx"
)

var unknownTimestampType = internal.TimestampType{
	Name:        "Unknown",
	Description: "A timestamp that was imported from a different system that didn't have enough information to decide what type it was. These are treated as unskippable",
	BaseEntity:  basicEntity(config.TIMESTAMP_ID_UNKNOWN),
}

// SeedUnknownTimestampType inserts the a new type, "Unknown"
var SeedUnknownTimestampType = &sqlx_migration.Migration{
	ID: "SEED_UNKNOWN_TIMESTAMP_TYPE",
	Up: func(tx *sqlx.Tx) error {
		return insertTimestampType(tx, unknownTimestampType)
	},
	Down: func(tx *sqlx.Tx) error {
		return deleteTimestampType(tx, unknownTimestampType.ID)
	},
}
