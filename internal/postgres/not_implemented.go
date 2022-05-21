package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	uuid "github.com/gofrs/uuid"
)

// EpisodeURL

func findEpisodeURLs(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) ([]internal.EpisodeURL, error) {
	return []internal.EpisodeURL{}, internal.NewNotImplemented("findEpisodeURLs")
}

func findEpisodeURL(ctx context.Context, tx internal.Tx, filter internal.EpisodeURLsFilter) (internal.EpisodeURL, error) {
	all, err := findEpisodeURLs(ctx, tx, filter)
	if err != nil {
		return internal.ZeroEpisodeURL, err
	} else if len(all) == 0 {
		return internal.ZeroEpisodeURL, internal.NewNotFound("Episode URL", "findEpisodeURL")
	}
	return all[0], nil
}

func createEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, createdBy uuid.UUID) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("createEpisodeURL")
}

func updateEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL, updatedBy uuid.UUID) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("updateEpisodeURL")
}

func deleteEpisodeURL(ctx context.Context, tx internal.Tx, episodeURL internal.EpisodeURL) (internal.EpisodeURL, error) {
	return episodeURL, internal.NewNotImplemented("deleteEpisodeURL")
}

// Template

func findTemplates(ctx context.Context, tx internal.Tx, filter internal.TemplatesFilter) ([]internal.Template, error) {
	return []internal.Template{}, internal.NewNotImplemented("findTemplates")
}

func findTemplate(ctx context.Context, tx internal.Tx, filter internal.TemplatesFilter) (internal.Template, error) {
	all, err := findTemplates(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTemplate, err
	} else if len(all) == 0 {
		return internal.ZeroTemplate, internal.NewNotFound("Template", "findTemplate")
	}
	return all[0], nil
}

func createTemplate(ctx context.Context, tx internal.Tx, template internal.Template, createdBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("createTemplate")
}

func updateTemplate(ctx context.Context, tx internal.Tx, template internal.Template, updatedBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("updateTemplate")
}

func deleteTemplate(ctx context.Context, tx internal.Tx, template internal.Template, deletedBy uuid.UUID) (internal.Template, error) {
	return template, internal.NewNotImplemented("deleteTemplate")
}

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

// TimestampType

func findTimestampTypes(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) ([]internal.TimestampType, error) {
	return []internal.TimestampType{}, internal.NewNotImplemented("findTimestampTypes")
}

func findTimestampType(ctx context.Context, tx internal.Tx, filter internal.TimestampTypesFilter) (internal.TimestampType, error) {
	all, err := findTimestampTypes(ctx, tx, filter)
	if err != nil {
		return internal.ZeroTimestampType, err
	} else if len(all) == 0 {
		return internal.ZeroTimestampType, internal.NewNotFound("Timestamp type", "findTimestampType")
	}
	return all[0], nil
}

func createTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, createdBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("createTimestampType")
}

func updateTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, updatedBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("updateTimestampType")
}

func deleteTimestampType(ctx context.Context, tx internal.Tx, timestampType internal.TimestampType, deletedBy uuid.UUID) (internal.TimestampType, error) {
	return timestampType, internal.NewNotImplemented("deleteTimestampType")
}
