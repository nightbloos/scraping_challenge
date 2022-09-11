package domain

import (
	"errors"
	"fmt"
)

const (
	// ErrBadRequest used to mark error as bad request
	ErrBadRequest string = "bad_request"
	// ErrNotFound used to mark error as not found
	ErrNotFound string = "not_found"
	// ErrInternal used to mark error as internal
	ErrInternal string = "internal"
	// ErrForbidden used to mark error as forbidden
	ErrForbidden string = "forbidden"
	// ErrUnauthorized used to mark error as unauthorized
	ErrUnauthorized string = "unauthorized"
)

// NewBadRequestError creates new Bad Request error
func NewBadRequestError(msg string) Error {
	return WrapWithBadRequestError(nil, msg)
}

// WrapWithBadRequestError wraps existing error with Bad Request error
func WrapWithBadRequestError(err error, msg string) Error {
	return Error{errorType: ErrBadRequest, err: err, msg: msg}
}

// NewNotFoundError creates new Not Found error
func NewNotFoundError(msg string) Error {
	return WrapWithNotFoundError(nil, msg)
}

// WrapWithNotFoundError wraps existing error with Not Found error
func WrapWithNotFoundError(err error, msg string) Error {
	return Error{errorType: ErrNotFound, err: err, msg: msg}
}

// NewInternalError creates new Internal error
func NewInternalError(msg string) Error {
	return WrapWithInternalError(nil, msg)
}

// WrapWithInternalError wraps existing error with Internal error
func WrapWithInternalError(err error, msg string) Error {
	return Error{errorType: ErrInternal, err: err, msg: msg}
}

// WrapWithForbiddenError wraps existing error with ErrForbidden error
func WrapWithForbiddenError(err error, msg string) Error {
	return Error{errorType: ErrForbidden, err: err, msg: msg}
}

// NewForbiddenError creates new Forbidden error
func NewForbiddenError(msg string) Error {
	return WrapWithForbiddenError(nil, msg)
}

// NewUnauthorizedError creates new Unauthorized error
func NewUnauthorizedError(msg string) Error {
	return WrapWithUnauthorizedError(nil, msg)
}

// WrapWithUnauthorizedError wraps existing error with Unauthorized error
func WrapWithUnauthorizedError(err error, msg string) Error {
	return Error{errorType: ErrUnauthorized, err: err, msg: msg}
}

// Error as common errors for our services
type Error struct {
	msg       string
	err       error
	errorType string
}

// GetErrorType returns error type
func (e Error) GetErrorType() string {
	return e.errorType
}

func (e Error) Error() string {
	if e.err == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.err.Error()
	}
	return fmt.Sprintf("%s: %s", e.msg, e.err.Error())
}

// Cause returns errors cause
func (e Error) Cause() error {
	return e.err
}

// Unwrap unwraps error
func (e Error) Unwrap() error {
	return e.err
}

// IsBadRequestError checks if Bad Request type error provided
func IsBadRequestError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrBadRequest {
		return true
	}
	return false
}

// IsForbiddenError checks if Forbidden type error provided
func IsForbiddenError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrForbidden {
		return true
	}
	return false
}

// IsNotFoundError checks if NotFound type error provided
func IsNotFoundError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrNotFound {
		return true
	}
	return false
}

// IsInternalError checks if Internal type error provided
func IsInternalError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrInternal {
		return true
	}
	return false
}

// IsUnauthorizedError checks if Unauthorized type error provided
func IsUnauthorizedError(e error) bool {
	var err Error
	if errors.As(e, &err) && err.GetErrorType() == ErrUnauthorized {
		return true
	}
	return false
}
