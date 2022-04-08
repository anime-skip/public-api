package utils

import (
	"math/rand"

	"anime-skip.com/public-api/internal/errors"
	"github.com/gofrs/uuid"
)

// RandomString of a specific length
func RandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		panic(errors.NewPanicedError("Failed to create id: %v", err))
	}
	return id
}
