package tables

// AddTimestampSource inserts the new source column with a default value of 0
var AddTimestampSource = sqlMigration(
	"MODIFY_TIMESTAMPS_TABLE__add_source",
	`
	ALTER TABLE timestamps
		ADD source integer NOT NULL
			CONSTRAINT source_default_value
			DEFAULT 0;
	`,
	`
	ALTER TABLE timestamps
		DROP COLUMN source;
	`,
)
