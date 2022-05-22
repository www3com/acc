package e

import "strconv"

type AppError struct {
	Code int
}

func New(code int) error {
	return &AppError{Code: code}
}

func (e *AppError) Error() string {
	return strconv.Itoa(e.Code)
}
