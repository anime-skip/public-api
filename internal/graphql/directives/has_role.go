package directives

import (
	"context"
	"log"

	gql "anime-skip.com/timestamps-service/internal/graphql"
	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role gql.Role) (interface{}, error) {
	log.Panic("Not implemented")
	return nil, nil
}
