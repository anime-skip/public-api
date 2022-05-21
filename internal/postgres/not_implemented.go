package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
)

// TemplateTimestamp

func findTemplateTimestamps(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) ([]internal.TemplateTimestamp, error) {
	return []internal.TemplateTimestamp{}, internal.NewNotImplemented("findTemplateTimestamps")
}

func findTemplateTimestamp(ctx context.Context, tx internal.Tx, filter internal.TemplateTimestampsFilter) (internal.TemplateTimestamp, error) {
	all, err := findTemplateTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTemplateTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTemplateTimestamp, internal.NewNotFound("Template timestamp", "findTemplateTimestamp")
	}
	return all[0], nil
}

func createTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("createTemplateTimestamp")
}

func updateTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("updateTemplateTimestamp")
}

func deleteTemplateTimestamp(ctx context.Context, tx internal.Tx, templateTimestamp internal.TemplateTimestamp) (internal.TemplateTimestamp, error) {
	return templateTimestamp, internal.NewNotImplemented("deleteTemplateTimestamp")
}
