package postgres

import (
	"context"

	internal "anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
)

func findExternalLinks(ctx context.Context, tx internal.Tx, filter internal.ExternalLinksFilter) ([]internal.ExternalLink, error) {
	var scanned internal.ExternalLink
	query := sqlbuilder.Select("external_links", map[string]any{
		"url":     &scanned.URL,
		"show_id": &scanned.ShowID,
	})
	if filter.URL != nil {
		query.Where("url = ?", *filter.URL)
	}
	if filter.ShowID != nil {
		query.Where("show_id = ?", *filter.ShowID)
	}
	if filter.ServiceID != nil {
		match, err := filter.ServiceID.Service.URLMatcher(filter.ServiceID.ServiceID)
		if err != nil {
			return nil, err
		}
		query.Where("url ILIKE ?", match)
	}

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findExternalLinks", err)
	}
	dest := query.ScanDest()
	result := make([]internal.ExternalLink, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findExternalLinks", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findExternalLink(ctx context.Context, tx internal.Tx, filter internal.ExternalLinksFilter) (internal.ExternalLink, error) {
	all, err := findExternalLinks(ctx, tx, filter)
	if err != nil {
		return internal.ZeroExternalLink, err
	} else if len(all) == 0 {
		return internal.ZeroExternalLink, internal.NewNotFound("ExternalLink", "findExternalLink")
	}
	return all[0], nil
}

func createExternalLink(ctx context.Context, tx internal.Tx, externalLink internal.ExternalLink) (internal.ExternalLink, error) {
	sql, args := sqlbuilder.Insert("external_links", map[string]any{
		"url":     externalLink.URL,
		"show_id": externalLink.ShowID,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return externalLink, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "ExternalLink with the generated UUID already exists, try again",
			Op:      "createExternalLink",
			Err:     err,
		}
	} else if err != nil {
		return externalLink, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create ExternalLink",
			Op:      "createExternalLink",
			Err:     err,
		}
	}
	return externalLink, nil
}

func deleteExternalLink(ctx context.Context, tx internal.Tx, externalLink internal.ExternalLink) (internal.ExternalLink, error) {
	_, err := tx.ExecContext(ctx, "DELETE FROM external_links WHERE url = $1", externalLink.URL)
	return externalLink, err
}

func deleteCascadeExternalLink(ctx context.Context, tx internal.Tx, externalLink internal.ExternalLink) (internal.ExternalLink, error) {
	log.V("Deleting external link (nothing to cascade): %s", externalLink.URL)
	return deleteExternalLink(ctx, tx, externalLink)
}
