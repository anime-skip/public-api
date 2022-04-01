package jwt

import (
	"fmt"
	"strings"
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
	secret []byte
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

func NewJWTAuthService(secret string) internal.AuthService {
	log.D("Using Custom JWT Authentication...")
	return &jwtAuthService{
		secret: []byte(secret),
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
		return "", fmt.Errorf("Internal error: failed to generate %s token", audience)
	}
	return tokenString, nil
}

func (s *jwtAuthService) validateToken(name string, token string, audience string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Validate Algorithm
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Internal error 401-01")
		}
		// return secret
		return s.secret, nil
	})
	if err != nil {
		log.V("%v", err)
		return nil, fmt.Errorf("Invalid %s token", strings.ToLower(name))
	}
	payload, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("%s token has invalid claims", name)
	}
	if isValidExpiresAt := payload.VerifyIssuedAt(time.Now().Unix(), true); !isValidExpiresAt {
		return nil, fmt.Errorf("%s token is expired", name)
	}
	if isValidAudience := payload.VerifyAudience(audience, true); !isValidAudience {
		return nil, fmt.Errorf("%s token has invalid audience", name)
	}
	if isValidIssuer := payload.VerifyIssuer(ISSUER, true); !isValidIssuer {
		return nil, fmt.Errorf("%s token has invalid issuer, expected '%s', but got '%v'", name, ISSUER, payload["iss"])
	}
	return payload, nil
}

func (s *jwtAuthService) mapAuthClaims(claims jwt.MapClaims) (internal.AuthClaims, error) {
	role := claims["role"].(float64)
	userID, err := uuid.FromString(claims["userId"].(string))
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return internal.AuthClaims{
		IsAdmin: role == internal.ROLE_ADMIN,
		IsDev:   role == internal.ROLE_DEV,
		UserID:  userID,
	}, nil
}

func (s *jwtAuthService) ValidateAccessToken(token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Access", token, AUD_ACCESS_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(claims)
}

func (s *jwtAuthService) ValidateRefreshToken(token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Refresh", token, AUD_REFRESH_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(claims)
}

func (s *jwtAuthService) ValidateVerifyEmailToken(token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Email", token, AUD_VERIFY_EMAIL_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(claims)
}

func (s *jwtAuthService) ValidateResetPasswordToken(token string) (internal.AuthClaims, error) {
	claims, err := s.validateToken("Password reset", token, AUD_RESET_PASSWORD_TOKEN)
	if err != nil {
		return internal.AuthClaims{}, err
	}
	return s.mapAuthClaims(claims)
}

func (s *jwtAuthService) ValidatePassword(inputPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	if err != nil {
		return fmt.Errorf("Incorrect password")
	}
	return nil
}

func (s *jwtAuthService) CreateAccessToken(user internal.User) (string, error) {
	return s.createToken(AUD_ACCESS_TOKEN, TIMEOUT_ACCESS_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
	})
}

func (s *jwtAuthService) CreateRefreshToken(user internal.User) (string, error) {
	return s.createToken(AUD_REFRESH_TOKEN, TIMEOUT_REFRESH_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
	})
}

func (s *jwtAuthService) CreateVerifyEmailToken(user internal.User) (string, error) {
	return s.createToken(AUD_VERIFY_EMAIL_TOKEN, TIMEOUT_VERIFY_EMAIL_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
	})
}

func (s *jwtAuthService) CreateResetPasswordToken(user internal.User) (string, error) {
	return s.createToken(AUD_RESET_PASSWORD_TOKEN, TIMEOUT_RESET_PASSWORD_TOKEN, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
	})
}

func (s *jwtAuthService) CreateEncryptedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("Failed to encrypt password")
	}
	return string(bytes), nil
}
