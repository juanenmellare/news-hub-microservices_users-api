package factories

import (
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/controllers"
	"news-hub-microservices_users-api/databases"
	"news-hub-microservices_users-api/repositories"
	"news-hub-microservices_users-api/services"
)

func buildHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func buildUsersController(relationalDatabase databases.RelationalDatabase, config configs.Config) controllers.UsersController {
	usersRepository := repositories.NewUsersRepository(relationalDatabase)
	userService := services.NewUsersService(usersRepository, config.GetBCryptCost())
	usersController := controllers.NewUsersController(userService)

	return usersController
}

type ControllersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
	GetUsersController() controllers.UsersController
}

type controllersFactoryImpl struct {
	healthChecksController controllers.HealthChecksController
	usersController        controllers.UsersController
}

func NewControllersFactory(relationalDatabase databases.RelationalDatabase, config configs.Config) ControllersFactory {
	return &controllersFactoryImpl{
		healthChecksController: buildHealthChecksController(),
		usersController:        buildUsersController(relationalDatabase, config),
	}
}

func (c controllersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return c.healthChecksController
}

func (c controllersFactoryImpl) GetUsersController() controllers.UsersController {
	return c.usersController
}
