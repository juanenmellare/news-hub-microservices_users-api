package rest

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"news-hub-microservices_users-api/internal/errors"
	"news-hub-microservices_users-api/test/mocks/models"
	"testing"
)

func assertTokenPanic(t *testing.T, recover interface{}, message string) {
	if r := recover; r != nil {
		assert.Equal(t, &errors.ApiError{Code: 400, Status: "Bad Request", Message: message}, r)
	} else {
		t.Errorf("did not panic")
	}
}

func TestNewUserToken(t *testing.T) {
	userMock := models.NewUserBuilder().Build()

	token := NewUserToken(1, &userMock)

	assert.NotNil(t, token.token)
	assert.NotNil(t, token.claims)
}

func TestUserToken_Verify_GetUserId(t *testing.T) {
	userMock := models.NewUserBuilder().Build()

	token := NewUserToken(1, &userMock)
	secretKey := ""
	tokenString := token.ToString(secretKey)

	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	token.Verify(secretKey, request)

	assert.Equal(t, userMock.ID.String(), token.GetUserId())
}

func TestUserToken_Verify_GetUserId_error_verifyUnexpectedError(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "token contains an invalid number of segments") }()

	userMock := models.NewUserBuilder().Build()

	token := NewUserToken(1, &userMock)
	secretKey := ""
	tokenString := token.ToString(secretKey)

	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	request.Header.Add("Authorization", fmt.Sprintf("%s foo", tokenString))
	token.Verify(secretKey, request)
}

func TestUserToken_Verify_GetUserId_error_ToString(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "the requested hash function is unavailable") }()

	userMock := models.NewUserBuilder().Build()

	token := NewUserToken(1, &userMock)
	token.token.Method = &jwt.SigningMethodHMAC{}

	secretKey := ""
	token.ToString(secretKey)
}

func TestUserToken_Verify_GetUserId_error(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "token missing claim: userId") }()

	userToken := jwt.New(jwt.SigningMethodHS256)
	claims := userToken.Claims.(jwt.MapClaims)

	token := UserToken{
		tokenImpl{userToken, &claims},
	}

	secretKey := ""
	tokenString := token.ToString(secretKey)

	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tokenString))
	token.Verify(secretKey, request)

	token.GetUserId()
}

func TestUserToken_Verify_GetUserId_error_getTokenFromRequest(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "invalid token") }()

	userMock := models.NewUserBuilder().Build()

	token := NewUserToken(1, &userMock)
	secretKey := ""
	tokenString := token.ToString(secretKey)

	request, _ := http.NewRequest(http.MethodPost, "/", nil)
	request.Header.Add("Authorization", tokenString)
	token.Verify(secretKey, request)
}

func TestUserToken_IsExpired_false(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	userMock := models.NewUserBuilder().Build()
	token := NewUserToken(1, &userMock)

	token.IsExpired()
}

func TestUserToken_IsExpired_missing_claim(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "token missing claim: expiration") }()

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	userToken := UserToken{tokenImpl{token, &claims}}

	userToken.IsExpired()
}

func TestUserToken_IsExpired_true_panic(t *testing.T) {
	defer func() { assertTokenPanic(t, recover(), "token has expired") }()

	userMock := models.NewUserBuilder().Build()
	token := NewUserToken(-1, &userMock)

	token.IsExpired()
}
