package e

import (
	"github.com/pkg/errors"
)

type Error struct {
	Code  int
	Cause error
}

func (e *Error) Error() string {
	return e.Cause.Error()
}

func Wrap(code int, err error, message string) error {
	return &Error{
		Code:  code,
		Cause: errors.Wrap(err, message),
	}
}

func WithStack(code int, err error) error {
	return &Error{
		Code:  code,
		Cause: errors.WithStack(err),
	}
}

func New(code int) *Error {
	return &Error{
		Code: code,
	}
}
