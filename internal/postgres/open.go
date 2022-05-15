package postgres

import (
	"database/sql"
	"fmt"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	_ "github.com/lib/pq"
)

func Open(url string, disableSsl bool, targetVersion int, enableSeeding bool) internal.Database {
	log.D("Connecting to postgres...")
	sslmode := "require"
	if disableSsl {
		sslmode = "disable"
	}
	connectionString := fmt.Sprintf("%s?sslmode=%s", url, sslmode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	err = migrate(db, targetVersion, enableSeeding)
	if err != nil {
		panic(err)
	}

	return db
}
