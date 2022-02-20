package resolvers

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
	"anime-skip.com/timestamps-service/internal/graphql"
	"github.com/gofrs/uuid"
)

type Resolver struct {
	// Give resolvers access to all the services
	*internal.Services
}

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*graphql.LoginData, error) {
	panic("mutationResolver.CreateAccount not implemented")
}

func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("mutationResolver.ChangePassword not implemented")
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context, recaptchaResponse string) (*bool, error) {
	panic("mutationResolver.ResendVerificationEmail not implemented")
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*graphql.Account, error) {
	panic("mutationResolver.VerifyEmailAddress not implemented")
}

func (r *mutationResolver) RequestPasswordReset(ctx context.Context, recaptchaResponse string, email string) (bool, error) {
	panic("mutationResolver.RequestPasswordReset not implemented")
}

func (r *mutationResolver) ResetPassword(ctx context.Context, passwordResetToken string, newPassword string, confirmNewPassword string) (*graphql.LoginData, error) {
	panic("mutationResolver.ResetPassword not implemented")
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, passwordHash string) (*graphql.Account, error) {
	panic("mutationResolver.DeleteAccountRequest not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, deleteToken string) (*graphql.Account, error) {
	panic("mutationResolver.DeleteAccount not implemented")
}

func (r *mutationResolver) SavePreferences(ctx context.Context, preferences graphql.InputPreferences) (*graphql.Preferences, error) {
	panic("mutationResolver.SavePreferences not implemented")
}

func (r *mutationResolver) CreateShow(ctx context.Context, showInput graphql.InputShow, becomeAdmin bool) (*graphql.Show, error) {
	panic("mutationResolver.CreateShow not implemented")
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID *uuid.UUID, newShow graphql.InputShow) (*graphql.Show, error) {
	panic("mutationResolver.UpdateShow not implemented")
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	panic("mutationResolver.DeleteShow not implemented")
}

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput graphql.InputShowAdmin) (*graphql.ShowAdmin, error) {
	panic("mutationResolver.CreateShowAdmin not implemented")
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	panic("mutationResolver.DeleteShowAdmin not implemented")
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID *uuid.UUID, episodeInput graphql.InputEpisode) (*graphql.Episode, error) {
	panic("mutationResolver.CreateEpisode not implemented")
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID *uuid.UUID, newEpisode graphql.InputEpisode) (*graphql.Episode, error) {
	panic("mutationResolver.UpdateEpisode not implemented")
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	panic("mutationResolver.DeleteEpisode not implemented")
}

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID *uuid.UUID, episodeURLInput graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	panic("mutationResolver.CreateEpisodeURL not implemented")
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	panic("mutationResolver.DeleteEpisodeURL not implemented")
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL graphql.InputEpisodeURL) (*graphql.EpisodeURL, error) {
	panic("mutationResolver.UpdateEpisodeURL not implemented")
}

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID *uuid.UUID, timestampInput graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("mutationResolver.CreateTimestamp not implemented")
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID *uuid.UUID, newTimestamp graphql.InputTimestamp) (*graphql.Timestamp, error) {
	panic("mutationResolver.UpdateTimestamp not implemented")
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID *uuid.UUID) (*graphql.Timestamp, error) {
	panic("mutationResolver.DeleteTimestamp not implemented")
}

func (r *mutationResolver) UpdateTimestamps(ctx context.Context, create []*graphql.InputTimestampOn, update []*graphql.InputExistingTimestamp, delete []*uuid.UUID) (*graphql.UpdatedTimestamps, error) {
	panic("mutationResolver.UpdateTimestamps not implemented")
}

func (r *mutationResolver) CreateTimestampType(ctx context.Context, timestampTypeInput graphql.InputTimestampType) (*graphql.TimestampType, error) {
	panic("mutationResolver.CreateTimestampType not implemented")
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID *uuid.UUID, newTimestampType graphql.InputTimestampType) (*graphql.TimestampType, error) {
	panic("mutationResolver.UpdateTimestampType not implemented")
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*graphql.TimestampType, error) {
	panic("mutationResolver.DeleteTimestampType not implemented")
}

func (r *mutationResolver) CreateTemplate(ctx context.Context, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("mutationResolver.CreateTemplate not implemented")
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID *uuid.UUID, newTemplate graphql.InputTemplate) (*graphql.Template, error) {
	panic("mutationResolver.UpdateTemplate not implemented")
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID *uuid.UUID) (*graphql.Template, error) {
	panic("mutationResolver.DeleteTemplate not implemented")
}

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("mutationResolver.AddTimestampToTemplate not implemented")
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp graphql.InputTemplateTimestamp) (*graphql.TemplateTimestamp, error) {
	panic("mutationResolver.RemoveTimestampFromTemplate not implemented")
}

func (r *queryResolver) Login(ctx context.Context, usernameEmail string, passwordHash string) (*graphql.LoginData, error) {
	panic("queryResolver.Login not implemented")
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*graphql.LoginData, error) {
	panic("queryResolver.LoginRefresh not implemented")
}

func (r *queryResolver) FindUser(ctx context.Context, userID *uuid.UUID) (*graphql.User, error) {
	return r.getUserById(ctx, userID)
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*graphql.User, error) {
	panic("queryResolver.FindUserByUsername not implemented")
}

func (r *queryResolver) FindShow(ctx context.Context, showID *uuid.UUID) (*graphql.Show, error) {
	panic("queryResolver.FindShow not implemented")
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*graphql.Show, error) {
	panic("queryResolver.SearchShows not implemented")
}

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID *uuid.UUID) (*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdmin not implemented")
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdminsByShowID not implemented")
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID *uuid.UUID) ([]*graphql.ShowAdmin, error) {
	panic("queryResolver.FindShowAdminsByUserID not implemented")
}

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*graphql.Episode, error) {
	panic("queryResolver.RecentlyAddedEpisodes not implemented")
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID *uuid.UUID) (*graphql.Episode, error) {
	panic("queryResolver.FindEpisode not implemented")
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Episode, error) {
	panic("queryResolver.FindEpisodesByShowID not implemented")
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *uuid.UUID, offset *int, limit *int, sort *string) ([]*graphql.Episode, error) {
	panic("queryResolver.SearchEpisodes not implemented")
}

func (r *queryResolver) FindEpisodeByName(ctx context.Context, name string) ([]*graphql.ThirdPartyEpisode, error) {
	panic("queryResolver.FindEpisodeByName not implemented")
}

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*graphql.EpisodeURL, error) {
	panic("queryResolver.FindEpisodeURL not implemented")
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.EpisodeURL, error) {
	panic("queryResolver.FindEpisodeUrlsByEpisodeID not implemented")
}

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID *uuid.UUID) (*graphql.Timestamp, error) {
	panic("queryResolver.FindTimestamp not implemented")
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID *uuid.UUID) ([]*graphql.Timestamp, error) {
	panic("queryResolver.FindTimestampsByEpisodeID not implemented")
}

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID *uuid.UUID) (*graphql.TimestampType, error) {
	panic("queryResolver.FindTimestampType not implemented")
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*graphql.TimestampType, error) {
	panic("queryResolver.AllTimestampTypes not implemented")
}

func (r *queryResolver) FindTemplate(ctx context.Context, templateID *uuid.UUID) (*graphql.Template, error) {
	panic("queryResolver.FindTemplate not implemented")
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID *uuid.UUID) ([]*graphql.Template, error) {
	panic("queryResolver.FindTemplatesByShowID not implemented")
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *uuid.UUID, showName *string, season *string) (*graphql.Template, error) {
	panic("queryResolver.FindTemplateByDetails not implemented")
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
