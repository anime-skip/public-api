package database

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/migrations"
	"github.com/aklinker1/anime-skip-backend/pkg/utils"
	log "github.com/aklinker1/anime-skip-backend/pkg/utils/log"

	// PostgreSQL dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var autoMigrate, logMode, seedDB bool
var connectionString, dialect string

// ORM struct to store the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	var sslmode = "require"
	if utils.EnvBool("POSTGRES_SSL_DISABLED") {
		sslmode = "disable"
	}
	connectionString = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		utils.EnvString("POSTGRES_HOST"),
		utils.EnvString("POSTGRES_PORT"),
		utils.EnvString("POSTGRES_USER"),
		utils.EnvString("POSTGRES_PASSWORD"),
		utils.EnvString("POSTGRES_DBNAME"),
		sslmode,
	)
	seedDB = utils.EnvBool("POSTGRES_ENABLE_SEEDING")
	logMode = utils.EnvBool("POSTGRES_ENABLE_LOGS")
	autoMigrate = utils.EnvBool("POSTGRES_ENABLE_AUTO_MIGRATE")
}

// Factory creates a db connection and returns the pointer to the GORM instance
func Factory() (*ORM, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}

	// Enable SQL logs
	db.LogMode(logMode)

	// Automigrate tables
	if autoMigrate {
		err = migrations.ServiceAutoMigration(orm.DB)
	}

	log.D("[ORM] Database connection initialized.")
	return orm, err
}
