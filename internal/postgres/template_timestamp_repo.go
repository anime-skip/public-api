package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/log"
)

func deleteCascadeTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	log.V("Deleting template timestamp (nothing to cascade): template=%v, timestamp=%v", templateTimestamp.TemplateID, templateTimestamp.TimestampID)
	return templateTimestamp, deleteTemplateTimestampInTx(ctx, tx, templateTimestamp)
}
