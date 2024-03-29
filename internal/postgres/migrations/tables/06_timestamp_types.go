package tables

var CreateTimestampTypesTable = createTable(
	"CREATE_TIMESTAMP_TYPES_TABLE",
	"timestamp_types",

	`CREATE TABLE timestamp_types (
		-- Soft Delete Entity
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,
		
		-- Custom Fields
		name text NOT NULL,
		description text NOT NULL,

		-- Constraints
		CONSTRAINT timestamp_types_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
