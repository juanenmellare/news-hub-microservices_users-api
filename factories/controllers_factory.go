package factories

import (
	"news-hub-microservices_users-api/controllers"
	"news-hub-microservices_users-api/databases"
)

type ControllersFactory interface {
	GetHealthChecksController() controllers.HealthChecksController
}

type controllersFactoryImpl struct {
	healthChecksController controllers.HealthChecksController
}

func NewControllersFactory(_ databases.RelationalDatabase) ControllersFactory {
	return &controllersFactoryImpl{
		healthChecksController: controllers.NewHealthChecksController(),
	}
}

func (c controllersFactoryImpl) GetHealthChecksController() controllers.HealthChecksController {
	return c.healthChecksController
}
