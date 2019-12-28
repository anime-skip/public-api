package resolvers

// import (
// 	"context"

// 	gql "github.com/aklinker1/anime-skip-backend/internal/graphql"
// 	"github.com/aklinker1/anime-skip-backend/internal/graphql/models"
// )

// // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// type Resolver struct{}

// func (r *Resolver) Episode() gql.EpisodeResolver {
// 	return &episodeResolver{r}
// }
// func (r *Resolver) EpisodeUrl() gql.EpisodeUrlResolver {
// 	return &episodeUrlResolver{r}
// }
// func (r *Resolver) Mutation() gql.MutationResolver {
// 	return &mutationResolver{r}
// }
// func (r *Resolver) MyUser() gql.MyUserResolver {
// 	return &myUserResolver{r}
// }
// func (r *Resolver) Preferences() gql.PreferencesResolver {
// 	return &preferencesResolver{r}
// }
// func (r *Resolver) Query() gql.QueryResolver {
// 	return &queryResolver{r}
// }
// func (r *Resolver) Show() gql.ShowResolver {
// 	return &showResolver{r}
// }
// func (r *Resolver) ShowAdmin() gql.ShowAdminResolver {
// 	return &showAdminResolver{r}
// }
// func (r *Resolver) Timestamp() gql.TimestampResolver {
// 	return &timestampResolver{r}
// }
// func (r *Resolver) User() gql.UserResolver {
// 	return &userResolver{r}
// }

// type episodeResolver struct{ *Resolver }

// func (r *episodeResolver) CreatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *episodeResolver) DeletedBy(ctx context.Context, obj *models.Episode) (*models.User, error) {
// 	panic("not implemented")
// }

// type episodeUrlResolver struct{ *Resolver }

// func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *models.EpisodeURL) (*models.User, error) {
// 	panic("not implemented")
// }

// type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (bool, error) {
// 	panic("not implemented")
// }
// func (r *mutationResolver) SavePreferences(ctx context.Context, id string, preferences models.InputPreferences) (*models.Preferences, error) {
// 	panic("not implemented")
// }

// type myUserResolver struct{ *Resolver }

// func (r *myUserResolver) AdminOfShows(ctx context.Context, obj *models.MyUser) ([]*models.ShowAdmin, error) {
// 	panic("not implemented")
// }
// func (r *myUserResolver) Preferences(ctx context.Context, obj *models.MyUser) (*models.Preferences, error) {
// 	panic("not implemented")
// }

// type preferencesResolver struct{ *Resolver }

// func (r *preferencesResolver) User(ctx context.Context, obj *models.Preferences) (*models.User, error) {
// 	panic("not implemented")
// }

// type queryResolver struct{ *Resolver }

// func (r *queryResolver) MyUser(ctx context.Context) (*models.MyUser, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindUserByID(ctx context.Context, userID string) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*models.User, error) {
// 	panic("not implemented")
// }

// type showResolver struct{ *Resolver }

// func (r *showResolver) CreatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showResolver) UpdatedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showResolver) DeletedBy(ctx context.Context, obj *models.Show) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showResolver) Admins(ctx context.Context, obj *models.Show) ([]*models.ShowAdmin, error) {
// 	panic("not implemented")
// }
// func (r *showResolver) Episodes(ctx context.Context, obj *models.Show) ([]*models.Episode, error) {
// 	panic("not implemented")
// }

// type showAdminResolver struct{ *Resolver }

// func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *showAdminResolver) Show(ctx context.Context, obj *models.ShowAdmin) (*models.Show, error) {
// 	panic("not implemented")
// }
// func (r *showAdminResolver) User(ctx context.Context, obj *models.ShowAdmin) (*models.User, error) {
// 	panic("not implemented")
// }

// type timestampResolver struct{ *Resolver }

// func (r *timestampResolver) CreatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *timestampResolver) DeletedBy(ctx context.Context, obj *models.Timestamp) (*models.User, error) {
// 	panic("not implemented")
// }
// func (r *timestampResolver) Type(ctx context.Context, obj *models.Timestamp) (*models.TimestampType, error) {
// 	panic("not implemented")
// }

// type userResolver struct{ *Resolver }

// func (r *userResolver) AdminOfShows(ctx context.Context, obj *models.User) ([]*models.ShowAdmin, error) {
// 	panic("not implemented")
// }
