package tables

var CreateTemplatesTable = migrateTable(
	"CREATE_TEMPLATES_TABLE",
	"templates",

	`CREATE TABLE public.templates (
		-- Soft Delete Entity
		id uuid NOT NULL DEFAULT uuid_generate_v4(),
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
		UNIQUE(source_episode_id) -- An episode cannot be the source for multiple templates
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;

	-- Indices
	CREATE INDEX "idx_template_show_id" ON public.templates("show_id");
	CREATE INDEX "idx_template_source_episode_id" ON public.templates("source_episode_id");`,
)
