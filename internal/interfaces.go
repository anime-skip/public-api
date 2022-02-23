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
}

type AuthClaims struct {
	IsAdmin bool
	IsDev   bool
	UserID  uuid.UUID
}
type AuthService interface {
	// ValidateAccessToken returns AuthenticationDetails or errors out if the token is invalid
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
}

type GetRecentlyAddedParams struct {
	Pagination
}
type EpisodeService interface {
	GetRecentlyAdded(ctx context.Context, params GetRecentlyAddedParams) ([]Episode, error)
	GetByID(ctx context.Context, id uuid.UUID) (Episode, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]Episode, error)
}

type EpisodeURLService interface {
	GetByURL(ctx context.Context, url string) (EpisodeURL, error)
	GetByEpisodeId(ctx context.Context, episodeID uuid.UUID) ([]EpisodeURL, error)
}

type PreferencesService interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (Preferences, error)
	CreateInTx(ctx context.Context, tx Tx, newPreferences Preferences) (Preferences, error)
	Update(ctx context.Context, newPreferences Preferences) (Preferences, error)
}

type ShowAdminService interface {
	GetByID(ctx context.Context, id uuid.UUID) (ShowAdmin, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]ShowAdmin, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]ShowAdmin, error)
}

type ShowService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Show, error)
}

type TemplateService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Template, error)
	GetByShowID(ctx context.Context, showID uuid.UUID) ([]Template, error)
	GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) (Template, error)
}

type TemplateTimestampService interface {
}

type ThirdPartyEpisodeService interface {
}

type ThirdPartyTimestampService interface {
}

type TimestampService interface {
	GetByID(ctx context.Context, id uuid.UUID) (Timestamp, error)
	GetByEpisodeID(ctx context.Context, episodeID uuid.UUID) ([]Timestamp, error)
}

type TimestampTypeService interface {
	GetByID(ctx context.Context, id uuid.UUID) (TimestampType, error)
	GetAll(ctx context.Context) ([]TimestampType, error)
}

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (User, error)
	CreateInTx(ctx context.Context, tx Tx, newUser User) (User, error)
	// DeleteInTx(ctx context.Context, tx Tx, user User) (EpisodeURL, error)
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
