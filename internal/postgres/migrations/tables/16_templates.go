package tables

var CreateTemplatesTable = createTable(
	"CREATE_TEMPLATES_TABLE",
	"templates",

	`CREATE TABLE templates (
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
		"type" integer NOT NULL,
		seasons text[],
		source_episode_id uuid NOT NULL,

		-- Constraints
		CONSTRAINT templates_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;

	-- Indices
	CREATE INDEX "idx_template_show_id" ON templates("show_id");
	CREATE INDEX "idx_template_source_episode_id" ON templates("source_episode_id");`,
)
