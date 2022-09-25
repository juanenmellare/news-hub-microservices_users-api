package configs

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Config_GetPort(t *testing.T) {
	expectedValue := "0000"
	_ = os.Setenv("PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetPort())
}

func Test_Config_GetPort_default(t *testing.T) {
	_ = os.Unsetenv("PORT")

	config := NewConfig()

	assert.Equal(t, "8081", config.GetPort())
}

func Test_Config_GetDatabasePort(t *testing.T) {
	expectedValue := "5431"
	_ = os.Setenv("DATABASE_PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePort())
}

func Test_Config_GetDatabasePort_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PORT")

	config := NewConfig()

	assert.Equal(t, "5432", config.GetDatabasePort())
}

func TestConfig_GetDatabaseHost(t *testing.T) {
	expectedValue := "foo-host"
	_ = os.Setenv("DATABASE_HOST", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseHost())
}

func TestConfig_GetDatabaseHost_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_HOST")

	config := NewConfig()

	assert.Equal(t, "localhost", config.GetDatabaseHost())
}

func TestConfig_GetDatabaseName(t *testing.T) {
	expectedValue := "foo-name"
	_ = os.Setenv("DATABASE_NAME", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseName())
}

func TestConfig_GetDatabaseName_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_NAME")

	config := NewConfig()

	assert.Equal(t, "development.news-hub_users-api", config.GetDatabaseName())
}

func TestConfig_GetDatabaseUser(t *testing.T) {
	expectedValue := "foo-user"
	_ = os.Setenv("DATABASE_USER", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseUser())
}

func TestConfig_GetDatabaseUser_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_USER")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetDatabaseUser())
}

func TestConfig_GetDatabasePass(t *testing.T) {
	expectedValue := "foo-pass"
	_ = os.Setenv("DATABASE_PASS", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePass())
}

func TestConfig_GetDatabasePass_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PASS")

	config := NewConfig()

	assert.Equal(t, "news-hub.2022", config.GetDatabasePass())
}

func TestConfig_GetBCryptCost(t *testing.T) {
	expectedValue := bcrypt.DefaultCost
	_ = os.Setenv("BCRYPT_COST", strconv.Itoa(expectedValue))

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetBCryptCost())
}

func TestConfig_GetBCryptCost_Panic_invalid_int(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("did not panic")
		} else {
			assert.Equal(t, "10c is not a valid int", r)
		}
	}()

	expectedValue := "10c"
	_ = os.Setenv("BCRYPT_COST", expectedValue)

	config := NewConfig()

	assert.Equal(t, bcrypt.MinCost, config.GetBCryptCost())
}

func TestConfig_GetBCryptCost_default(t *testing.T) {
	_ = os.Unsetenv("BCRYPT_COST")

	config := NewConfig()

	assert.Equal(t, bcrypt.MinCost, config.GetBCryptCost())
}

func Test_config_GetTokenUserSecretKey(t *testing.T) {
	_ = os.Unsetenv("USER_TOKEN_SECRET_KEY")

	config := NewConfig()

	assert.Equal(t, "foo", config.GetTokenUserSecretKey())
}

func Test_config_GetTokenUserExpirationHours(t *testing.T) {
	_ = os.Unsetenv("USER_TOKEN_EXPIRATION_HOURS")

	config := NewConfig()

	assert.Equal(t, 1, config.GetTokenUserExpirationHours())
}
