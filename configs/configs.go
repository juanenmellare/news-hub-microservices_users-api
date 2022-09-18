package configs

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

type Config interface {
	GetPort() string
	GetDatabaseHost() string
	GetDatabaseName() string
	GetDatabasePort() string
	GetDatabaseUser() string
	GetDatabasePass() string
	GetBCryptCost() int
	GetTokenUserSecretKey() string
	GetTokenUserExpirationHours() int
}

type configImpl struct {
	port                     string
	databaseHost             string
	databaseName             string
	databasePort             string
	databaseUser             string
	databasePass             string
	bCryptCost               int
	userTokenSecretKey       string
	userTokenExpirationHours int
}

func NewConfig() Config {
	return &configImpl{
		port:                     getStringValueOrDefault("PORT", "8081"),
		databaseHost:             getStringValueOrDefault("DATABASE_HOST", "localhost"),
		databaseName:             getStringValueOrDefault("DATABASE_NAME", "development.news-hub_users-api"),
		databasePort:             getStringValueOrDefault("DATABASE_PORT", "5432"),
		databaseUser:             getStringValueOrDefault("DATABASE_USER", "admin"),
		databasePass:             getStringValueOrDefault("DATABASE_PASS", "news-hub.2022"),
		bCryptCost:               getIntValueOrDefault("BCRYPT_COST", bcrypt.MinCost),
		userTokenSecretKey:       getStringValueOrDefault("USER_TOKEN_SECRET_KEY", "foo"),
		userTokenExpirationHours: getIntValueOrDefault("USER_TOKEN_EXPIRATION_HOURS", 1),
	}
}

func getStringValueOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func getIntValueOrDefault(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("%s is not a valid int", value))
	}

	return intValue
}

func (c configImpl) GetPort() string {
	return c.port
}

func (c configImpl) GetDatabaseHost() string {
	return c.databaseHost
}

func (c configImpl) GetDatabaseName() string {
	return c.databaseName
}

func (c configImpl) GetDatabasePort() string {
	return c.databasePort
}

func (c configImpl) GetDatabaseUser() string {
	return c.databaseUser
}

func (c configImpl) GetDatabasePass() string {
	return c.databasePass
}

func (c configImpl) GetBCryptCost() int {
	return c.bCryptCost
}

func (c configImpl) GetTokenUserSecretKey() string {
	return c.userTokenSecretKey
}

func (c configImpl) GetTokenUserExpirationHours() int {
	return c.userTokenExpirationHours
}
