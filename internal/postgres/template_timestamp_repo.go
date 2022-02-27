package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

func deleteCascadeTemplateTimestamp(ctx context.Context, tx internal.Tx, template internal.TemplateTimestamp) (internal.Episode, error) {
	panic("deleteCascadeTemplateTimestamp not implemented")
}
