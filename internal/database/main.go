package database

import (
	"fmt"

	"github.com/aklinker1/anime-skip-backend/internal/database/migrations"
	"github.com/aklinker1/anime-skip-backend/internal/utils"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"

	// PostgreSQL dialect
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var autoMigrate, seedDB bool
var connectionString, dialect string

// ORM struct to store the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

var ORMInstance *ORM

func init() {
	var sslmode = "require"
	if utils.EnvBool("POSTGRES_DISABLE_SSL") {
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
}

// Factory creates a db connection and returns the pointer to the GORM instance
func Factory() (*ORM, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Error: ", err)
	}
	log.D("Connected to PostgreSQL")
	ORMInstance = &ORM{
		DB: db,
	}

	// Enable SQL logs
	db.LogMode(utils.EnvBool("LOG_SQL"))
	// db.SetLogger(log.SQLLogger)

	// Automigrate tables
	log.D("Running migrations if necessary")
	err = migrations.Run(ORMInstance.DB)

	// Adding plugins
	db.Callback().Create().Register("anime_skip:update_created_by", updateColumn("created_by_user_id"))
	db.Callback().Update().Register("anime_skip:update_updated_by", updateColumn("updated_by_user_id"))
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	fmt.Println()
	return ORMInstance, err
}
