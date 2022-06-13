package tables

var CreateTemplateTimestampsTable = createTable(
	"CREATE_TEMPLATE_TIMESTAMPS_TABLE",
	"template_timestamps",

	`CREATE TABLE template_timestamps (
		-- Many to many
		template_id uuid NOT NULL,
		timestamp_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT template_timestamps_pkey PRIMARY KEY (template_id, timestamp_id),
		UNIQUE(timestamp_id) -- Timestamps can only be on a single template
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
