package repositories

import "news-hub-microservices_users-api/databases"

type UsersRepository interface {
}

type usersRepositoryImpl struct {
	relationalDatabase databases.RelationalDatabase
}

func NewUsersRepository(relationalDatabase databases.RelationalDatabase) UsersRepository {
	return &usersRepositoryImpl{
		relationalDatabase,
	}
}
