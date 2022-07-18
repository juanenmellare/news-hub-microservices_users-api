package configs

import (
	"os"
)

type Config interface {
	GetPort() string
	GetDatabaseHost() string
	GetDatabaseName() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePass() string
}

type ConfigImpl struct {
	port         string
	databaseHost string
	databaseName string
	databasePort string
	databaseUser string
	databasePass string
}

func NewConfig() Config {
	return &ConfigImpl{
		port:         getValueOrDefault("PORT", "8081"),
		databaseHost: getValueOrDefault("DATABASE_HOST", "localhost"),
		databaseName: getValueOrDefault("DATABASE_NAME", "development.news-hub_users-api"),
		databasePort: getValueOrDefault("DATABASE_PORT", "5432"),
		databaseUser: getValueOrDefault("DATABASE_USER", "admin"),
		databasePass: getValueOrDefault("DATABASE_PASS", "news-hub.2022"),
	}
}

func getValueOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func (c ConfigImpl) GetPort() string {
	return c.port
}

func (c ConfigImpl) GetDatabaseHost() string {
	return c.databaseHost
}

func (c ConfigImpl) GetDatabaseName() string {
	return c.databaseName
}

func (c ConfigImpl) GetDatabasePort() string {
	return c.databasePort
}

func (c ConfigImpl) GetDatabaseUser() string {
	return c.databaseUser
}

func (c ConfigImpl) GetDatabasePass() string {
	return c.databasePass
}
