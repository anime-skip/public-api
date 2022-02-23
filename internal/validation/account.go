package validation

import (
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
)

func AccountUsername(username string) error {
	if len(strings.TrimSpace(username)) < 3 {
		return errors.New("Username must be at least 3 characters long")
	}
	return nil
}

func AccountEmail(email string) error {
	if govalidator.IsEmail(email) {
		return nil
	}
	return errors.New("Email is not valid")
}
