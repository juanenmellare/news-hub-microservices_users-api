package api

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/internal/errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"news-hub-microservices_users-api/internal/factories"
)

func Test_New(t *testing.T) {
	config := configs.NewConfig()
	DomainLayersFactory := factories.NewControllersFactory(nil, config)
	engine := NewRouter(DomainLayersFactory, configs.NewConfig())
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

func Test_BasicAuthenticationMiddleware_Unauthorized_empty(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{}

	BasicAuthenticationMiddleware(c, "admin", "password")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401,\"status\":\"Unauthorized\",\"message\":\"invalid credentials\"}", w.Body.String())
}

func Test_BasicAuthenticationMiddleware_Unauthorized_basic_incomplete(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: map[string][]string{},
	}
	c.Request.Header.Set("Authorization", "Basic")

	BasicAuthenticationMiddleware(c, "admin", "password")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401,\"status\":\"Unauthorized\",\"message\":\"invalid credentials\"}", w.Body.String())
}

func Test_BasicAuthenticationMiddleware_Unauthorized_basic_base64_bad_encode(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: map[string][]string{},
	}
	c.Request.Header.Set("Authorization", "Basic ;;;;;;")

	BasicAuthenticationMiddleware(c, "admin", "password")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401,\"status\":\"Unauthorized\",\"message\":\"invalid credentials\"}", w.Body.String())
}

func Test_BasicAuthenticationMiddleware_Unauthorized_basic_base64_bad_format(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: map[string][]string{},
	}
	c.Request.Header.Set("Authorization", "Basic Zm9v")

	BasicAuthenticationMiddleware(c, "admin", "password")

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401,\"status\":\"Unauthorized\",\"message\":\"invalid credentials\"}", w.Body.String())
}

func Test_BasicAuthenticationMiddleware(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Header: map[string][]string{},
	}
	c.Request.Header.Set("Authorization", "Basic YWRtaW46cGFzc3dvcmQ=, Bearer foo")

	BasicAuthenticationMiddleware(c, "admin", "password")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Bearer foo", c.Request.Header.Get("Authorization"))
}

func TestNewRouter(t *testing.T) {
	controllersFactoryMock := factories.NewControllersFactory(nil, configs.NewConfig())

	gin.SetMode(gin.TestMode)
	router := NewRouter(controllersFactoryMock, configs.NewConfig())

	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/v1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":401,\"status\":\"Unauthorized\",\"message\":\"invalid credentials\"}", w.Body.String())
}
