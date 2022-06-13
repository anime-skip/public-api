package postgres

import (
	"context"
	"database/sql/driver"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
	"github.com/lib/pq"
)

func findTemplates(ctx context.Context, tx internal.Tx, filter internal.TemplatesFilter) ([]internal.Template, error) {
	var scanned internal.Template
	query := sqlbuilder.Select("templates", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"show_id":            &scanned.ShowID,
		"source_episode_id":  &scanned.SourceEpisodeID,
		"\"type\"":           &scanned.Type,
		"seasons":            pq.Array(&scanned.Seasons),
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Season != nil {
		query.Where("season = ?", *filter.Season)
	}
	if filter.ShowID != nil {
		query.Where("show_id = ?", *filter.ShowID)
	}
	if filter.SourceEpisodeID != nil {
		query.Where("source_episode_id = ?", *filter.SourceEpisodeID)
	}
	if filter.Type != nil {
		query.Where("type = ?", *filter.Type)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("type", "ASC")

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findTemplates", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Template, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findTemplates", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
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
	id, err := utils.RandomID()
	if err != nil {
		return template, err
	}
	template.ID = id
	template.CreatedAt = *now()
	template.CreatedByUserID = &createdBy
	template.UpdatedAt = *now()
	template.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("templates", map[string]any{
		"id":                 template.ID,
		"created_at":         template.CreatedAt,
		"created_by_user_id": template.CreatedByUserID,
		"updated_at":         template.UpdatedAt,
		"updated_by_user_id": template.UpdatedByUserID,
		"show_id":            template.ShowID,
		"source_episode_id":  template.SourceEpisodeID,
		"type":               driver.Valuer(&template.Type),
		"seasons":            pq.Array(template.Seasons),
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return template, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "Template with the generated UUID already exists, try again",
			Op:      "createTemplate",
			Err:     err,
		}
	} else if err != nil {
		return template, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Template",
			Op:      "createTemplate",
			Err:     err,
		}
	}
	return template, nil
}

func updateTemplate(ctx context.Context, tx internal.Tx, template internal.Template, updatedBy uuid.UUID) (internal.Template, error) {
	template.UpdatedAt = *now()
	template.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("templates", template.ID, map[string]any{
		"updated_at":         template.UpdatedAt,
		"updated_by_user_id": template.UpdatedByUserID,
		"deleted_at":         template.DeletedAt,
		"deleted_by_user_id": template.DeletedByUserID,
		"show_id":            template.ShowID,
		"source_episode_id":  template.SourceEpisodeID,
		"type":               driver.Valuer(&template.Type),
		"seasons":            pq.Array(template.Seasons),
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return template, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Template",
			Op:      "updateTemplate",
			Err:     err,
		}
	}

	return template, nil
}

func deleteTemplate(ctx context.Context, tx internal.Tx, template internal.Template, deletedBy uuid.UUID) (internal.Template, error) {
	template.DeletedByUserID = &deletedBy
	template.DeletedAt = now()
	return updateTemplate(ctx, tx, template, deletedBy)
}

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
