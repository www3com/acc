package e

type AppError struct {
	Code int
	Err  string
}

func (e *AppError) Error() string {
	return e.Err
}
