package tables

var CreateUserReportsTable = createTable(
	"CREATE_USER_REPORTS_TABLE",
	"user_reports",

	`CREATE TABLE user_reports (
		-- Soft Delete Entity
		id uuid NOT NULL,
		created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_by_user_id uuid NOT NULL,
		updated_at timestamp with time zone NOT NULL,
		updated_by_user_id uuid NOT NULL,
		deleted_at timestamp with time zone,
		deleted_by_user_id uuid,

		-- Custom Fields
		message varchar(500) NOT NULL,
		reported_from_url text NOT NULL,
		resolved boolean NOT NULL DEFAULT false,
		timestamp_id uuid,
		episode_id uuid,
		episode_url text,
		show_id uuid,

		-- Constraints
		CONSTRAINT user_reports_pkey PRIMARY KEY (id)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;

	-- Indices
	CREATE INDEX "idx_user_created_at" ON user_reports("created_at");
	CREATE INDEX "idx_user_resolved" ON user_reports("resolved");
	CREATE INDEX "idx_user_report_timestamp_id" ON user_reports("timestamp_id");
	CREATE INDEX "idx_user_report_episode_id" ON user_reports("episode_id");
	CREATE INDEX "idx_user_report_episode_url" ON user_reports("episode_url");
	CREATE INDEX "idx_user_report_show_id" ON user_reports("show_id");`,
)
