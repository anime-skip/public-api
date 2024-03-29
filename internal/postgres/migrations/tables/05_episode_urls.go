package tables

var CreateEpisodeURLsTable = createTable(
	"CREATE_EPISODE_URLS_TABLE",
	"episode_urls",

	`CREATE TABLE episode_urls (
		-- Soft Delete Entity
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,

		-- Custom Fields
		url text NOT NULL,
		episode_id uuid NOT NULL,
		source integer NOT NULL,
		
		-- Constraints
		CONSTRAINT episode_urls_pkey PRIMARY KEY (url)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
