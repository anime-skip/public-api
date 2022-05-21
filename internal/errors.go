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

func SQLFailure(op string, err error) error {
	return &Error{
		Code:    EINTERNAL,
		Message: "Unhandled SQL error",
		Op:      op,
		Err:     err,
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
	} else if e1, ok := err.(*Error); ok && e1.Message != "" {
		return e1.Message
	} else if ok && e1.Err != nil {
		return ErrorMessage(e1.Err)
	} else if e2, ok := err.(error); ok {
		return e2.Error()
	}
	return "An internal error has occurred. Contact support@anime-skip.com if the error persists"
}

func ErrorCode(err any) string {
	if err == nil {
		return ""
	} else if e1, ok := err.(*Error); ok {
		return e1.Code
	}
	return ""
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
