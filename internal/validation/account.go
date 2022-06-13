package validation

import (
	"strings"

	"anime-skip.com/public-api/internal"
	"github.com/asaskevich/govalidator"
)

func AccountUsername(username string) error {
	if len(strings.TrimSpace(username)) < 3 {
		return &internal.Error{
			Code:    internal.EINVALID,
			Message: "Username must be at least 3 characters long",
		}
	}
	return nil
}

func AccountEmail(email string) error {
	if govalidator.IsEmail(email) {
		return nil
	}
	return &internal.Error{
		Code:    internal.EINVALID,
		Message: "Email is not valid",
	}
}
