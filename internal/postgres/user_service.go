package postgres

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
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
