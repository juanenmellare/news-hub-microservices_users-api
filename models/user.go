package models

import (
	"github.com/gofrs/uuid"
	"news-hub-microservices_users-api/utils"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
}

func NewUser(firstName, lastName, email, password string, cost int) *User {
	passwordHashed := utils.HashPassword(password, cost)
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  passwordHashed,
	}
}
