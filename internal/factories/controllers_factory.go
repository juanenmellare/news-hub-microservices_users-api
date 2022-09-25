package factories

import (
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/internal/controllers"
	"news-hub-microservices_users-api/internal/databases"
	"news-hub-microservices_users-api/internal/repositories"
	"news-hub-microservices_users-api/internal/services"
)

func buildHealthChecksController() controllers.HealthChecksController {
	return controllers.NewHealthChecksController()
}

func buildUsersController(relationalDatabase databases.RelationalDatabase, config configs.Config) controllers.UsersController {
	usersRepository := repositories.NewUsersRepository(relationalDatabase)
	userService := services.NewUsersService(usersRepository, config.GetBCryptCost())
	usersController := controllers.NewUsersController(userService, config.GetTokenUserSecretKey(), config.GetTokenUserExpirationHours())

	return usersController
}

type ControllersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
	GetUsersController() controllers.UsersController
}

type controllersFactory struct {
	healthChecksController controllers.HealthChecksController
	usersController        controllers.UsersController
}

func NewControllersFactory(relationalDatabase databases.RelationalDatabase, config configs.Config) ControllersFactory {
	return &controllersFactory{
		healthChecksController: buildHealthChecksController(),
		usersController:        buildUsersController(relationalDatabase, config),
	}
}

func (c controllersFactory) GetHealthChecksController() controllers.HealthChecksController {
	return c.healthChecksController
}

func (c controllersFactory) GetUsersController() controllers.UsersController {
	return c.usersController
}
