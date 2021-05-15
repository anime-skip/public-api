package tables

var CreateEpisodesTable = migrateTable(
	"CREATE_EPISODES_TABLE",
	"episodes",

	`CREATE TABLE public.episodes (
		-- Soft Delete Entity
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,
		
		-- Custom Fields
		season integer,
		"number" integer,
		absolute_number integer,
		name text,
		show_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT episodes_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
