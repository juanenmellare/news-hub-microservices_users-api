package repositories

import (
	"news-hub-microservices_users-api/internal/databases"
	"news-hub-microservices_users-api/internal/models"
)

type UsersRepository interface {
	Create(user *models.User)
	FindByEmail(email string) *models.User
	FindById(id string) *models.User
}

type usersRepository struct {
	relationalDatabase databases.RelationalDatabase
}

func (u usersRepository) Create(user *models.User) {
	result := u.relationalDatabase.Get().Create(user)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (u usersRepository) FindByEmail(email string) *models.User {
	var user models.User
	result := u.relationalDatabase.Get().First(&user, "email = ?", email)
	if result.Error != nil {
		return nil
	}
	return &user
}

func (u usersRepository) FindById(id string) *models.User {
	var user models.User
	result := u.relationalDatabase.Get().First(&user, "id = ?", id)
	if result.Error != nil {
		return nil
	}
	return &user
}

func NewUsersRepository(relationalDatabase databases.RelationalDatabase) UsersRepository {
	return &usersRepository{
		relationalDatabase,
	}
}
