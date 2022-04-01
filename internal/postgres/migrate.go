package postgres

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/postgres/migrations/seeders"
	"anime-skip.com/public-api/internal/postgres/migrations/sqlx_migration"
	"anime-skip.com/public-api/internal/postgres/migrations/tables"
)

// migrate all the migrations and seeders
func migrate(db internal.Database, dbVersion int, enableSeeding bool) error {
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
		/* 20 */ tables.AddMissingPreferenceDefaults(),
	}, dbVersion)
	if err != nil {
		return err
	}

	if !enableSeeding {
		return nil
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
