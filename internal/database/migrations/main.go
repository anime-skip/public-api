package migrations

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	"github.com/aklinker1/anime-skip-backend/internal/database/migrations/seeders"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		// List entities that will be automatically migrated
		&entities.Episode{},
		&entities.EpisodeURL{},
		&entities.Preferences{},
		&entities.Show{},
		&entities.ShowAdmin{},
		&entities.Timestamp{},
		&entities.TimestampType{},
		&entities.User{},
	).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.V("Initializing database schema")

		// Add the UUID extension
		db.Exec("create extension \"uuid-ossp\";")

		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	// Seed the database
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		seeders.SeedAdminUser,
		seeders.SeedTimestampTypes,
	})
	return m.Migrate()
}
