package appError

import "net/http"

type AppError struct {
	Message    string      `json:"message"`
	AppCode    string      `json:"errorCode"`
	Data       interface{} `json:"data,omitempty"`
	StatusCode int         `json:"-"`
}

func NewAppError() *AppError {
	return &AppError{}
}

func NewServerError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusInternalServerError).
		WithAppCode(ErrorCodeInternalServer).
		WithMessage(ErrorMessageInternalServer)
}

func NewBadRequestError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusBadRequest).
		WithAppCode(ErrorCodeBadRequest).
		WithMessage(ErrorMessageBadRequest)
}

func NewNotFoundError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusNotFound).
		WithAppCode(ErrorCodeNotFound).
		WithMessage(ErrorMessageNotFound)
}

func NewUnauthorizedError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusUnauthorized).
		WithAppCode(ErrorCodeUnauthorized).
		WithMessage(ErrorMessageUnauthorized)
}

func NewForbiddenError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusForbidden).
		WithAppCode(ErrorCodeForbidden).
		WithMessage(ErrorMessageForbidden)
}

func NewExpectationFailedError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusExpectationFailed).
		WithAppCode(ErrorCodeExpectationFailed).
		WithMessage(ErrorMessageExpectationFailed)
}

func NewConflictError() *AppError {
	return NewAppError().
		WithStatusCode(http.StatusConflict).
		WithAppCode(ErrorCodeConflict).
		WithMessage(ErrorMessageConflict)
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) WithMessage(message string) *AppError {
	e.Message = message
	return e
}

func (e *AppError) DefaultMessage() *AppError {
	e.Message = ErrorMessageInternalServer
	return e
}

func (e *AppError) WithAppCode(appCode string) *AppError {
	e.AppCode = appCode
	return e
}

func (e *AppError) WithData(data interface{}) *AppError {
	e.Data = data
	return e
}

func (e *AppError) WithError(err error) *AppError {
	if appError, ok := err.(*AppError); ok {
		e.AppCode = appError.AppCode
		e.Message = appError.Message

		if appError.StatusCode != 0 {
			e.StatusCode = appError.StatusCode
		}

		if appError.Data != nil {
			e.Data = appError.Data
		}
	} else {
		e.Message = err.Error()
	}
	return e
}

func (e *AppError) DefaultAppCode() *AppError {
	e.AppCode = ErrorCodeInternalServer
	return e
}

func (e *AppError) WithStatusCode(statusCode int) *AppError {
	e.StatusCode = statusCode
	return e
}

func (e *AppError) DefaultStatusCode() *AppError {
	e.StatusCode = http.StatusInternalServerError
	return e
}

func (e *AppError) AddDefaultValuesIfMissing() *AppError {
	if len(e.Message) == 0 {
		e.DefaultMessage()
	}

	if len(e.AppCode) == 0 {
		e.DefaultAppCode()
	}

	if e.StatusCode == 0 {
		e.DefaultStatusCode()
	}

	return e
}

const (
	ErrorCodeInternalServer    = "internal_server_error"
	ErrorCodeBadRequest        = "bad_request"
	ErrorCodeNotFound          = "not_found"
	ErrorCodeUnauthorized      = "unauthorized"
	ErrorCodeForbidden         = "forbidden"
	ErrorCodeExpectationFailed = "expectation_failed"
	ErrorCodeConflict          = "conflict"
)

const (
	ErrorMessageInternalServer    = "internal server error"
	ErrorMessageBadRequest        = "bad request"
	ErrorMessageNotFound          = "not found"
	ErrorMessageUnauthorized      = "unauthorized"
	ErrorMessageForbidden         = "forbidden"
	ErrorMessageExpectationFailed = "expectation failed"
	ErrorMessageConflict          = "conflict"
)
