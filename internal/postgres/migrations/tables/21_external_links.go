package tables

var CreateExternalLinksTable = createTable(
	"CREATE_EXTERNAL_LINKS_TABLE",
	"external_links",

	`CREATE TABLE external_links (
		url text NOT NULL,
		show_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT external_links_pkey PRIMARY KEY (url, show_id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
