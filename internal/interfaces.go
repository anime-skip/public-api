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
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]APIClient, error)
	Create(ctx context.Context, newAPIClient APIClient) (APIClient, error)
	Update(ctx context.Context, newAPIClient APIClient) (APIClient, error)
	Delete(ctx context.Context, apiClient APIClient) (APIClient, error)
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
	CreateAccessToken(user User) (string, error)
	CreateRefreshToken(user User) (string, error)
	CreateVerifyEmailToken(user User) (string, error)
	CreateResetPasswordToken(user User) (string, error)
	CreateEncryptedPassword(password string) (string, error)
}

type EmailService interface {
	SendWelcome(user User) error
	SendVerification(user User, token string) error
	SendResetPassword(user User, token string) error
}

type GetRecentlyAddedParams struct {
	Pagination
}
type EpisodeService interface {
	GetRecentlyAdded(ctx context.Context, params GetRecentlyAddedParams) ([]Episode, error)
	GetByID(ctx context.Context, id uuid.UUID) (Episode, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]Episode, error)
	Create(ctx context.Context, newEpisode Episode) (Episode, error)
	Update(ctx context.Context, newEpisode Episode) (Episode, error)
	Delete(ctx context.Context, episode Episode) (Episode, error)
}

type EpisodeURLService interface {
	GetByURL(ctx context.Context, url string) (EpisodeURL, error)
	GetByEpisodeId(ctx context.Context, episodeID uuid.UUID) ([]EpisodeURL, error)
	Create(ctx context.Context, newEpisodeURL EpisodeURL) (EpisodeURL, error)
	Update(ctx context.Context, newEpisodeURL EpisodeURL) (EpisodeURL, error)
	Delete(ctx context.Context, episodeURL EpisodeURL) (EpisodeURL, error)
}

type PreferencesService interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (Preferences, error)
	NewDefault(ctx context.Context, userID uuid.UUID) Preferences
	CreateInTx(ctx context.Context, tx Tx, newPreferences Preferences) (Preferences, error)
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	GetByID(ctx context.Context, id uuid.UUID) (ShowAdmin, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]ShowAdmin, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]ShowAdmin, error)
	Create(ctx context.Context, newShowAdmin ShowAdmin) (ShowAdmin, error)
	Update(ctx context.Context, newShowAdmin ShowAdmin) (ShowAdmin, error)
	Delete(ctx context.Context, showAdmin ShowAdmin) (ShowAdmin, error)
}

type ShowService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Show, error)
	GetSeasonCount(ctx context.Context, id uuid.UUID) (int, error)
	Create(ctx context.Context, newShow Show) (Show, error)
	Update(ctx context.Context, newShow Show) (Show, error)
	Delete(ctx context.Context, show Show) (Show, error)
}

type TemplateService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Template, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]Template, error)
	GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) (Template, error)
	Create(ctx context.Context, newTemplate Template) (Template, error)
	Update(ctx context.Context, newTemplate Template) (Template, error)
	Delete(ctx context.Context, template Template) (Template, error)
}

type TemplateTimestampService interface {
	GetByTimestampID(ctx context.Context, timestampID uuid.UUID) (TemplateTimestamp, error)
	GetByTemplateID(ctx context.Context, templateID uuid.UUID) ([]TemplateTimestamp, error)
	Create(ctx context.Context, newTemplateTimestamp TemplateTimestamp) (TemplateTimestamp, error)
	Delete(ctx context.Context, templateTimestamp TemplateTimestamp) (TemplateTimestamp, error)
}

type TimestampService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Timestamp, error)
	GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) ([]Timestamp, error)
	Create(ctx context.Context, newTimestamp Timestamp) (Timestamp, error)
	Update(ctx context.Context, newTimestamp Timestamp) (Timestamp, error)
	Delete(ctx context.Context, timestamp Timestamp) (Timestamp, error)
}

type TimestampTypeService interface {
	GetByID(ctx context.Context, id uuid.UUID) (TimestampType, error)
	GetAll(ctx context.Context) ([]TimestampType, error)
	Create(ctx context.Context, newTimestampType TimestampType) (TimestampType, error)
	Update(ctx context.Context, newTimestampType TimestampType) (TimestampType, error)
	Delete(ctx context.Context, timestampType TimestampType) (TimestampType, error)
}

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (User, error)
	CreateInTx(ctx context.Context, tx Tx, newUser User) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, user User) (User, error)
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
}

type DirectiveServices struct {
	AuthService      AuthService
	ShowAdminService ShowAdminService
}
