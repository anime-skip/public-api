package auth

import (
	"fmt"
	"regexp"
	"time"

	"anime-skip.com/backend/internal/database/entities"
	log "anime-skip.com/backend/internal/utils/log"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"anime-skip.com/backend/internal/utils/env"
)

var jwtSecret = []byte(env.JWT_SECRET)

var day time.Duration = 24 * time.Hour
var week time.Duration = 7 * day

// TODO: simplify audiences, switch to including permissions in the token instead
// We need to make sure refreshing is working in web and web-ext first

const (
	AUD_AUTH_TOKEN               = "anime-skip.com"                      // TODO: Change to api.anime-skip.com
	AUD_REFRESH_TOKEN            = "anime-skip.com/graphql?loginRefresh" // TODO: Change to api.anime-skip.com
	AUD_EMAIL_VERIFICATION_TOKEN = "anime-skip.com/verify-email-address"
	AUD_RESET_PASSWORD_TOKEN     = "anime-skip.com/forgot-password"
)

const ISSUER = "anime-skip.com" // TODO: Change to api.anime-skip.com

// ValidateAuthHeader parses the authorization header and decides whether or not the token is valid.
func ValidateAuthHeader(authHeader string) (jwt.MapClaims, error) {
	re := regexp.MustCompile(`Bearer (.*?\..*?\..*)`)
	matches := re.FindStringSubmatch(authHeader)
	if len(matches) == 0 {
		return nil, fmt.Errorf("authorization header must be formated as 'Bearer <token>'")
	}

	tokenString := matches[1]
	return validateToken(tokenString)
}

// validateToken returns the payload for the token
func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate Algorithm
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Internal error 401-01")
		}
		// return secret
		return jwtSecret, nil
	})
	if err != nil {
		log.V("%v", err)
		return nil, fmt.Errorf("Invalid Token")
	}
	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid claims")
	}
	if isValidExpiresAt := payload.VerifyIssuedAt(time.Now().Unix(), true); !isValidExpiresAt {
		return nil, fmt.Errorf("Token is expired")
	}
	if isValidIssuer := payload.VerifyIssuer("anime-skip.com", true); !isValidIssuer {
		return nil, fmt.Errorf("Invalid issuer, expected 'anime-skip.com', but got '%v'", payload["iss"])
	}
	return payload, nil
}

func validateTokenWithAud(token, aud string) (jwt.MapClaims, error) {
	claims, err := validateToken(token)
	if err != nil {
		return nil, err
	}
	if isValidAud := claims.VerifyAudience(aud, true); !isValidAud {
		return nil, fmt.Errorf("Invalid aud, expected '%s'", aud)
	}
	return claims, nil
}

func ValidateAuthToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, AUD_AUTH_TOKEN)
}

func ValidateRefreshToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, AUD_REFRESH_TOKEN)
}

func ValidateEmailVerificationToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, AUD_EMAIL_VERIFICATION_TOKEN)
}

func ValidateResetPasswordToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, AUD_RESET_PASSWORD_TOKEN)
}

func generateGeneralToken(
	label string,
	duration time.Duration,
	customClaims jwt.MapClaims,
) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"exp": now.Add(duration).Unix(),
		"iat": now.Unix(),
		"iss": ISSUER,
	}
	for key, value := range customClaims {
		claims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.E("%v", err)
		return "", fmt.Errorf("Internal error: failed to generate %s token", label)
	}
	return tokenString, nil
}

// GenerateAuthToken creates a jwt token for the specified user
func GenerateAuthToken(user *entities.User) (string, error) {
	return generateGeneralToken("auth", 12*time.Hour, jwt.MapClaims{
		"aud":    AUD_AUTH_TOKEN,
		"userId": user.ID,
		"role":   user.Role,
	})
}

// GenerateRefreshToken creates a refresh token to be used with the login query
func GenerateRefreshToken(user *entities.User) (string, error) {
	return generateGeneralToken("refresh", 1*week, jwt.MapClaims{
		"aud":    AUD_REFRESH_TOKEN,
		"userId": user.ID,
	})
}

func GenerateVerifyEmailToken(user *entities.User) (string, error) {
	return generateGeneralToken("verify_email_address", 2*day, jwt.MapClaims{
		"aud":    AUD_EMAIL_VERIFICATION_TOKEN,
		"userId": user.ID,
	})
}

func GenerateResetPasswordToken(user *entities.User) (string, error) {
	return generateGeneralToken("reset_password", 10*time.Minute, jwt.MapClaims{
		"aud":    AUD_RESET_PASSWORD_TOKEN,
		"userId": user.ID,
	})
}

// ValidatePassword checks the password is valid against the bcyrpted hash in the database
func ValidatePassword(password string, bcryptHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptHash), []byte(password))
	if err != nil {
		return fmt.Errorf("Incorrect password")
	}
	return nil
}

// GenerateEncryptedPassword takes a md5 hashed password and
func GenerateEncryptedPassword(passwordHash string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordHash), 14)
	if err != nil {
		return "", fmt.Errorf("Failed to encrypt password")
	}
	return string(bytes), nil
}
