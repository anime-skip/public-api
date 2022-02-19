package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
)

type Resolver struct {
	// Give resolvers access to all the services
	*internal.Services
}

func (r *accountResolver) AdminOfShows(ctx context.Context, obj *graphql.Account) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *accountResolver) Preferences(ctx context.Context, obj *graphql.Account) (*graphql.Preferences, error) {
	panic("not implemented")
}

func (r *episodeResolver) CreatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	panic("not implemented")
}

func (r *episodeResolver) UpdatedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	panic("not implemented")
}

func (r *episodeResolver) DeletedBy(ctx context.Context, obj *graphql.Episode) (*graphql.User, error) {
	panic("not implemented")
}

func (r *episodeResolver) Show(ctx context.Context, obj *graphql.Episode) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *episodeResolver) Timestamps(ctx context.Context, obj *graphql.Episode) ([]*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *episodeResolver) Urls(ctx context.Context, obj *graphql.Episode) ([]*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *episodeResolver) Template(ctx context.Context, obj *graphql.Episode) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *episodeUrlResolver) CreatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	panic("not implemented")
}

func (r *episodeUrlResolver) UpdatedBy(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.User, error) {
	panic("not implemented")
}

func (r *episodeUrlResolver) Episode(ctx context.Context, obj *graphql.EpisodeURL) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*graphql.LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context, recaptchaResponse string) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*graphql.Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) RequestPasswordReset(ctx context.Context, recaptchaResponse string, email string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) ResetPassword(ctx context.Context, passwordResetToken string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, passwordHash string) (*graphql.Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, deleteToken string) (*graphql.Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) SavePreferences(ctx context.Context, preferences graphql.InputPreferences) (*graphql.Preferences, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateShow(ctx context.Context, showInput graphql.InputShow, becomeAdmin bool) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID string, newShow graphql.InputShow) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID string) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput graphql.InputShowAdmin) (*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID string) (*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID string, episodeInput graphql.InputEpisode) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID string, newEpisode graphql.InputEpisode) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID string) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID string, episodeURLInput graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID string, timestampInput graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID string, newTimestamp graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID string) (*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestamps(ctx context.Context, create []*graphql.InputTimestampOn, update []*graphql.InputExistingTimestamp, delete []string) (*graphql.UpdatedTimestamps, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTimestampType(ctx context.Context, timestampTypeInput graphql.InputTimestampType) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID string, newTimestampType graphql.InputTimestampType) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID string) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTemplate(ctx context.Context, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID string, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID string) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("not implemented")
}

func (r *preferencesResolver) User(ctx context.Context, obj *graphql.Preferences) (*graphql.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Login(ctx context.Context, usernameEmail string, passwordHash string) (*graphql.LoginData, error) {
	panic("not implemented")
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*graphql.LoginData, error) {
	panic("not implemented")
}

func (r *queryResolver) FindUser(ctx context.Context, userID string) (*graphql.User, error) {
	panic("not implemented")
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*graphql.User, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShow(ctx context.Context, showID string) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*graphql.Show, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID string) (*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID string) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID string) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*graphql.Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID string) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID string) ([]*graphql.Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *string, offset *int, limit *int, sort *string) ([]*graphql.Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeByName(ctx context.Context, name string) ([]*graphql.ThirdPartyEpisode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID string) ([]*graphql.EpisodeURL, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID string) (*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID string) ([]*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID string) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplate(ctx context.Context, templateID string) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID string) ([]*graphql.Template, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *string, showName *string, season *string) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *showResolver) CreatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showResolver) UpdatedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showResolver) DeletedBy(ctx context.Context, obj *graphql.Show) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showResolver) Admins(ctx context.Context, obj *graphql.Show) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

func (r *showResolver) Episodes(ctx context.Context, obj *graphql.Show) ([]*graphql.Episode, error) {
	panic("not implemented")
}

func (r *showResolver) Templates(ctx context.Context, obj *graphql.Show) ([]*graphql.Template, error) {
	panic("not implemented")
}

func (r *showResolver) SeasonCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("not implemented")
}

func (r *showResolver) EpisodeCount(ctx context.Context, obj *graphql.Show) (int, error) {
	panic("not implemented")
}

func (r *showAdminResolver) CreatedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showAdminResolver) UpdatedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showAdminResolver) DeletedBy(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	panic("not implemented")
}

func (r *showAdminResolver) Show(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *showAdminResolver) User(ctx context.Context, obj *graphql.ShowAdmin) (*graphql.User, error) {
	panic("not implemented")
}

func (r *templateResolver) CreatedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	panic("not implemented")
}

func (r *templateResolver) UpdatedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	panic("not implemented")
}

func (r *templateResolver) DeletedBy(ctx context.Context, obj *graphql.Template) (*graphql.User, error) {
	panic("not implemented")
}

func (r *templateResolver) Show(ctx context.Context, obj *graphql.Template) (*graphql.Show, error) {
	panic("not implemented")
}

func (r *templateResolver) SourceEpisode(ctx context.Context, obj *graphql.Template) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *templateResolver) Timestamps(ctx context.Context, obj *graphql.Template) ([]*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *templateResolver) TimestampIds(ctx context.Context, obj *graphql.Template) ([]string, error) {
	panic("not implemented")
}

