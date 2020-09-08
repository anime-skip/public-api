package tables

// AddTimestampSource inserts the admin user
var AddTimestampSource = migrateTable(
	"MODIFY_TIMESTAMPS_TABLE__add_source",
	"timestamps",
	[]string{
		"ALTER TABLE public.timestamps",
		"ADD source integer NOT NULL",
		"CONSTRAINT source_default_value",
		"DEFAULT 0;",
	},
)
