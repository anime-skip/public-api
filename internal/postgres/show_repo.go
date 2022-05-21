package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
	uuid "github.com/gofrs/uuid"
)

func findShows(ctx context.Context, tx internal.Tx, filter internal.ShowsFilter) ([]internal.Show, error) {
	var scanned internal.Show
	query := sqlbuilder.Select("shows", map[string]any{
		"id":                 &scanned.ID,
		"created_at":         &scanned.CreatedAt,
		"created_by_user_id": &scanned.CreatedByUserID,
		"updated_at":         &scanned.UpdatedAt,
		"updated_by_user_id": &scanned.UpdatedByUserID,
		"deleted_at":         &scanned.DeletedAt,
		"deleted_by_user_id": &scanned.DeletedByUserID,
		"name":               &scanned.Name,
		"original_name":      &scanned.OriginalName,
		"website":            &scanned.Website,
		"image":              &scanned.Image,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Name != nil {
		query.Where("name = ?", *filter.Name)
	}
	if filter.NameContains != nil {
		query.Where("name ILIKE ?", "%"+*filter.NameContains+"%")
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("name", filter.Sort)

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findShows", err)
	}
	dest := query.ScanDest()
	result := make([]internal.Show, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findShows", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findShow(ctx context.Context, tx internal.Tx, filter internal.ShowsFilter) (internal.Show, error) {
	all, err := findShows(ctx, tx, filter)
	if err != nil {
		return internal.ZeroShow, err
	} else if len(all) == 0 {
		return internal.ZeroShow, internal.NewNotFound("Show", "findShow")
	}
	return all[0], nil
}

func countShowSeasons(ctx context.Context, tx internal.Tx, id uuid.UUID) (int, error) {
	return 0, internal.NewNotImplemented("countShowSeasons")
}

func createShow(ctx context.Context, tx internal.Tx, show internal.Show, createdBy uuid.UUID) (internal.Show, error) {
	id, err := utils.RandomID()
	if err != nil {
		return show, err
	}
	show.ID = id
	show.CreatedAt = *now()
	show.CreatedByUserID = &createdBy
	show.UpdatedAt = *now()
	show.UpdatedByUserID = &createdBy

	sql, args := sqlbuilder.Insert("shows", map[string]any{
		"id":                 show.ID,
		"created_at":         show.CreatedAt,
		"created_by_user_id": show.CreatedByUserID,
		"updated_at":         show.UpdatedAt,
		"updated_by_user_id": show.UpdatedByUserID,
		"name":               show.Name,
		"original_name":      show.OriginalName,
		"website":            show.Website,
		"image":              show.Image,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if isConflict(err) {
		return show, &internal.Error{
			Code:    internal.ECONFLICT,
			Message: "Show with the generated UUID already exists, try again",
			Op:      "createShow",
			Err:     err,
		}
	} else if err != nil {
		return show, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to create Show",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}
	return show, nil
}

func updateShow(ctx context.Context, tx internal.Tx, show internal.Show, updatedBy uuid.UUID) (internal.Show, error) {
	show.UpdatedAt = *now()
	show.UpdatedByUserID = &updatedBy

	sql, args := sqlbuilder.Update("shows", show.ID, map[string]any{
		"updated_at":         show.UpdatedAt,
		"updated_by_user_id": show.UpdatedByUserID,
		"deleted_at":         show.DeletedAt,
		"deleted_by_user_id": show.DeletedByUserID,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return show, &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to update Show",
			Op:      "sqlite.createApp",
			Err:     err,
		}
	}

	return show, nil
}

func deleteShow(ctx context.Context, tx internal.Tx, show internal.Show, deletedBy uuid.UUID) (internal.Show, error) {
	show.DeletedByUserID = &deletedBy
	show.DeletedAt = now()
	return updateShow(ctx, tx, show, deletedBy)
}

func deleteCascadeShow(ctx context.Context, tx internal.Tx, show internal.Show, deletedBy uuid.UUID) (internal.Show, error) {
	log.V("Deleting show: %v", show.ID)
	deletedShow, err := deleteShow(ctx, tx, show, deletedBy)
	if err != nil {
		return internal.Show{}, err
	}

	log.V("Deleting show admins")
	admins, err := findShowAdmins(ctx, tx, internal.ShowAdminsFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, admin := range admins {
		_, err := deleteCascadeShowAdmin(ctx, tx, admin, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show templates")
	templates, err := findTemplates(ctx, tx, internal.TemplatesFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, template := range templates {
		_, err := deleteCascadeTemplate(ctx, tx, template, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Deleting show episodes")
	episodes, err := findEpisodes(ctx, tx, internal.EpisodesFilter{
		ShowID: show.ID,
	})
	if err != nil {
		return internal.Show{}, err
	}
	for _, episode := range episodes {
		_, err := deleteCascadeEpisode(ctx, tx, episode, deletedBy)
		if err != nil {
			return internal.Show{}, err
		}
	}

	log.V("Done deleting show: %v", show.ID)
	return deletedShow, nil
}
