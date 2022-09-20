package factories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/internal/controllers"
	"testing"
)

func buildDomainLayersFactory() ControllersFactory {
	config := configs.NewConfig()
	return NewControllersFactory(nil, config)
}

func TestNewDomainLayersFactory(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*ControllersFactory)(nil), domainLayersFactory)
}

func Test_domainLayersFactoryImpl_GetHealthChecksController(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*controllers.HealthChecksController)(nil), domainLayersFactory.GetHealthChecksController())
}

func Test_domainLayersFactoryImpl_GetUsersController(t *testing.T) {
	domainLayersFactory := buildDomainLayersFactory()

	assert.Implements(t, (*controllers.UsersController)(nil), domainLayersFactory.GetUsersController())
}
