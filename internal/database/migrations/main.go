package migrations

import (
	"anime-skip.com/backend/internal/database/migrations/seeders"
	"anime-skip.com/backend/internal/database/migrations/tables"
	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func getCurrentMigration(migrations []*gormigrate.Migration) int {
	finalMigration := len(migrations) - 1
	currentMigrationPtr := env.DATABASE_MIGRATION

	if currentMigrationPtr == nil {
		return finalMigration
	} else {
		return *currentMigrationPtr
	}
}

// Run migrates all the tables and modifications to the connected source
func Run(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.V("Initializing database schema")
		// Add the UUID extension
		return db.Exec("create extension \"uuid-ossp\";").Error
	})
	err := m.Migrate()
	if err != nil {
		return err
	}

	// Create the Tables
	migrations := []*gormigrate.Migration{
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
	}
	currentMigration := getCurrentMigration(migrations)
	currentMigrationId := migrations[currentMigration].ID
	log.D("Current migration: %s", currentMigrationId)

	// Migrate
	log.D("Running migrations if necessary...")
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations[0:currentMigration+1])
	err = m.Migrate()
	if err != nil {
		return err
	}

	// Rollback
	r := gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	log.D("Running rollbacks if necessary...")
	err = r.RollbackTo(currentMigrationId)
	if err != nil {
		return err
	}

	if !env.DATABASE_ENABLE_SEEDING {
		return nil
	}

	// Seed
	log.D("Running seeders if necessary...")
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		seeders.SeedAdminUser,
		seeders.SeedTimestampTypes,
		seeders.SeedUnknownTimestampType,
	})
	return m.Migrate()
}
