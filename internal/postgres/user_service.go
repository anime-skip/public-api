package postgres

import (
	"context"

	"anime-skip.com/public-api/internal"
)

type userService struct {
	db internal.Database
}

func NewUserService(db internal.Database) internal.UserService {
	return &userService{db}
}

func (s *userService) Get(ctx context.Context, filter internal.UsersFilter) (internal.FullUser, error) {
	return inTx(ctx, s.db, false, internal.ZeroFullUser, func(tx internal.Tx) (internal.FullUser, error) {
		return findUser(ctx, tx, filter)
	})
}

func (s *userService) List(ctx context.Context, filter internal.UsersFilter) ([]internal.FullUser, error) {
	return inTx(ctx, s.db, false, nil, func(tx internal.Tx) ([]internal.FullUser, error) {
		return findUsers(ctx, tx, filter)
	})
}

func (s *userService) CreateAccount(ctx context.Context, newUser internal.FullUser) (internal.FullUser, error) {
	return inTx(ctx, s.db, true, internal.ZeroFullUser, func(tx internal.Tx) (internal.FullUser, error) {
		user, err := createUser(ctx, tx, newUser)
		if err != nil {
			return internal.ZeroFullUser, err
		}
		defaultPreferences := internal.NewPreferences(ctx, user.ID)
		_, err = createPreferences(ctx, tx, defaultPreferences)
		if err != nil {
			return internal.ZeroFullUser, err
		}
		return user, nil
	})
}

func (s *userService) Update(ctx context.Context, newUser internal.FullUser) (internal.FullUser, error) {
	return inTx(ctx, s.db, true, internal.ZeroFullUser, func(tx internal.Tx) (internal.FullUser, error) {
		return updateUser(ctx, tx, newUser)
	})
}

func (s *userService) Count(ctx context.Context) (int, error) {
	return inTx(ctx, s.db, false, 0, func(tx internal.Tx) (int, error) {
		return count(ctx, tx, "SELECT COUNT(*) FROM users")
	})
}
