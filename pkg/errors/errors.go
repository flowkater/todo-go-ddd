package errors

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("resource not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnknownCommand = errors.New("unknown command")
	ErrUnknownQuery   = errors.New("unknown query")
)

type Error struct {
	Code    string
	Message string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func Wrap(err error, code, message string) error {
	return &Error{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
