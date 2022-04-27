package err

type AppError interface {
	Code() int
	Message() string
	Error() string
}
