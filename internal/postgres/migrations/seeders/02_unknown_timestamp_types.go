package seeders

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/config"
	"anime-skip.com/public-api/internal/postgres/migrations/sqlx_migration"
)

var unknownTimestampType = PartialTimestamp{
	Name:        "Unknown",
	Description: "A timestamp that was imported from a different system that didn't have enough information to decide what type it was. These are treated as unskippable",
	ID:          config.TIMESTAMP_ID_UNKNOWN,
}

// SeedUnknownTimestampType inserts the a new type, "Unknown"
var SeedUnknownTimestampType = &sqlx_migration.Migration{
	ID: "SEED_UNKNOWN_TIMESTAMP_TYPE",
	Up: func(tx internal.Tx) error {
		return insertTimestampType(tx, unknownTimestampType)
	},
	Down: func(tx internal.Tx) error {
		return deleteTimestampType(tx, unknownTimestampType.ID)
	},
}
