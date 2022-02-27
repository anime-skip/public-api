package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/errors"
	"github.com/gofrs/uuid"
)

type userService struct {
	db internal.Database
}

func NewUserService(db internal.Database) internal.UserService {
	return &userService{db}
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (internal.User, error) {
	return getUserByID(ctx, s.db, id)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (internal.User, error) {
	return getUserByUsername(ctx, s.db, username)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (internal.User, error) {
	return getUserByEmail(ctx, s.db, email)
}

func (s *userService) GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (internal.User, error) {
	user, err := getUserByUsername(ctx, s.db, usernameOrEmail)
	if err == nil {
		return user, nil
	}
	if !errors.IsRecordNotFound(err) {
		return internal.User{}, err
	}
	return getUserByEmail(ctx, s.db, usernameOrEmail)
}

func (s *userService) CreateInTx(ctx context.Context, tx internal.Tx, user internal.User) (internal.User, error) {
	return insertUserInTx(ctx, tx, user)
}

func (s *userService) Update(ctx context.Context, newUser internal.User) (internal.User, error) {
	return updateUser(ctx, s.db, newUser)
}
