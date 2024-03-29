package tables

var CreateShowsTable = createTable(
	"CREATE_SHOWS_TABLE",

	"shows",

	`CREATE TABLE shows (
		-- Soft Delete Entity
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,
		
		-- Custom Fields
		name text,
		original_name text,
		website text,
		image text,

		-- Constraints
		CONSTRAINT shows_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
