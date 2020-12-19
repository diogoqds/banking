package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewInternalError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}
