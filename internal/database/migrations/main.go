package migrations

import (
	"github.com/aklinker1/anime-skip-backend/internal/database/migrations/seeders"
	"github.com/aklinker1/anime-skip-backend/internal/database/migrations/tables"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// Run migrates all the tables and modifications to the connected source
func Run(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.V("Initializing database schema")

		// Add the UUID extension
		return db.Exec("create extension \"uuid-ossp\";").Error
	})
	m.Migrate()

	// Create the Tables
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		tables.CreateUsersTable,
		tables.CreatePreferencesTable,
		tables.CreateShowsTable,
		tables.CreateShowAdminsTable,
		tables.CreateEpisodesTable,
		tables.CreateEpisodeURLsTable,
		tables.CreateTimestampTypesTable,
		tables.CreateTimestampsTable,
	})
	m.Migrate()

	// Seed the database
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		seeders.SeedAdminUser,
		seeders.SeedTimestampTypes,
	})
	return m.Migrate()
}
