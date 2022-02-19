package postgres

import (
	"fmt"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Open(url string, disableSsl bool, targetVersion int) internal.Database {
	log.D("Connecting to postgres...")
	sslmode := "require"
	if disableSsl {
		sslmode = "disable"
	}
	connectionString := fmt.Sprintf("%s?sslmode=%s", url, sslmode)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
		return nil
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	err = migrate(db, targetVersion)
	if err != nil {
		log.Panic(err)
		return nil
	}

	return db
}
