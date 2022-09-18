package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"news-hub-microservices_users-api/factories"
)

func Test_New(t *testing.T) {
	config := configs.NewConfig()
	DomainLayersFactory := factories.NewControllersFactory(nil, config)
	engine := New(DomainLayersFactory)
	s := httptest.NewServer(engine)

	response, _ := http.Get(fmt.Sprintf("%s/ping", s.URL))

	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(response.Body)
	responseBodyString := buf.String()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "{\"message\":\"pong\"}", responseBodyString)

	s.Close()
}

func Test_HandlePanicRecoveryMiddleware_apiError_NewAlreadyExistModelError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HandlePanicRecoveryMiddleware(c, errors.NewAlreadyExistModelError("foo"))

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"code\":400,\"status\":\"Bad Request\",\"message\":\"foo already exist\"}", w.Body.String())
}

func Test_HandlePanicRecoveryMiddleware_apiError_NewInvalidEmailOrPasswordError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HandlePanicRecoveryMiddleware(c, errors.NewInvalidEmailOrPasswordError())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"code\":400,\"status\":\"Bad Request\",\"message\":\"Invalid Email or Password\"}", w.Body.String())
}

func Test_HandlePanicRecoveryMiddleware_unexpected_apiError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HandlePanicRecoveryMiddleware(c, errors.NewBadRequestApiError("error"))

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "{\"code\":400,\"status\":\"Bad Request\",\"message\":\"error\"}", w.Body.String())
}

func Test_HandlePanicRecoveryMiddleware_unexpected_error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	HandlePanicRecoveryMiddleware(c, errors.NewError("error"))

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"code\":500,\"status\":\"Internal Server Error\",\"message\":\"unexpected error: error\"}", w.Body.String())
}

func Test_HandlePanicRecoveryMiddleware_unhandled_error(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	type unhandledStruct struct {
		Message string
	}

	HandlePanicRecoveryMiddleware(c, unhandledStruct{Message: "error"})

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "{\"code\":500,\"status\":\"Internal Server Error\",\"message\":\"unhandled error: {error}\"}", w.Body.String())
}
