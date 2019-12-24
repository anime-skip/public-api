package resolvers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/gql"
)

// Resolver stores the instance of gorm so it can be accessed in each of the resolvers
type Resolver struct {
	ORM *database.ORM
}

// Mutation returns the root mutation for the schema
func (r *Resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}

// Query returns the root query for the schema
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
