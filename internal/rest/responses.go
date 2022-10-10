package rest

import (
	"fmt"
	"github.com/gofrs/uuid"
	"news-hub-microservices_users-api/internal/models"
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

func NewAuthenticateResponse(token string) *AuthenticateResponse {
	return &AuthenticateResponse{
		Token: fmt.Sprintf("Bearer %s", token),
	}
}

type GetResponse struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func NewGetResponse(user *models.User) *GetResponse {
	return &GetResponse{
		user.ID.String(),
		user.FirstName,
		user.LastName,
		user.Email,
	}
}
