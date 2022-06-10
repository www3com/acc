package e

import (
	"fmt"
)

type WrapError struct {
	Code    int
	Message string
	Cause   error
}

func (e *WrapError) Error() string {
	return e.Message + ": " + e.Cause.Error()
}

func New(code int) error {
	return &WrapError{
		Code: code,
	}
}

func WithCode(code int, err error, message string) error {
	return &WrapError{
		Code:    code,
		Message: message,
		Cause:   err,
	}
}

func WithCodef(code int, err error, format string, args ...interface{}) error {
	return &WrapError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
		Cause:   err,
	}
}

func Wrap(err error, message string) error {
	return WithCode(ERROR, err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return WithCodef(ERROR, err, format, args...)
}

func UnWrap(err error) (*WrapError, bool) {
	v, ok := err.(*WrapError)
	return v, ok
}
