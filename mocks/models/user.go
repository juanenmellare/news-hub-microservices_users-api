package models

import (
	"github.com/gofrs/uuid"
	"news-hub-microservices_users-api/models"
)

type UserBuilder struct {
	user models.User
}

func (u UserBuilder) Build() models.User {
	return u.user
}

func NewUserBuilder() *UserBuilder {
	uuidMock, _ := uuid.NewV4()

	return &UserBuilder{
		user: models.User{
			ID:        uuidMock,
			FirstName: "foo-firstname",
			LastName:  "foo-lastname",
			Email:     "foo-email@email.com",
			Password:  "password",
			Salt:      "10",
		},
	}
}
