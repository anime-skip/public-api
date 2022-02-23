package utils

import (
	"fmt"
	"math/rand"

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
		panic(fmt.Errorf("Failed to create id: %v", err))
	}
	return id
}
