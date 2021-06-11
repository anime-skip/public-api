package database

import (
	"fmt"

	"anime-skip.com/backend/internal/database/migrations"
	"anime-skip.com/backend/internal/utils/env"
	log "anime-skip.com/backend/internal/utils/log"

	// PostgreSQL dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ORM struct to store the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

var ORMInstance *ORM

// Factory creates a db connection and returns the pointer to the GORM instance
func Factory() (*ORM, error) {
	var sslmode = "require"
	if env.DATABASE_DISABLE_SSL {
		sslmode = "disable"
	}
	connectionString := fmt.Sprintf("%s?sslmode=%s", env.DATABASE_URL, sslmode)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Error: ", err)
	}
	log.I("Connected to PostgreSQL")
	ORMInstance = &ORM{
		DB: db,
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	// Enable SQL logs
	db.LogMode(env.LOG_SQL)
	db.SetLogger(log.SQLLogger)

	// Migrations
	if !env.DATABASE_SKIP_MIGRATIONS {
		err = migrations.Run(ORMInstance.DB)
		if err != nil {
			panic(err)
		}
	}

	// Adding plugins
	log.D("Registering update callbacks...")
	db.Callback().Create().Before("gorm:create").Register("anime_skip_create:update_updated_by", updateColumn("UpdatedByUserId"))
	db.Callback().Create().Before("gorm:create").Register("anime_skip_create:update_created_by", updateColumn("CreatedByUserId"))
	db.Callback().Update().Before("gorm:update").Register("anime_skip_update:update_updated_by", updateColumn("UpdatedByUserId"))
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	return ORMInstance, err
}
