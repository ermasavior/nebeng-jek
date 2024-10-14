package error

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Raw     error
}

func (e AppError) Error() string {
	return e.Message
}

func NewInternalServerError(err error, msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Raw:     err,
		Message: msg,
	}
}

func NewNotFound(err error, msg string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Raw:     err,
		Message: msg,
	}
}

func NewForbidden(err error, msg string) *AppError {
	return &AppError{
		Code:    http.StatusForbidden,
		Raw:     err,
		Message: msg,
	}
}

func NewBadRequest(err error, msg string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		Raw:     err,
		Message: msg,
	}
}

func NewUnauthorized(err error, msg string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		Raw:     err,
		Message: msg,
	}
}
