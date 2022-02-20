package postgres

import (
	"context"
	"database/sql"

	"anime-skip.com/timestamps-service/internal"
	"github.com/jmoiron/sqlx"
)

func getUserByID(ctx context.Context, tx *sqlx.Tx, params internal.GetUserByIDParams) (internal.User, error) {
	var user internal.User
	err := tx.GetContext(ctx, &user, `SELECT * FROM users WHERE id=$1`, params.UserID)
	return user, err
}

type userService struct {
	db internal.Database
}

func NewUserService(db internal.Database) internal.UserService {
	return &userService{db}
}

func (s *userService) GetUserByID(ctx context.Context, params internal.GetUserByIDParams) (internal.User, error) {
	tx := s.db.MustBeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	user, err := getUserByID(ctx, tx, params)
	if err != nil {
		return internal.User{}, err
	}

	tx.Commit()
	return user, nil
}
