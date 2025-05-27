package errors

import (
	"net/http"
)

const (
	ValidationError   = "VALIDATION_ERROR"
	NotFoundError     = "NOT_FOUND_ERROR"
	ConflictError     = "CONFLICT_ERROR"
	UnauthorizedError = "UNAUTHORIZED_ERROR"
	InternalError     = "INTERNAL_ERROR"
	AppError          = "APP_ERROR"
)

type APPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"-"`
	Data    any    `json:"data,omitempty"`
}

func NewAPPError(code int, message string, errType string, data any) *APPError {
	return &APPError{Code: code, Message: message, Type: errType, Data: data}
}

func (e *APPError) Error() string {
	return e.Message
}

func NewValidationError(message string, data any) *APPError {
	return NewAPPError(http.StatusBadRequest, message, ValidationError, data)
}

func NewNotFoundError(message string, data any) *APPError {
	return NewAPPError(http.StatusNotFound, message, NotFoundError, data)
}

func NewConflictError(message string, data any) *APPError {
	return NewAPPError(http.StatusConflict, message, ConflictError, data)
}

func NewUnauthorizedError(message string, data any) *APPError {
	return NewAPPError(http.StatusUnauthorized, message, UnauthorizedError, data)
}

func NewInternalError(message string, data any) *APPError {
	return NewAPPError(http.StatusInternalServerError, message, InternalError, data)
}
