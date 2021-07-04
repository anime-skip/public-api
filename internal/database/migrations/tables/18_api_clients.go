package tables

var CreateClientsTable = migrateTable(
	"CREATE_CLIENTS_TABLE",
	"api_clients",

	`CREATE TABLE public.api_clients (
		-- Soft Delete Entity
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,

		-- Custom Fields
		client_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT template_timestamps_pkey PRIMARY KEY (template_id, timestamp_id),
		UNIQUE(timestamp_id) -- Timestamps can only be on a single template
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
