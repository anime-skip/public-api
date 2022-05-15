package mappers

import (
	"anime-skip.com/public-api/internal"
)

func ToAccount(user internal.FullUser) internal.Account {
	return internal.Account{
		ID:            &user.ID,
		CreatedAt:     user.CreatedAt,
		DeletedAt:     user.DeletedAt,
		Username:      user.Username,
		Email:         user.Email,
		ProfileURL:    user.ProfileURL,
		EmailVerified: user.EmailVerified,
		Role:          ToRoleEnum(user.Role),
	}
}

func ToUser(user internal.FullUser) internal.User {
	return internal.User{
		ID:         &user.ID,
		CreatedAt:  user.CreatedAt,
		DeletedAt:  user.DeletedAt,
		Username:   user.Username,
		ProfileURL: user.ProfileURL,
	}
}