func (r *templateTimestampResolver) Template(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Template, error) {
	panic("not implemented")
}

func (r *templateTimestampResolver) Timestamp(ctx context.Context, obj *graphql.TemplateTimestamp) (*graphql.Timestamp, error) {
	panic("not implemented")
}

func (r *thirdPartyEpisodeResolver) Timestamps(ctx context.Context, obj *graphql.ThirdPartyEpisode) ([]*graphql.ThirdPartyTimestamp, error) {
	panic("not implemented")
}

func (r *thirdPartyEpisodeResolver) Show(ctx context.Context, obj *graphql.ThirdPartyEpisode) (*graphql.ThirdPartyShow, error) {
	panic("not implemented")
}

func (r *thirdPartyTimestampResolver) Type(ctx context.Context, obj *graphql.ThirdPartyTimestamp) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *timestampResolver) CreatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	panic("not implemented")
}

func (r *timestampResolver) UpdatedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	panic("not implemented")
}

func (r *timestampResolver) DeletedBy(ctx context.Context, obj *graphql.Timestamp) (*graphql.User, error) {
	panic("not implemented")
}

func (r *timestampResolver) Type(ctx context.Context, obj *graphql.Timestamp) (*graphql.TimestampType, error) {
	panic("not implemented")
}

func (r *timestampResolver) Episode(ctx context.Context, obj *graphql.Timestamp) (*graphql.Episode, error) {
	panic("not implemented")
}

func (r *timestampTypeResolver) CreatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	panic("not implemented")
}

func (r *timestampTypeResolver) UpdatedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	panic("not implemented")
}

func (r *timestampTypeResolver) DeletedBy(ctx context.Context, obj *graphql.TimestampType) (*graphql.User, error) {
	panic("not implemented")
}

func (r *userResolver) AdminOfShows(ctx context.Context, obj *graphql.User) ([]*graphql.ShowAdmin, error) {
	panic("not implemented")
}

// Account returns graphql.AccountResolver implementation.
func (r *Resolver) Account() graphql.AccountResolver { return &accountResolver{r} }

// Episode returns graphql.EpisodeResolver implementation.
func (r *Resolver) Episode() graphql.EpisodeResolver { return &episodeResolver{r} }

// EpisodeUrl returns graphql.EpisodeUrlResolver implementation.
func (r *Resolver) EpisodeUrl() graphql.EpisodeUrlResolver { return &episodeUrlResolver{r} }

// Mutation returns graphql.MutationResolver implementation.
func (r *Resolver) Mutation() graphql.MutationResolver { return &mutationResolver{r} }

// Preferences returns graphql.PreferencesResolver implementation.
func (r *Resolver) Preferences() graphql.PreferencesResolver { return &preferencesResolver{r} }

// Query returns graphql.QueryResolver implementation.
func (r *Resolver) Query() graphql.QueryResolver { return &queryResolver{r} }

// Show returns graphql.ShowResolver implementation.
func (r *Resolver) Show() graphql.ShowResolver { return &showResolver{r} }

// ShowAdmin returns graphql.ShowAdminResolver implementation.
func (r *Resolver) ShowAdmin() graphql.ShowAdminResolver { return &showAdminResolver{r} }

// Template returns graphql.TemplateResolver implementation.
func (r *Resolver) Template() graphql.TemplateResolver { return &templateResolver{r} }

// TemplateTimestamp returns graphql.TemplateTimestampResolver implementation.
func (r *Resolver) TemplateTimestamp() graphql.TemplateTimestampResolver {
	return &templateTimestampResolver{r}
}

// ThirdPartyEpisode returns graphql.ThirdPartyEpisodeResolver implementation.
func (r *Resolver) ThirdPartyEpisode() graphql.ThirdPartyEpisodeResolver {
	return &thirdPartyEpisodeResolver{r}
}

// ThirdPartyTimestamp returns graphql.ThirdPartyTimestampResolver implementation.
func (r *Resolver) ThirdPartyTimestamp() graphql.ThirdPartyTimestampResolver {
	return &thirdPartyTimestampResolver{r}
}

// Timestamp returns graphql.TimestampResolver implementation.
func (r *Resolver) Timestamp() graphql.TimestampResolver { return &timestampResolver{r} }

// TimestampType returns graphql.TimestampTypeResolver implementation.
func (r *Resolver) TimestampType() graphql.TimestampTypeResolver { return &timestampTypeResolver{r} }

// User returns graphql.UserResolver implementation.
func (r *Resolver) User() graphql.UserResolver { return &userResolver{r} }

type accountResolver struct{ *Resolver }
type episodeResolver struct{ *Resolver }
type episodeUrlResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type preferencesResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type showResolver struct{ *Resolver }
type showAdminResolver struct{ *Resolver }
type templateResolver struct{ *Resolver }
type templateTimestampResolver struct{ *Resolver }
type thirdPartyEpisodeResolver struct{ *Resolver }
type thirdPartyTimestampResolver struct{ *Resolver }
type timestampResolver struct{ *Resolver }
type timestampTypeResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
