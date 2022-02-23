package errors

import (
	"fmt"
	"strings"
)

const RECORD_NOT_FOUND_MESSAGE = "record not found"

func NewRecordNotFound(where string) error {
	return fmt.Errorf("%s (%s)", RECORD_NOT_FOUND_MESSAGE, where)
}

func IsRecordNotFound(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), RECORD_NOT_FOUND_MESSAGE)
}
