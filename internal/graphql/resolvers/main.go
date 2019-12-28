package resolvers

import (
	"github.com/aklinker1/anime-skip-backend/internal/database"
	"github.com/aklinker1/anime-skip-backend/internal/graphql"
)

// Resolver stores the instance of gorm so it can be accessed in each of the resolvers
type Resolver struct {
	ORM *database.ORM
}

func (r *Resolver) Episode() gql.EpisodeResolver {
	return &episodeResolver{r}
}
func (r *Resolver) EpisodeUrl() gql.EpisodeUrlResolver {
	return &episodeUrlResolver{r}
}
func (r *Resolver) MyUser() gql.MyUserResolver {
	return &myUserResolver{r}
}
func (r *Resolver) Preferences() gql.PreferencesResolver {
	return &preferencesResolver{r}
}
func (r *Resolver) Show() gql.ShowResolver {
	return &showResolver{r}
}
func (r *Resolver) ShowAdmin() gql.ShowAdminResolver {
	return &showAdminResolver{r}
}
func (r *Resolver) Timestamp() gql.TimestampResolver {
	return &timestampResolver{r}
}
func (r *Resolver) User() gql.UserResolver {
	return &userResolver{r}
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