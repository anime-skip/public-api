package mappers

import (
	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
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
