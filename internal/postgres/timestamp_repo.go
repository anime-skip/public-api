package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/errors"
	"anime-skip.com/public-api/internal/log"
)

func deleteCascadeTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp) (internal.Timestamp, error) {
	log.V("Deleting timestamp: %v", timestamp.ID)
	deletedTimestamp, err := deleteTimestampInTx(ctx, tx, timestamp)
	if err != nil {
		return internal.Timestamp{}, err
	}

	log.V("Deleting timestamp template timestamp")
	templateTimestamp, err := getTemplateTimestampByTimestampIDInTx(ctx, tx, timestamp.ID)
	if err == nil {
		_, err = deleteCascadeTemplateTimestamp(ctx, tx, templateTimestamp)
	}
	if !errors.IsRecordNotFound(err) {
		return internal.Timestamp{}, err
	}

	log.V("Done deleting timestamp: %v", timestamp.ID)
	return deletedTimestamp, nil
}
