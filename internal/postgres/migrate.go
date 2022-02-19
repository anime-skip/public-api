package postgres

import (
	"anime-skip.com/timestamps-service/internal/postgres/migrations/seeders"
	"anime-skip.com/timestamps-service/internal/postgres/migrations/sqlx_migration"
	"anime-skip.com/timestamps-service/internal/postgres/migrations/tables"
	"github.com/jmoiron/sqlx"
)

// migrate all the migrations and seeders
func migrate(db *sqlx.DB, dbVersion int) error {
	tx := db.MustBegin()
	defer tx.Rollback()

	err := sqlx_migration.RunMigrations(tx, "schema migrations", []*sqlx_migration.Migration{
		/* 0  */ tables.CreateUsersTable,
		/* 1  */ tables.CreatePreferencesTable,
		/* 2  */ tables.CreateShowsTable,
		/* 3  */ tables.CreateShowAdminsTable,
		/* 4  */ tables.CreateEpisodesTable,
		/* 5  */ tables.CreateEpisodeURLsTable,
		/* 6  */ tables.CreateTimestampTypesTable,
		/* 7  */ tables.CreateTimestampsTable,
		/* 8  */ tables.ModifyEpisodeUrlsTableHardDelete,
		/* 9  */ tables.LowercaseAllEmails,
		/* 10 */ tables.EpisodeColumnsToStrings,
		/* 11 */ tables.AddTimestampSource,
		/* 12 */ tables.AddBaseDurationToEpisodes,
		/* 13 */ tables.AddDurationToEpisodeUrls,
		/* 14 */ tables.AddTimestampsOffsetToEpisodeUrls,
		/* 15 */ tables.AddHideTimelinePreferences,
		/* 16 */ tables.CreateTemplatesTable,
		/* 17 */ tables.CreateTemplateTimestampsTable,
		/* 18 */ tables.CreateAPIClientsTable,
		/* 19 */ tables.AddColorThemePreference,
	}, dbVersion)
	if err != nil {
		return err
	}

	err = sqlx_migration.RunAllMigrations(tx, "seeders", []*sqlx_migration.Migration{
		/* 0  */ seeders.SeedAdminUser,
		/* 1  */ seeders.SeedTimestampTypes,
		/* 2  */ seeders.SeedUnknownTimestampType,
		/* 3  */ seeders.SeedKnownClientIDs,
	})
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
