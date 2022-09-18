package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_NewApiError(t *testing.T) {
	code := http.StatusInternalServerError
	err := errors.New("panic")

	apiError := newApiError(code, err.Error())

	assert.Equal(t, "Internal Server Error", apiError.Status)
	assert.Equal(t, code, apiError.Code)
	assert.Equal(t, err.Error(), apiError.Message)
}

func Test_NewInternalServerApiError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewInternalServerApiError(err.Error())

	assert.Equal(t, "Internal Server Error", apiError.Status)
	assert.Equal(t, http.StatusInternalServerError, apiError.Code)
	assert.Equal(t, err.Error(), apiError.Message)
}

func Test_NewBadRequestApiError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewBadRequestApiError(err.Error())

	assert.Equal(t, "Bad Request", apiError.Status)
	assert.Equal(t, http.StatusBadRequest, apiError.Code)
	assert.Equal(t, err.Error(), apiError.Message)
}

func Test_NewError(t *testing.T) {
	message := "panic"
	err := NewError(message)

	assert.NotNil(t, err)
	assert.Equal(t, message, err.Error())
}

func Test_NewNotFoundError(t *testing.T) {
	err := errors.New("panic")

	apiError := NewNotFoundError(err.Error())

	assert.Equal(t, "Not Found", apiError.Status)
	assert.Equal(t, http.StatusNotFound, apiError.Code)
	assert.Equal(t, err.Error(), apiError.Message)
}

func Test_NewRequestFieldsShouldNotBeEmptyError_oneField(t *testing.T) {
	fields := []string{"foo"}

	apiError := NewRequestFieldsShouldNotBeEmptyError(fields)

	assert.Equal(t, "Bad Request", apiError.Status)
	assert.Equal(t, http.StatusBadRequest, apiError.Code)
	assert.Equal(t, "the field 'foo' should not be empty", apiError.Message)
}

func Test_NewRequestFieldsShouldNotBeEmptyError_multiFields(t *testing.T) {
	fields := []string{"foo", "foo2"}

	apiError := NewRequestFieldsShouldNotBeEmptyError(fields)

	assert.Equal(t, "Bad Request", apiError.Status)
	assert.Equal(t, http.StatusBadRequest, apiError.Code)
	assert.Equal(t, "the fields 'foo, foo2' should not be empty", apiError.Message)
}

func TestNewAlreadyExistModelError(t *testing.T) {
	err := NewAlreadyExistModelError("foo")

	assert.Equal(t, "foo already exist", err.Message)
}

func TestNewInvalidEmailOrPasswordError(t *testing.T) {
	err := NewInvalidEmailOrPasswordError()

	assert.Equal(t, "Invalid Email or Password", err.Message)
}
