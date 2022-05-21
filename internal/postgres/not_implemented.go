package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
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

// Timestamp

func findTimestamps(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) ([]internal.Timestamp, error) {
	return []internal.Timestamp{}, internal.NewNotImplemented("findTimestamps")
}

func findTimestamp(ctx context.Context, tx internal.Tx, filter internal.TimestampsFilter) (internal.Timestamp, error) {
	all, err := findTimestamps(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestamp, err
	} else if len(all) == 0 {
		return internal.ZeroTimestamp, internal.NewNotFound("Timestamp", "findTimestamp")
	}
	return all[0], nil
}

func createTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, createdBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("createTimestamp")
}

func updateTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, updatedBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("updateTimestamp")
}

func deleteTimestamp(ctx context.Context, tx internal.Tx, timestamp internal.Timestamp, deletedBy uuid.UUID) (internal.Timestamp, error) {
	return timestamp, internal.NewNotImplemented("deleteTimestamp")
}
