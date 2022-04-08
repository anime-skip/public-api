package errors

import "fmt"

type PanicedError struct {
	message string
}

var Paniced = &PanicedError{}

func NewPanicedError(format string, vars ...interface{}) *PanicedError {
	return &PanicedError{
		message: fmt.Sprintf(format, vars...),
	}
}

func (err *PanicedError) Error() string {
	return err.message
}
