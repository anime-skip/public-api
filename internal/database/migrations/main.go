package migrations

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/models"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.V("[Migration.InitSchema] Initializing database schema")
		switch db.Dialect().GetName() {
		case "postgres":
			// Let's create the UUID extension, the user has to ahve superuser
			// permission for now
			db.Exec("create extension \"uuid-ossp\";")
		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	// m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
	// 	jobs.SeedUsers,
	// })
	return m.Migrate()
}
