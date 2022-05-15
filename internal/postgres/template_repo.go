package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	uuid "github.com/gofrs/uuid"
)

func deleteCascadeTemplate(ctx context.Context, tx internal.Tx, template internal.Template, deletedBy uuid.UUID) (internal.Template, error) {
	log.V("Deleting template: %v", template.ID)
	deletedTemplate, err := deleteTemplate(ctx, tx, template, deletedBy)
	if err != nil {
		return internal.Template{}, err
	}

	log.V("Deleting template timestamps")
	templateTimestamps, err := findTemplateTimestamps(ctx, tx, internal.TemplateTimestampsFilter{
		TemplateID: template.ID,
	})
	if err != nil {
		return internal.Template{}, err
	}
	for _, templateTimestamp := range templateTimestamps {
		_, err := deleteCascadeTemplateTimestamp(ctx, tx, templateTimestamp)
		if err != nil {
			return internal.Template{}, err
		}
	}

	log.V("Done deleting template: %v", template.ID)
	return deletedTemplate, err
}
