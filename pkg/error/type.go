package error

const (
	ErrBaseMsg             = "Business error"
	ErrUnauthorizedMsg     = "Unauthorized user"
	ErrBadRequestMsg       = "Bad request"
	ErrResourceNotFoundMsg = "Resource is not found"
	ErrForbiddenMsg        = "Forbidden action"

	ErrBaseCode = 4000 + iota
	ErrBadRequestCode
	ErrInternalErrorCode
	ErrUnauthorizedCode
	ErrResourceNotFoundCode
	ErrForbiddenCode
)

type AppError interface {
	GetCode() int
	GetMessage() string
	Error() string
}
