package factories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/controllers"
	"testing"
)

func buildDomainLayersFactory() ControllersFactory {
	return NewControllersFactory(nil)
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
