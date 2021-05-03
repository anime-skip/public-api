package resolvers

import (
	"context"

	"anime-skip.com/backend/internal/database"
	gql "anime-skip.com/backend/internal/graphql"
	"anime-skip.com/backend/internal/utils/constants"
	"anime-skip.com/backend/internal/utils/context_utils"
	"github.com/jinzhu/gorm"
)

// Resolver stores the instance of gorm so it can be accessed in each of the resolvers
type Resolver struct {
	_orm *database.ORM
}

func ResolverWithORM(orm *database.ORM) *Resolver {
	return &Resolver{
		_orm: orm,
	}
}

func (r *Resolver) DB(ctx context.Context) *gorm.DB {
	if userID, err := context_utils.UserID(ctx); err == nil {
		r._orm.DB = r._orm.DB.Set(constants.CTX_USER_ID, userID)
	}
	return r._orm.DB
}

func (r *Resolver) Episode() gql.EpisodeResolver {
	return &episodeResolver{r}
}
func (r *Resolver) ThirdPartyEpisode() gql.ThirdPartyEpisodeResolver {
	return &thirdPartyEpisodeResolver{r}
}
func (r *Resolver) EpisodeUrl() gql.EpisodeUrlResolver {
	return &episodeUrlResolver{r}
}
func (r *Resolver) Account() gql.AccountResolver {
	return &AccountResolver{r}
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
func (r *Resolver) ThirdPartyTimestamp() gql.ThirdPartyTimestampResolver {
	return &thirdPartyTimestampResolver{r}
}
func (r *Resolver) TimestampType() gql.TimestampTypeResolver {
	return &timestampTypeResolver{r}
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
