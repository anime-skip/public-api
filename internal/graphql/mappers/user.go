package mappers

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/graphql"
)

func ToGraphqlUser(user internal.User) graphql.User {
	return graphql.User{
		ID:         &user.ID,
		CreatedAt:  user.CreatedAt,
		DeletedAt:  user.DeletedAt,
		Username:   user.Username,
		ProfileURL: user.ProfileURL,
	}
}
