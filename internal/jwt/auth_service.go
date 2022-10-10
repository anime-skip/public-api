package jwt

import (
	"context"
	"fmt"
	"time"

	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

var day time.Duration = 24 * time.Hour

type jwtAuthService struct {
	secret      []byte
	userService internal.UserService
}

// TODO: simplify audiences, switch to including permissions in the token instead
// We need to make sure refreshing is working in web and web-ext first

const (
	AUD_ACCESS_TOKEN         = "anime-skip.com"                      // TODO: Change to api.anime-skip.com
	AUD_REFRESH_TOKEN        = "anime-skip.com/graphql?loginRefresh" // TODO: Change to api.anime-skip.com
	AUD_VERIFY_EMAIL_TOKEN   = "anime-skip.com/verify-email-address" // TODO: Change to api.anime-skip.com
	AUD_RESET_PASSWORD_TOKEN = "anime-skip.com/forgot-password"      // TODO: Change to api.anime-skip.com
)

var TIMEOUT_ACCESS_TOKEN = 7 * day
var TIMEOUT_REFRESH_TOKEN = 30 * day
var TIMEOUT_VERIFY_EMAIL_TOKEN = 2 * day
var TIMEOUT_RESET_PASSWORD_TOKEN = 10 * time.Minute

// TODO: Change to api.anime-skip.com
const ISSUER = "anime-skip.com"

func NewJWTAuthService(secret string, userService internal.UserService) internal.AuthService {
	log.D("Using Custom JWT Authentication...")
	return &jwtAuthService{
		secret:      []byte(secret),
		userService: userService,
	}
}

func (s *jwtAuthService) createToken(
	audience string,
	duration time.Duration,
	customClaims jwt.MapClaims,
) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"exp": now.Add(duration).Unix(),
		"iat": now.Unix(),
		"iss": ISSUER,
		"aud": audience,
	}
	for key, value := range customClaims {
		claims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.secret)
	if err != nil {
		log.E("%v", err)
		return "", &internal.Error{
			Code:    internal.EINTERNAL,
			Message: fmt.Sprintf("Internal error: failed to generate %s token", audience),
			Op:      "createToken",
			Err:     err,
		}
	}
	return tokenString, nil
}

func (s *jwtAuthService) validateToken(name string, token string, audience string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// Validate Algorithm
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &internal.Error{
				Code:    internal.EINVALID,
				Message: "Internal error 401-01",
				Op:      "validateToken",
			}
		}
		// return secret
		return s.secret, nil
	})
	if err != nil {
		log.V("%v", err)
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: "Invalid Token", // TODO: Update extension to look for something other than the message
			Op:      "validateToken",
		}
	}
	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: fmt.Sprintf("%s token has invalid claims", name),
			Op:      "validateToken",
		}
	}
	if isValidExpiresAt := payload.VerifyIssuedAt(time.Now().Unix(), true); !isValidExpiresAt {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: fmt.Sprintf("%s token is expired", name),
			Op:      "validateToken",
		}
	}
	if isValidAudience := payload.VerifyAudience(audience, true); !isValidAudience {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: fmt.Sprintf("%s token has invalid audience", name),
			Op:      "validateToken",
		}
	}
	if isValidIssuer := payload.VerifyIssuer(ISSUER, true); !isValidIssuer {
		return nil, &internal.Error{
			Code:    internal.EINVALID,
			Message: fmt.Sprintf("%s token has invalid issuer, expected '%s', but got '%v'", name, ISSUER, payload["iss"]),
			Op:      "validateToken",
		}
	}
	return payload, nil
}

func (s *jwtAuthService) mapAuthClaims(ctx context.Context, claims jwt.MapClaims) (internal.AuthClaims, error) {
	userID, err := uuid.FromString(claims["userId"].(string))
	if err != nil {
		return internal.AuthClaims{}, err
	}
	user, err := s.userService.Get(ctx, internal.UsersFilter{
		ID: &userID,
	})
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return internal.AuthClaims{
		Role:   user.Role,
		UserID: user.ID,
	}, nil
}

func (s *jwtAuthService) ValidateAccessToken(ctx context.Context, token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Access", token, AUD_ACCESS_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(ctx, claims)
}

func (s *jwtAuthService) ValidateRefreshToken(ctx context.Context, token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Refresh", token, AUD_REFRESH_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(ctx, claims)
}

func (s *jwtAuthService) ValidateVerifyEmailToken(ctx context.Context, token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Email", token, AUD_VERIFY_EMAIL_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(ctx, claims)
}

func (s *jwtAuthService) ValidateResetPasswordToken(ctx context.Context, token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Password reset", token, AUD_RESET_PASSWORD_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(ctx, claims)
}

func (s *jwtAuthService) ValidatePassword(inputPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		return &internal.Error{
			Code:    internal.EINVALID,
			Message: "Incorrect password",
			Op:      "ValidatePassword",
		}
	}
	return nil
}

func (s *jwtAuthService) CreateAccessToken(user internal.FullUser) (string, error) {
	roleInt, err := user.Role.Value()
	if err != nil {
		return "", err
	}
	return s.createToken(AUD_ACCESS_TOKEN, TIMEOUT_ACCESS_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   roleInt,
	})
}

func (s *jwtAuthService) CreateRefreshToken(user internal.FullUser) (string, error) {
	roleInt, err := user.Role.Value()
	if err != nil {
		return "", err
	}
	return s.createToken(AUD_REFRESH_TOKEN, TIMEOUT_REFRESH_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   roleInt,
	})
}

func (s *jwtAuthService) CreateVerifyEmailToken(user internal.FullUser) (string, error) {
	roleInt, err := user.Role.Value()
	if err != nil {
		return "", err
	}
	return s.createToken(AUD_VERIFY_EMAIL_TOKEN, TIMEOUT_VERIFY_EMAIL_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   roleInt,
	})
}

func (s *jwtAuthService) CreateResetPasswordToken(user internal.FullUser) (string, error) {
	roleInt, err := user.Role.Value()
	if err != nil {
		return "", err
	}
	return s.createToken(AUD_RESET_PASSWORD_TOKEN, TIMEOUT_RESET_PASSWORD_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   roleInt,
	})
}

func (s *jwtAuthService) CreateEncryptedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", &internal.Error{
			Code:    internal.EINVALID,
			Message: "Failed to encrypt password",
			Op:      "CreateEncryptedPassword",
		}
	}
	return string(bytes), nil
}
