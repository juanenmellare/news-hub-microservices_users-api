package models

import "github.com/gofrs/uuid"

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	FirstName string    `json:"firstName" gorm:"column:first_name"`
	LastName  string    `json:"lastName" gorm:"column:last_name"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	Salt      string    `json:"salt" gorm:"column:salt"`
}
