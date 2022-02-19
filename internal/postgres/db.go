package postgres

import (
	"database/sql"
	"fmt"

	"anime-skip.com/timestamps-service/internal"
	_ "github.com/lib/pq"
)

func Open(url string, disableSsl bool) internal.Database {
	println("Connecting to postgres...")
	sslmode := "require"
	if disableSsl {
		sslmode = "disable"
	}
	connectionString := fmt.Sprintf("%s?sslmode=%s", url, sslmode)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}
