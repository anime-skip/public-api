package internal

import (
	"bytes"
	"fmt"
)

type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

const (
	ECONFLICT = "conflict"  // action cannot be performed
	EINTERNAL = "internal"  // internal error
	EINVALID  = "invalid"   // validation failed
	ENOTFOUND = "not_found" // entity does not exist
)

func NewNotImplemented(op string) error {
	return &Error{
		Message: op + " not implemented",
		Code:    EINTERNAL,
		Op:      op,
	}
}

func NewNotFound(recordName string, op string) error {
	return &Error{
		Message: recordName + " not found",
		Code:    ENOTFOUND,
		Op:      op,
	}
}

func IsNotFound(err error) bool {
	if err == nil {
		return false
	} else if e, ok := err.(*Error); ok {
		return e.Code == ENOTFOUND
	} else if ok && e.Err != nil {
		return IsNotFound(e.Err)
	}
	return false
}

func ErrorMessage(err any) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred. Contact support@anime-skip.com if the error persists"
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
