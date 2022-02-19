package graphql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"anime-skip.com/timestamps-service/internal"
)

// Resolver Struct Types

type Resolver struct {
	db internal.Database
}

type accountFieldResolver struct{ *Resolver }
type episodeFieldResolver struct{ *Resolver }
type episodeUrlFieldResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type preferencesFieldResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type showFieldResolver struct{ *Resolver }
type showAdminFieldResolver struct{ *Resolver }
type templateFieldResolver struct{ *Resolver }
type templateTimestampFieldResolver struct{ *Resolver }
type thirdPartyEpisodeFieldResolver struct{ *Resolver }
type thirdPartyTimestampFieldResolver struct{ *Resolver }
type timestampFieldResolver struct{ *Resolver }
type timestampTypeFieldResolver struct{ *Resolver }
type userFieldResolver struct{ *Resolver }

// Resolver Constructors

func (r *Resolver) Account() AccountResolver {
	return &accountFieldResolver{r}
}
func (r *Resolver) Episode() EpisodeResolver {
	return &episodeFieldResolver{r}
}
func (r *Resolver) EpisodeUrl() EpisodeUrlResolver {
	return &episodeUrlFieldResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Preferences() PreferencesResolver {
	return &preferencesFieldResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Show() ShowResolver {
	return &showFieldResolver{r}
}
func (r *Resolver) ShowAdmin() ShowAdminResolver {
	return &showAdminFieldResolver{r}
}
func (r *Resolver) Template() TemplateResolver {
	return &templateFieldResolver{r}
}
func (r *Resolver) TemplateTimestamp() TemplateTimestampResolver {
	return &templateTimestampFieldResolver{r}
}
func (r *Resolver) ThirdPartyEpisode() ThirdPartyEpisodeResolver {
	return &thirdPartyEpisodeFieldResolver{r}
}
func (r *Resolver) ThirdPartyTimestamp() ThirdPartyTimestampResolver {
	return &thirdPartyTimestampFieldResolver{r}
}
func (r *Resolver) Timestamp() TimestampResolver {
	return &timestampFieldResolver{r}
}
func (r *Resolver) TimestampType() TimestampTypeResolver {
	return &timestampTypeFieldResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userFieldResolver{r}
}

// Mutations

func (r *mutationResolver) CreateAccount(ctx context.Context, username string, email string, passwordHash string, recaptchaResponse string) (*LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string, confirmNewPassword string) (*LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) ResendVerificationEmail(ctx context.Context, recaptchaResponse string) (*bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) VerifyEmailAddress(ctx context.Context, validationToken string) (*Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) RequestPasswordReset(ctx context.Context, recaptchaResponse string, email string) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) ResetPassword(ctx context.Context, passwordResetToken string, newPassword string, confirmNewPassword string) (*LoginData, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAccountRequest(ctx context.Context, passwordHash string) (*Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteAccount(ctx context.Context, deleteToken string) (*Account, error) {
	panic("not implemented")
}

func (r *mutationResolver) SavePreferences(ctx context.Context, preferences InputPreferences) (*Preferences, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateShow(ctx context.Context, showInput InputShow, becomeAdmin bool) (*Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateShow(ctx context.Context, showID string, newShow InputShow) (*Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteShow(ctx context.Context, showID string) (*Show, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateShowAdmin(ctx context.Context, showAdminInput InputShowAdmin) (*ShowAdmin, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteShowAdmin(ctx context.Context, showAdminID string) (*ShowAdmin, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, showID string, episodeInput InputEpisode) (*Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, episodeID string, newEpisode InputEpisode) (*Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteEpisode(ctx context.Context, episodeID string) (*Episode, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateEpisodeURL(ctx context.Context, episodeID string, episodeURLInput InputEpisodeURL) (*EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteEpisodeURL(ctx context.Context, episodeURL string) (*EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateEpisodeURL(ctx context.Context, episodeURL string, newEpisodeURL InputEpisodeURL) (*EpisodeURL, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTimestamp(ctx context.Context, episodeID string, timestampInput InputTimestamp) (*Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestamp(ctx context.Context, timestampID string, newTimestamp InputTimestamp) (*Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTimestamp(ctx context.Context, timestampID string) (*Timestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestamps(ctx context.Context, create []*InputTimestampOn, update []*InputExistingTimestamp, delete []string) (*UpdatedTimestamps, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTimestampType(ctx context.Context, timestampTypeInput InputTimestampType) (*TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTimestampType(ctx context.Context, timestampTypeID string, newTimestampType InputTimestampType) (*TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTimestampType(ctx context.Context, timestampTypeID string) (*TimestampType, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateTemplate(ctx context.Context, newTemplate InputTemplate) (*Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, templateID string, newTemplate InputTemplate) (*Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) DeleteTemplate(ctx context.Context, templateID string) (*Template, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddTimestampToTemplate(ctx context.Context, templateTimestamp InputTemplateTimestamp) (*TemplateTimestamp, error) {
	panic("not implemented")
}

func (r *mutationResolver) RemoveTimestampFromTemplate(ctx context.Context, templateTimestamp InputTemplateTimestamp) (*TemplateTimestamp, error) {
	panic("not implemented")
}

// Queries

func (r *queryResolver) Account(ctx context.Context) (*Account, error) {
	panic("not implemented")
}

func (r *queryResolver) Login(ctx context.Context, usernameEmail string, passwordHash string) (*LoginData, error) {
	panic("not implemented")
}

func (r *queryResolver) LoginRefresh(ctx context.Context, refreshToken string) (*LoginData, error) {
	panic("not implemented")
}

func (r *queryResolver) FindUser(ctx context.Context, userID string) (*User, error) {
	panic("not implemented")
}

func (r *queryResolver) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShow(ctx context.Context, showID string) (*Show, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchShows(ctx context.Context, search *string, offset *int, limit *int, sort *string) ([]*Show, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdmin(ctx context.Context, showAdminID string) (*ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdminsByShowID(ctx context.Context, showID string) ([]*ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) FindShowAdminsByUserID(ctx context.Context, userID string) ([]*ShowAdmin, error) {
	panic("not implemented")
}

func (r *queryResolver) RecentlyAddedEpisodes(ctx context.Context, limit *int, offset *int) ([]*Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisode(ctx context.Context, episodeID string) (*Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodesByShowID(ctx context.Context, showID string) ([]*Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) SearchEpisodes(ctx context.Context, search *string, showID *string, offset *int, limit *int, sort *string) ([]*Episode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeByName(ctx context.Context, name string) ([]*ThirdPartyEpisode, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeURL(ctx context.Context, episodeURL string) (*EpisodeURL, error) {
	panic("not implemented")
}

func (r *queryResolver) FindEpisodeUrlsByEpisodeID(ctx context.Context, episodeID string) ([]*EpisodeURL, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestamp(ctx context.Context, timestampID string) (*Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestampsByEpisodeID(ctx context.Context, episodeID string) ([]*Timestamp, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTimestampType(ctx context.Context, timestampTypeID string) (*TimestampType, error) {
	panic("not implemented")
}

func (r *queryResolver) AllTimestampTypes(ctx context.Context) ([]*TimestampType, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplate(ctx context.Context, templateID string) (*Template, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplatesByShowID(ctx context.Context, showID string) ([]*Template, error) {
	panic("not implemented")
}

func (r *queryResolver) FindTemplateByDetails(ctx context.Context, episodeID *string, showName *string, season *string) (*Template, error) {
	panic("not implemented")
}

// Account Fields

func (r *accountFieldResolver) AdminOfShows(ctx context.Context, obj *Account) ([]*ShowAdmin, error) {
	panic("not implemented")
}

func (r *accountFieldResolver) Preferences(ctx context.Context, obj *Account) (*Preferences, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) CreatedBy(ctx context.Context, obj *Episode) (*User, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) UpdatedBy(ctx context.Context, obj *Episode) (*User, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) DeletedBy(ctx context.Context, obj *Episode) (*User, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) Show(ctx context.Context, obj *Episode) (*Show, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) Timestamps(ctx context.Context, obj *Episode) ([]*Timestamp, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) Urls(ctx context.Context, obj *Episode) ([]*EpisodeURL, error) {
	panic("not implemented")
}

func (r *episodeFieldResolver) Template(ctx context.Context, obj *Episode) (*Template, error) {
	panic("not implemented")
}

// Episode URL Fields

func (r *episodeUrlFieldResolver) CreatedBy(ctx context.Context, obj *EpisodeURL) (*User, error) {
	panic("not implemented")
}

func (r *episodeUrlFieldResolver) UpdatedBy(ctx context.Context, obj *EpisodeURL) (*User, error) {
	panic("not implemented")
}

func (r *episodeUrlFieldResolver) Episode(ctx context.Context, obj *EpisodeURL) (*Episode, error) {
	panic("not implemented")
}

// Preferences Fields

func (r *preferencesFieldResolver) User(ctx context.Context, obj *Preferences) (*User, error) {
	panic("not implemented")
}

// Show Fields

func (r *showFieldResolver) CreatedBy(ctx context.Context, obj *Show) (*User, error) {
	panic("not implemented")
}

func (r *showFieldResolver) UpdatedBy(ctx context.Context, obj *Show) (*User, error) {
	panic("not implemented")
}

func (r *showFieldResolver) DeletedBy(ctx context.Context, obj *Show) (*User, error) {
	panic("not implemented")
}

func (r *showFieldResolver) Admins(ctx context.Context, obj *Show) ([]*ShowAdmin, error) {
	panic("not implemented")
}

func (r *showFieldResolver) Episodes(ctx context.Context, obj *Show) ([]*Episode, error) {
	panic("not implemented")
}

func (r *showFieldResolver) Templates(ctx context.Context, obj *Show) ([]*Template, error) {
	panic("not implemented")
}

func (r *showFieldResolver) SeasonCount(ctx context.Context, obj *Show) (int, error) {
	panic("not implemented")
}

func (r *showFieldResolver) EpisodeCount(ctx context.Context, obj *Show) (int, error) {
	panic("not implemented")
}

// Show Admin Fields

func (r *showAdminFieldResolver) CreatedBy(ctx context.Context, obj *ShowAdmin) (*User, error) {
	panic("not implemented")
}

func (r *showAdminFieldResolver) UpdatedBy(ctx context.Context, obj *ShowAdmin) (*User, error) {
	panic("not implemented")
}

func (r *showAdminFieldResolver) DeletedBy(ctx context.Context, obj *ShowAdmin) (*User, error) {
	panic("not implemented")
}

func (r *showAdminFieldResolver) Show(ctx context.Context, obj *ShowAdmin) (*Show, error) {
	panic("not implemented")
}

func (r *showAdminFieldResolver) User(ctx context.Context, obj *ShowAdmin) (*User, error) {
	panic("not implemented")
}

// Template Fields

func (r *templateFieldResolver) CreatedBy(ctx context.Context, obj *Template) (*User, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) UpdatedBy(ctx context.Context, obj *Template) (*User, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) DeletedBy(ctx context.Context, obj *Template) (*User, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) Show(ctx context.Context, obj *Template) (*Show, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) SourceEpisode(ctx context.Context, obj *Template) (*Episode, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) Timestamps(ctx context.Context, obj *Template) ([]*Timestamp, error) {
	panic("not implemented")
}

func (r *templateFieldResolver) TimestampIds(ctx context.Context, obj *Template) ([]string, error) {
	panic("not implemented")
}

// Template Timestamp Fields

func (r *templateTimestampFieldResolver) Template(ctx context.Context, obj *TemplateTimestamp) (*Template, error) {
	panic("not implemented")
}

func (r *templateTimestampFieldResolver) Timestamp(ctx context.Context, obj *TemplateTimestamp) (*Timestamp, error) {
	panic("not implemented")
}

// Third Party Episode Fields

func (r *thirdPartyEpisodeFieldResolver) Timestamps(ctx context.Context, obj *ThirdPartyEpisode) ([]*ThirdPartyTimestamp, error) {
	panic("not implemented")
}

func (r *thirdPartyEpisodeFieldResolver) Show(ctx context.Context, obj *ThirdPartyEpisode) (*ThirdPartyShow, error) {
	panic("not implemented")
}

// Third Party Timestamp Fields

func (r *thirdPartyTimestampFieldResolver) Type(ctx context.Context, obj *ThirdPartyTimestamp) (*TimestampType, error) {
	panic("not implemented")
}

// Timestamp Fields

func (r *timestampFieldResolver) CreatedBy(ctx context.Context, obj *Timestamp) (*User, error) {
	panic("not implemented")
}

func (r *timestampFieldResolver) UpdatedBy(ctx context.Context, obj *Timestamp) (*User, error) {
	panic("not implemented")
}

func (r *timestampFieldResolver) DeletedBy(ctx context.Context, obj *Timestamp) (*User, error) {
	panic("not implemented")
}

func (r *timestampFieldResolver) Type(ctx context.Context, obj *Timestamp) (*TimestampType, error) {
	panic("not implemented")
}

func (r *timestampFieldResolver) Episode(ctx context.Context, obj *Timestamp) (*Episode, error) {
	panic("not implemented")
}

// Timestamp Type Fields

func (r *timestampTypeFieldResolver) CreatedBy(ctx context.Context, obj *TimestampType) (*User, error) {
	panic("not implemented")
}

func (r *timestampTypeFieldResolver) UpdatedBy(ctx context.Context, obj *TimestampType) (*User, error) {
	panic("not implemented")
}

func (r *timestampTypeFieldResolver) DeletedBy(ctx context.Context, obj *TimestampType) (*User, error) {
	panic("not implemented")
}

// User Fields

func (r *userFieldResolver) AdminOfShows(ctx context.Context, obj *User) ([]*ShowAdmin, error) {
	panic("not implemented")
}
