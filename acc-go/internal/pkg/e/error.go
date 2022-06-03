package e

import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[%d]%v", e.Code, e.Message)
}

func New(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
