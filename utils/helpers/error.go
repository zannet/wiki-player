package helpers

type AppError struct {
	Error   error
	Code    int
	Message string
}
