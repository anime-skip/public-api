package postgres

import (
	"context"
	"database/sql"

	"anime-skip.com/timestamps-service/internal"
	"github.com/jmoiron/sqlx"
)

func getPreferencesByUserID(ctx context.Context, tx *sqlx.Tx, params internal.GetPreferencesByUserIDParams) (internal.Preferences, error) {
	var preferences internal.Preferences
	err := tx.GetContext(ctx, &preferences, `SELECT * FROM preferences WHERE user_id=$1`, params.UserID)
	return preferences, err
}

type preferencesService struct {
	db internal.Database
}

func NewPreferencesService(db internal.Database) internal.PreferencesService {
	return &preferencesService{db}
}

func (s *preferencesService) GetPreferencesByUserID(ctx context.Context, params internal.GetPreferencesByUserIDParams) (internal.Preferences, error) {
	tx := s.db.MustBeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	defer tx.Rollback()

	preferences, err := getPreferencesByUserID(ctx, tx, params)
	if err != nil {
		return internal.Preferences{}, err
	}

	tx.Commit()
	return preferences, nil
}
