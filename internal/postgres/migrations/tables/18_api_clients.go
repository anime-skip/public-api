package tables

var CreateAPIClientsTable = createTable(
	"CREATE_CLIENTS_TABLE",
	"api_clients",

	`CREATE TABLE api_clients (
		-- Soft Delete Entity
		id text NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,

		-- Custom Fields
		user_id uuid NOT NULL,
		app_name text NOT NULL,
		description text NOT NULL,
		allowed_origins text[],
		rate_limit_rpm integer,

		-- Constraints
		CONSTRAINT api_clients_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
