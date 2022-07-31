package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/utils"
	"strings"
	"testing"
)

func Test_CreateUserRequest_MarshallAndValidate_Ok(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)

	createUserRequest := &CreateUserRequest{
		FirstName:      utils.NewStringPointer("firstName"),
		LastName:       utils.NewStringPointer("lastName"),
		Email:          utils.NewStringPointer("email"),
		Password:       utils.NewStringPointer("password"),
		PasswordRepeat: utils.NewStringPointer("password"),
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(createUserRequest)
	context.Request, _ = http.NewRequest(http.MethodPost, "/", reqBodyBytes)

	createUserRequest.MarshallAndValidate(context)
}

func Test_CreateUserRequest_MarshallAndValidate_Panic_bindJSON(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			apiErr, _ := r.(*errors.ApiError)
			assert.Equal(t, 400, apiErr.Code)
			assert.Equal(t, "unexpected EOF", apiErr.Message)
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	context.Request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader("{"))

	createUserRequest := &CreateUserRequest{}
	createUserRequest.MarshallAndValidate(context)
}

func Test_CreateUserRequest_MarshallAndValidate_Panic_passwordMatch(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			apiErr, _ := r.(*errors.ApiError)
			assert.Equal(t, 400, apiErr.Code)
			assert.Equal(t, "the fields 'password' and 'passwordRepeat' doesn't match", apiErr.Message)
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)

	createUserRequest := &CreateUserRequest{
		FirstName:      utils.NewStringPointer("firstName"),
		LastName:       utils.NewStringPointer("lastName"),
		Email:          utils.NewStringPointer("email"),
		Password:       utils.NewStringPointer("password"),
		PasswordRepeat: utils.NewStringPointer("password."),
	}

	_ = json.NewEncoder(reqBodyBytes).Encode(createUserRequest)
	context.Request, _ = http.NewRequest(http.MethodPost, "/", reqBodyBytes)

	createUserRequest.MarshallAndValidate(context)
}

func Test_CreateUserRequest_MarshallAndValidate_Panic_missingFields(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			apiErr, _ := r.(*errors.ApiError)
			assert.Equal(t, 400, apiErr.Code)
			assert.Equal(t, "the fields 'firstName, lastName, email, password, passwordRepeat' should not be empty", apiErr.Message)
		}
	}()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)
	reqBodyBytes := new(bytes.Buffer)

	createUserRequest := &CreateUserRequest{}

	_ = json.NewEncoder(reqBodyBytes).Encode(createUserRequest)
	context.Request, _ = http.NewRequest(http.MethodPost, "/", reqBodyBytes)

	createUserRequest.MarshallAndValidate(context)
}
