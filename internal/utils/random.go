package utils

import (
	"math/rand"

	"anime-skip.com/public-api/internal"
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

func RandomID(target *uuid.UUID) error {
	id, err := uuid.NewV4()
	if err != nil {
		return &internal.Error{
			Code:    internal.EINTERNAL,
			Message: "Failed to generate uuid",
			Op:      "utils.RandomID",
			Err:     err,
		}
	}
	*target = id
	return nil
}
