package error

type appError struct {
	code    int
	message string
}

func NewInternalServerError(msg string) AppError {
	return &appError{
		code:    ErrInternalErrorCode,
		message: msg,
	}
}

func NewNotFoundError(msg string) AppError {
	return &appError{
		code:    ErrResourceNotFoundCode,
		message: msg,
	}
}

func NewForbiddenError(msg string) AppError {
	return &appError{
		code:    ErrForbiddenCode,
		message: msg,
	}
}

func NewBadRequestError(msg string) AppError {
	return &appError{
		code:    ErrBadRequestCode,
		message: msg,
	}
}

func NewUnauthorizedError(msg string) AppError {
	return &appError{
		code:    ErrUnauthorizedCode,
		message: msg,
	}
}

func NewUnprocessableError(msg string) AppError {
	return &appError{
		code:    ErrResourceUnprocessableCode,
		message: msg,
	}
}

func (e appError) GetCode() int {
	return e.code
}

func (e appError) GetMessage() string {
	return e.message
}

func (e appError) Error() string {
	return e.message
}

func (e appError) String() string {
	return e.message
}
