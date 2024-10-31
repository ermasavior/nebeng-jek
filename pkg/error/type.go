package error

const (
	ErrBaseMsg             = "business error"
	ErrUnauthorizedMsg     = "unauthorized user"
	ErrBadRequestMsg       = "bad request"
	ErrResourceNotFoundMsg = "resource is not found"
	ErrForbiddenMsg        = "forbidden action"

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
