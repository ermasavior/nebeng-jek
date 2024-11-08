package error

import (
	"net/http"
)

func ToHttpError(err AppError) int {
	switch err.GetCode() {
	case ErrBaseCode, ErrInternalErrorCode:
		return http.StatusInternalServerError
	case ErrBadRequestCode:
		return http.StatusBadRequest
	case ErrUnauthorizedCode:
		return http.StatusUnauthorized
	case ErrResourceNotFoundCode:
		return http.StatusNotFound
	case ErrForbiddenCode:
		return http.StatusForbidden
	case ErrResourceUnprocessableCode:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
