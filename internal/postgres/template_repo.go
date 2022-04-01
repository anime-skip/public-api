package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

func deleteCascadeTemplate(ctx context.Context, tx internal.Tx, template internal.Template) (internal.Template, error) {
	log.V("Deleting template: %v", template.ID)
	deletedTemplate, err := deleteTemplateInTx(ctx, tx, template)
	if err != nil {
		return internal.Template{}, err
	}

	log.V("Deleting template timestamps")
	templateTimestamps, err := getTemplateTimestampsByTemplateIDInTx(ctx, tx, template.ID)
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
