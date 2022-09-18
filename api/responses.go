package api

import (
	"fmt"
	"github.com/gofrs/uuid"
)

type CreateUserResponse struct {
	UserId string `json:"userId"`
}

func NewCreateUserResponse(userId uuid.UUID) *CreateUserResponse {
	return &CreateUserResponse{
		UserId: userId.String(),
	}
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func NewAuthenticateResponse(token UserToken, userTokenSecretKey string) *AuthenticateResponse {
	return &AuthenticateResponse{
		Token: fmt.Sprintf("Bearer %s", token.ToString(userTokenSecretKey)),
	}
}

type GetResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func NewGetResponse(firstName, lastName, email string) *GetResponse {
	return &GetResponse{
		firstName,
		lastName,
		email,
	}
}
