package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type ApiError struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewError(message string) error {
	return errors.New(message)
}

func newApiError(code int, message string) *ApiError {
	return &ApiError{
		Status:  http.StatusText(code),
		Code:    code,
		Message: message,
	}
}

func NewInternalServerApiError(message string) *ApiError {
	return newApiError(http.StatusInternalServerError, message)
}

func NewBadRequestApiError(message string) *ApiError {
	return newApiError(http.StatusBadRequest, message)
}

func NewNotFoundError(message string) *ApiError {
	return newApiError(http.StatusNotFound, message)
}

func NewRequestFieldsShouldNotBeEmptyError(fields []string) *ApiError {
	fieldString := strings.Join(fields, ", ")
	grammaticalNumber := "field"
	if len(fields) > 1 {
		grammaticalNumber = grammaticalNumber + "s"
	}
	return NewBadRequestApiError(fmt.Sprintf("the %s '%s' should not be empty", grammaticalNumber, fieldString))
}
