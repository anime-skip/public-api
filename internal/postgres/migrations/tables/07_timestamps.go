package tables

var CreateTimestampsTable = createTable(
	"CREATE_TIMESTAMPS_TABLE",
	"timestamps",

	`CREATE TABLE timestamps (
		-- Soft Delete Entity
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,
		
		-- Custom Fields
		at numeric,
		type_id uuid NOT NULL,
		episode_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT timestamps_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
