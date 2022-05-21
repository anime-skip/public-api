package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/postgres/sqlbuilder"
	"anime-skip.com/public-api/internal/utils"
)

func findUsers(ctx context.Context, tx internal.Tx, filter internal.UsersFilter) ([]internal.FullUser, error) {
	var scanned internal.FullUser
	query := sqlbuilder.Select("users", map[string]any{
		"id":             &scanned.ID,
		"created_at":     &scanned.CreatedAt,
		"deleted_at":     &scanned.DeletedAt,
		"username":       &scanned.Username,
		"email":          &scanned.Email,
		"password_hash":  &scanned.PasswordHash,
		"profile_url":    &scanned.ProfileURL,
		"email_verified": &scanned.EmailVerified,
		"role":           &scanned.Role,
	})
	if filter.IncludeDeleted {
		query.IncludeSoftDeleted()
	}
	if filter.ID != nil {
		query.Where("id = ?", *filter.ID)
	}
	if filter.Username != nil {
		query.Where("username = ?", *filter.Username)
	}
	if filter.Email != nil {
		query.Where("email = ?", *filter.Email)
	}
	if filter.UsernameOrEmail != nil {
		query.Where("(username = ? OR email = ?)", *filter.Username, *filter.Email)
	}
	if filter.Pagination != nil {
		query.Paginate(*filter.Pagination)
	}
	query.OrderBy("username", "ASC")

	sql, args := query.ToSQL()
	rows, err := tx.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, internal.SQLFailure("findUsers", err)
	}
	dest := query.ScanDest()
	result := make([]internal.FullUser, 0)
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return nil, internal.SQLFailure("findUsers", err)
		}
		result = append(result, scanned)
	}
	return result, rows.Err()
}

func findUser(ctx context.Context, tx internal.Tx, filter internal.UsersFilter) (internal.FullUser, error) {
	all, err := findUsers(ctx, tx, filter)
	if err != nil {
		return internal.ZeroFullUser, err
	} else if len(all) == 0 {
		return internal.ZeroFullUser, internal.NewNotFound("User", "findUser")
	}
	return all[0], nil
}

func createUser(ctx context.Context, tx internal.Tx, user internal.FullUser) (internal.FullUser, error) {
	id, err := utils.RandomID()
	if err != nil {
		return user, err
	}
	user.ID = *id
	user.CreatedAt = *now()

	sql, args := sqlbuilder.Insert("users", map[string]any{
		"id":             user.ID,
		"created_at":     user.CreatedAt,
		"deleted_at":     user.DeletedAt,
		"username":       user.Username,
		"email":          user.Email,
		"password_hash":  user.PasswordHash,
		"profile_url":    user.ProfileURL,
		"email_verified": user.EmailVerified,
		"role":           user.Role,
	}).ToSQL()

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return user, internal.SQLFailure("createUser", err)
	}
	return user, nil
}

func updateUser(ctx context.Context, tx internal.Tx, user internal.FullUser) (internal.FullUser, error) {
	sql, args := sqlbuilder.Update("users", user.ID, map[string]any{
		"deleted_at":     user.DeletedAt,
		"username":       user.Username,
		"email":          user.Email,
		"password_hash":  user.PasswordHash,
		"profile_url":    user.ProfileURL,
		"email_verified": user.EmailVerified,
		"role":           user.Role,
	}).ToSQL()

	_, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		return user, internal.SQLFailure("updateUser", err)
	}
	return user, nil
}
