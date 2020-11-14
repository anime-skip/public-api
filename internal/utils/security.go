package utils

import (
	"fmt"
	"regexp"

	"anime-skip.com/backend/internal/database/entities"
	log "anime-skip.com/backend/internal/utils/log"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(ENV.JWT_SECRET)

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
	if isValidExpiresAt := payload.VerifyIssuedAt(CurrentTimeSec(), true); !isValidExpiresAt {
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
	return validateTokenWithAud(token, "anime-skip.com")
}

func ValidateRefreshToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, "anime-skip.com/graphql?loginRefresh")
}

func ValidateEmailVerificationToken(token string) (jwt.MapClaims, error) {
	return validateTokenWithAud(token, "anime-skip.com/verify-email-address")
}

// GenerateAuthToken creates a jwt token for the specified user
func GenerateAuthToken(user *entities.User) (string, error) {
	now := CurrentTimeSec()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"aud":    "anime-skip.com",
			"exp":    now + 43200, // 12hrs in seconds = 12*60*60
			"iat":    now,
			"iss":    "anime-skip.com",
			"userId": user.ID,
			"role":   user.Role,
		},
	)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.E("%v", err)
		return "", fmt.Errorf("Internal error: failed to generate auth token")
	}
	return tokenString, nil
}

// GenerateRefreshToken creates a refresh token to be used with the login query
func GenerateRefreshToken(user *entities.User) (string, error) {
	now := CurrentTimeSec()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"aud":    "anime-skip.com/graphql?loginRefresh",
			"exp":    now + 604800, // 7 days in seconds = 7*24*60*60
			"iat":    now,
			"iss":    "anime-skip.com",
			"userId": user.ID,
		},
	)
	refreshTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.E("%v", err)
		return "", fmt.Errorf("Internal error: failed to generate refresh token")
	}
	return refreshTokenString, nil
}

func GenerateVerifyEmailToken(user *entities.User) (string, error) {
	now := CurrentTimeSec()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"aud":    "anime-skip.com/verify-email-address",
			"exp":    now + 172800, // 2 days in seconds = 2*24*60*60
			"iat":    now,
			"iss":    "anime-skip.com",
			"userId": user.ID,
		},
	)
	verifyEmailTokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.E("%v", err)
		return "", fmt.Errorf("Internal error: failed to generate verify email token")
	}
	return verifyEmailTokenString, nil
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
