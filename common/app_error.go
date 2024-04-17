package common

import (
	"errors"
	"net/http"
)

const (
	DB_ERROR     = "DB_ERROR"
	SERVER_ERROR = "INTERNAL_SERVER_ERROR"
	BAD_REQUEST  = "BAD_REQUEST"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"root_error"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewAppErrorResponse(statusCode int, rootErr error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(message string, rootErr error, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorizedResponse(message string, rootErr error, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewInternalErrorResponse(message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func NewBadRequestResponse(message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

func ErrDB(err error) *AppError {
	return NewAppErrorResponse(http.StatusInternalServerError, nil, "Something went wrong with DB", err.Error(), DB_ERROR)
}

func ErrBadRequest(err error) *AppError {
	return NewBadRequestResponse("The request was malformed or contained invalid parameters", err.Error(), BAD_REQUEST)
}

func ErrInternal(err error) *AppError {
	return NewInternalErrorResponse("Something went wrong with server", err.Error(), SERVER_ERROR)
}

var RecordNotFound = errors.New("Record not found")
