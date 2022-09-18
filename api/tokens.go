package api

import (
	"net/http"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/models"
	"strings"
	"time"
)

import "github.com/golang-jwt/jwt/v4"

type Token interface {
	Verify(userTokenSecretKey string, r *http.Request)
	ToString(userTokenSecretKey string) string
}

type tokenImpl struct {
	token  *jwt.Token
	claims *jwt.MapClaims
}

func getTokenFromRequest(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	bearTokenParts := strings.Split(bearToken, " ")
	if len(bearTokenParts) != 2 {
		panic(errors.NewBadRequestApiError("invalid token"))
	}
	return bearTokenParts[1]
}

func (t *tokenImpl) Verify(userTokenSecretKey string, r *http.Request) {
	tokenString := getTokenFromRequest(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(userTokenSecretKey), nil
	})
	if err != nil {
		panic(errors.NewBadRequestApiError(err.Error()))
	}

	t.token = token

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		t.claims = &claims
	}
}

func (t *tokenImpl) ToString(userTokenSecretKey string) string {
	secretKey := []byte(userTokenSecretKey)
	tokenString, err := t.token.SignedString(secretKey)
	if err != nil {
		panic(errors.NewBadRequestApiError(err.Error()))
	}

	return tokenString
}

type UserToken struct {
	tokenImpl
}

func (t UserToken) GetUserId() string {
	if userId, ok := (*t.claims)["userId"].(string); ok {
		return userId
	}
	panic(errors.NewBadRequestApiError("token missing claim: userId"))
}

func (t UserToken) IsExpired() {
	if expirationUnix, ok := (*t.claims)["expiration"].(float64); ok {
		if time.Now().Unix() > int64(expirationUnix) {
			panic(errors.NewBadRequestApiError("token has expired"))
		}
	} else {
		panic(errors.NewBadRequestApiError("token missing claim: expiration"))
	}
}

func NewUserToken(userTokenExpirationHours int, user *models.User) UserToken {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.ID
	claims["expiration"] = time.Now().Add(time.Duration(userTokenExpirationHours) * time.Hour).Unix()

	return UserToken{
		tokenImpl{token, nil},
	}
}
