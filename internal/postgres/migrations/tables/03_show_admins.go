package tables

var CreateShowAdminsTable = createTable(
	"CREATE_SHOW_ADMINS_TABLE",
	"show_admins",

	`CREATE TABLE show_admins (
		-- Soft Delete Entity
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,
		
		-- Custom Fields
		show_id uuid NOT NULL,
		user_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT show_admins_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;`,
)
