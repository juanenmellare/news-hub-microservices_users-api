package factories

import (
	"github.com/stretchr/testify/assert"
	"news-hub-microservices_users-api/controllers"
	"testing"
)

func TestNewDomainLayersFactory(t *testing.T) {
	domainLayersFactory := NewControllersFactory(nil)

	assert.Implements(t, (*ControllersFactory)(nil), domainLayersFactory)
}

func Test_domainLayersFactoryImpl_GetHealthChecksController(t *testing.T) {
	domainLayersFactory := NewControllersFactory(nil)

	assert.Implements(t, (*controllers.HealthChecksController)(nil), domainLayersFactory.GetHealthChecksController())
}
