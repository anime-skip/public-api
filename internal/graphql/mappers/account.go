package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

func ToGraphqlAccount(user internal.User) graphql.Account {
	return graphql.Account{
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
