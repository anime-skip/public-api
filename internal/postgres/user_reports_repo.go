package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findUserReports(ctx context.Context, tx internal.Tx, filter internal.UserReportsFilter) ([]internal.UserReport, error) {
	var scanned internal.UserReport
	query := sqlbuilder.Select("user_reports", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"message":            &scanned.Message,
		"reported_from_url":  &scanned.ReportedFromURL,
		"resolved":           &scanned.Resolved,
		"timestamp_id":       &scanned.TimestampID,
		"episode_id":         &scanned.EpisodeID,
		"episode_url":        &scanned.EpisodeURLString,
		"show_id":            &scanned.ShowID,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Resolved != nil {
		if *filter.Resolved {
			query.Where("resolved = ?", true)
			query.IncludeSoftDeleted() // resolved reports are also deleted
		} else {
			query.Where("resolved = ?", false)
		}
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("created_at", filter.Sort)

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findUserReports", err)
	}
	dest := query.ScanDest()
	result := make([]internal.UserReport, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findUserReports", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findUserReport(ctx context.Context, tx internal.Tx, filter internal.UserReportsFilter) (internal.UserReport, error) {
	all, err := findUserReports(ctx, tx, filter)
	if err != nil {
		return internal.ZeroUserReport, err
	} else if len(all) == 0 {
		return internal.ZeroUserReport, internal.NewNotFound("UserReport", "findUserReport")
	}
	return all[0], nil
}

func createUserReport(ctx context.Context, tx internal.Tx, report internal.UserReport, createdBy uuid.UUID) (internal.UserReport, error) {
	id, err := utils.RandomID()
	if err != nil {
		return report, err
	}
	report.ID = id
	report.CreatedAt = *now()
	report.CreatedByUserID = &createdBy
	report.UpdatedAt = *now()
	report.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("user_reports", map[string]any{
		"id":                 report.ID,
		"created_at":         report.CreatedAt,
		"created_by_user_id": report.CreatedByUserID,
		"updated_at":         report.UpdatedAt,
		"updated_by_user_id": report.UpdatedByUserID,
		"deleted_at":         report.DeletedAt,
		"deleted_by_user_id": report.DeletedByUserID,
		"message":            report.Message,
		"reported_from_url":  report.ReportedFromURL,
		"resolved":           report.Resolved,
		"timestamp_id":       report.TimestampID,
		"episode_id":         report.EpisodeID,
		"episode_url":        report.EpisodeURLString,
		"show_id":            report.ShowID,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return report, internal.SQLFailure("createUserReport", err)
	}
	return report, nil
}

func updateUserReport(ctx context.Context, tx internal.Tx, report internal.UserReport, updatedBy uuid.UUID) (internal.UserReport, error) {
	report.UpdatedAt = *now()
	report.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("user_reports", report.ID, map[string]any{
		"id":                 report.ID,
		"updated_at":         report.UpdatedAt,
		"updated_by_user_id": report.UpdatedByUserID,
		"deleted_at":         report.DeletedAt,
		"deleted_by_user_id": report.DeletedByUserID,
		"message":            report.Message,
		"reported_from_url":  report.ReportedFromURL,
		"resolved":           report.Resolved,
		"timestamp_id":       report.TimestampID,
		"episode_id":         report.EpisodeID,
		"episode_url":        report.EpisodeURLString,
		"show_id":            report.ShowID,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return report, internal.SQLFailure("updateUserReport", err)
	}
	return report, nil
}

func deleteUserReport(ctx context.Context, tx internal.Tx, report internal.UserReport, deletedBy uuid.UUID) (internal.UserReport, error) {
	report.DeletedByUserID = &deletedBy
	report.DeletedAt = now()
	return updateUserReport(ctx, tx, report, deletedBy)
}

func deleteCascadeUserReport(ctx context.Context, tx internal.Tx, report internal.UserReport, deletedBy uuid.UUID) (internal.UserReport, error) {
	log.V("Deleting user report (nothing to cascade): %v", report.ID)
	return deleteUserReport(ctx, tx, report, deletedBy)
}
