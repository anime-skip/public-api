package postgres

import (
	"context"
	"database/sql"
	"time"

	"anime-skip.com/public-api/internal"
)

func now() *time.Time {
	n := time.Now()
	return &n
}

func count(ctx context.Context, tx internal.Tx, query string, args ...any) (int, error) {
	row, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	var count int
	row.Next()
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func inTx[T any](ctx context.Context, db internal.Database, write bool, zeroValue T, exec func(tx internal.Tx) (T, error)) (T, error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: !write})
	if err != nil {
		return zeroValue, err
	}
	defer tx.Rollback()

	res, err := exec(tx)
	if err != nil {
		return zeroValue, err
	}

	tx.Commit()
	return res, nil
}
