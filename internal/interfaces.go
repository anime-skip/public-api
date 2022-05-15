package internal

import (
	"context"

	"github.com/gofrs/uuid"
)

type Server interface {
	Start() error
}

type RecaptchaService interface {
	Verify(ctx context.Context, response string) error
}

type APIClientService interface {
	Get(ctx context.Context, filter APIClientsFilter) (APIClient, error)
	List(ctx context.Context, filter APIClientsFilter) ([]APIClient, error)
	Create(ctx context.Context, newAPIClient APIClient, createdBy uuid.UUID) (APIClient, error)
	Update(ctx context.Context, newAPIClient APIClient, updatedBy uuid.UUID) (APIClient, error)
	Delete(ctx context.Context, id string, deletedBy uuid.UUID) (APIClient, error)
}

type AuthClaims struct {
	IsAdmin bool
	IsDev   bool
	UserID  uuid.UUID
}
type AuthService interface {
	ValidateAccessToken(token string) (AuthClaims, error)
	ValidateRefreshToken(token string) (AuthClaims, error)
	ValidateVerifyEmailToken(token string) (AuthClaims, error)
	ValidateResetPasswordToken(token string) (AuthClaims, error)
	ValidatePassword(inputPasswordHash, knownPasswordHash string) error
	CreateAccessToken(user FullUser) (string, error)
	CreateRefreshToken(user FullUser) (string, error)
	CreateVerifyEmailToken(user FullUser) (string, error)
	CreateResetPasswordToken(user FullUser) (string, error)
	CreateEncryptedPassword(password string) (string, error)
}

type EmailService interface {
	SendWelcome(ctx context.Context, user FullUser) error
	SendVerification(ctx context.Context, user FullUser, token string) error
	SendResetPassword(ctx context.Context, user FullUser, token string) error
}

type EpisodeService interface {
	ListRecentlyAdded(ctx context.Context, filter RecentlyAddedEpisodesFilter) ([]Episode, error)
	Get(ctx context.Context, filter EpisodesFilter) (Episode, error)
	List(ctx context.Context, filter EpisodesFilter) ([]Episode, error)
	Create(ctx context.Context, newEpisode Episode, createdBy uuid.UUID) (Episode, error)
	Update(ctx context.Context, newEpisode Episode, updatedBy uuid.UUID) (Episode, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (Episode, error)
}

type EpisodeURLService interface {
	Get(ctx context.Context, filter EpisodeURLsFilter) (EpisodeURL, error)
	List(ctx context.Context, filter EpisodeURLsFilter) ([]EpisodeURL, error)
	Create(ctx context.Context, newEpisodeURL EpisodeURL, createdBy uuid.UUID) (EpisodeURL, error)
	Update(ctx context.Context, newEpisodeURL EpisodeURL, updatedBy uuid.UUID) (EpisodeURL, error)
	Delete(ctx context.Context, url string) (EpisodeURL, error)
}

type PreferencesService interface {
	Get(ctx context.Context, filter PreferencesFilter) (Preferences, error)
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	Get(ctx context.Context, filter ShowAdminsFilter) (ShowAdmin, error)
	List(ctx context.Context, filter ShowAdminsFilter) ([]ShowAdmin, error)
	Create(ctx context.Context, newShowAdmin ShowAdmin, createdBy uuid.UUID) (ShowAdmin, error)
	Update(ctx context.Context, newShowAdmin ShowAdmin, updatedBy uuid.UUID) (ShowAdmin, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (ShowAdmin, error)
}

type ShowService interface {
	Get(ctx context.Context, filter ShowsFilter) (Show, error)
	List(ctx context.Context, filter ShowsFilter) ([]Show, error)
	GetSeasonCount(ctx context.Context, id uuid.UUID) (int, error)
	Create(ctx context.Context, newShow Show, createdBy uuid.UUID) (Show, error)
	Update(ctx context.Context, newShow Show, updatedBy uuid.UUID) (Show, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (Show, error)
}

type TemplateService interface {
	Get(ctx context.Context, filter TemplatesFilter) (Template, error)
	List(ctx context.Context, filter TemplatesFilter) ([]Template, error)
	Create(ctx context.Context, newTemplate Template, createdBy uuid.UUID) (Template, error)
	Update(ctx context.Context, newTemplate Template, updatedBy uuid.UUID) (Template, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (Template, error)
}

type TemplateTimestampService interface {
	Get(ctx context.Context, filter TemplateTimestampsFilter) (TemplateTimestamp, error)
	List(ctx context.Context, filter TemplateTimestampsFilter) ([]TemplateTimestamp, error)
	Create(ctx context.Context, newTemplateTimestamp TemplateTimestamp) (TemplateTimestamp, error)
	Delete(ctx context.Context, templateTimestamp InputTemplateTimestamp) (TemplateTimestamp, error)
}

type TimestampService interface {
	Get(ctx context.Context, filter TimestampsFilter) (Timestamp, error)
	List(ctx context.Context, filter TimestampsFilter) ([]Timestamp, error)
	Create(ctx context.Context, newTimestamp Timestamp, createdBy uuid.UUID) (Timestamp, error)
	Update(ctx context.Context, newTimestamp Timestamp, updatedBy uuid.UUID) (Timestamp, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (Timestamp, error)
}

type TimestampTypeService interface {
	Get(ctx context.Context, filter TimestampTypesFilter) (TimestampType, error)
	List(ctx context.Context, filter TimestampTypesFilter) ([]TimestampType, error)
	Create(ctx context.Context, newTimestampType TimestampType, createdBy uuid.UUID) (TimestampType, error)
	Update(ctx context.Context, newTimestampType TimestampType, updatedBy uuid.UUID) (TimestampType, error)
	Delete(ctx context.Context, id uuid.UUID, deletedBy uuid.UUID) (TimestampType, error)
}

type UserService interface {
	Get(ctx context.Context, filter UsersFilter) (FullUser, error)
	List(ctx context.Context, filter UsersFilter) ([]FullUser, error)
	CreateAccount(ctx context.Context, newUser FullUser) (FullUser, error)
	Update(ctx context.Context, newUser FullUser) (FullUser, error)
}

type ThirdPartyService interface {
	Name() string
	FindEpisodeByName(ctx context.Context, name string) ([]ThirdPartyEpisode, error)
}

type Services struct {
	APIClientService         APIClientService
	AuthService              AuthService
	EmailService             EmailService
	EpisodeService           EpisodeService
	EpisodeURLService        EpisodeURLService
	PreferencesService       PreferencesService
	RecaptchaService         RecaptchaService
	ShowAdminService         ShowAdminService
	ShowService              ShowService
	TemplateService          TemplateService
	TemplateTimestampService TemplateTimestampService
	TimestampService         TimestampService
	TimestampTypeService     TimestampTypeService
	UserService              UserService
	ThirdPartyService        ThirdPartyService
}
