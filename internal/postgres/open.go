package postgres

import (
	"fmt"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:generate go run ../../cmd/sql-gen/main.go

func Open(url string, disableSsl bool, targetVersion int) internal.Database {
	log.D("Connecting to postgres...")
	sslmode := "require"
	if disableSsl {
		sslmode = "disable"
	}
	connectionString := fmt.Sprintf("%s?sslmode=%s", url, sslmode)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	err = migrate(db, targetVersion)
	if err != nil {
		panic(err)
	}

	return db
}
