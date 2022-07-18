package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"news-hub-microservices_users-api/configs"
	"news-hub-microservices_users-api/databases"
	"news-hub-microservices_users-api/factories"
	"news-hub-microservices_users-api/router"
)

func main() {
	logger := log.Default()
	config := configs.NewConfig()

	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		config.GetDatabaseUser(), config.GetDatabasePass(), config.GetDatabaseHost(),
		config.GetDatabasePort(), config.GetDatabaseName())

	relationalDatabase := databases.NewConnection(gorm.Open(postgres.Open(connectionString), &gorm.Config{}))
	relationalDatabase.DoMigration()

	domainLayersFactory := factories.NewControllersFactory(relationalDatabase)

	port := ":" + config.GetPort()
	if err := router.New(domainLayersFactory).Run(port); err != nil {
		logger.Fatalf("Error while trying to create the router: " + err.Error())
	}
}
