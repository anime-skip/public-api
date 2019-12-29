package utils

import (
	"fmt"
	"regexp"

	"github.com/aklinker1/anime-skip-backend/internal/database/entities"
	log "github.com/aklinker1/anime-skip-backend/internal/utils/log"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(EnvString("JWT_SECRET"))

// ValidateAuthHeader parses the authorization header and decides whether or not the token is valid.
func ValidateAuthHeader(authHeader string) (jwt.MapClaims, error) {
	re := regexp.MustCompile(`Bearer (.*?\..*?\..*)`)
	matches := re.FindStringSubmatch(authHeader)
	if len(matches) == 0 {
		return nil, fmt.Errorf("authorization header must be formated as 'Bearer <token>'")
	}

	tokenString := matches[1]
	return ValidateToken(tokenString)
}

// ValidateToken returns the payload for the token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
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
		log.V("%v", err)
		return nil, fmt.Errorf("Invalid claims")
	}
	return payload, nil
}

// GenerateToken creates a jwt token for the specified user
func GenerateToken(user *entities.User) (string, error) {
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
			"aud":    "anime-skip.com",
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

// ValidatePassword checks the password is valid against the bcyrpted hash in the database
func ValidatePassword(password string, bcryptHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(bcryptHash), []byte(password))
	if err != nil {
		return fmt.Errorf("Incorrect password")
	}
	return nil
}

// GenerateEncryptedPassword takes a md5 hashed password and
func GenerateEncryptedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("Failed to encrypt password")
	}
	return string(bytes), nil
}
