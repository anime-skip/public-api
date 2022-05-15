package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	uuid "github.com/gofrs/uuid"
)

func deleteCascadeTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, deletedBy uuid.UUID) (internal.Timestamp, error) {
	log.V("Deleting timestamp: %v", timestamp.ID)
	deletedTimestamp, err := deleteTimestamp(ctx, tx, timestamp, deletedBy)
	if err != nil {
		return internal.Timestamp{}, err
	}

	log.V("Deleting timestamp template timestamps")
	templateTimestamps, err := findTemplateTimestamps(ctx, tx, internal.TemplateTimestampsFilter{
		TimestampID: timestamp.ID,
	})
	if err != nil {
		return internal.ZeroTimestamp, err
	}
	for _, templateTimestamp := range templateTimestamps {
		_, err := deleteCascadeTemplateTimestamp(ctx, tx, templateTimestamp)
		if err != nil {
			return internal.ZeroTimestamp, err
		}
	}

	log.V("Done deleting timestamp: %v", timestamp.ID)
	return deletedTimestamp, nil
}
