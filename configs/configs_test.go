package configs

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ConfigImpl_GetPort(t *testing.T) {
	expectedValue := "0000"
	_ = os.Setenv("PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetPort())
}

func Test_ConfigImpl_GetPort_default(t *testing.T) {
	_ = os.Unsetenv("PORT")

	config := NewConfig()

	assert.Equal(t, "8081", config.GetPort())
}

func Test_ConfigImpl_GetDatabasePort(t *testing.T) {
	expectedValue := "5431"
	_ = os.Setenv("DATABASE_PORT", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePort())
}

func Test_ConfigImpl_GetDatabasePort_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PORT")

	config := NewConfig()

	assert.Equal(t, "5432", config.GetDatabasePort())
}

func TestConfigImpl_GetDatabaseHost(t *testing.T) {
	expectedValue := "foo-host"
	_ = os.Setenv("DATABASE_HOST", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseHost())
}

func TestConfigImpl_GetDatabaseHost_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_HOST")

	config := NewConfig()

	assert.Equal(t, "localhost", config.GetDatabaseHost())
}

func TestConfigImpl_GetDatabaseName(t *testing.T) {
	expectedValue := "foo-name"
	_ = os.Setenv("DATABASE_NAME", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseName())
}

func TestConfigImpl_GetDatabaseName_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_NAME")

	config := NewConfig()

	assert.Equal(t, "development.news-hub_users-api", config.GetDatabaseName())
}

func TestConfigImpl_GetDatabaseUser(t *testing.T) {
	expectedValue := "foo-user"
	_ = os.Setenv("DATABASE_USER", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabaseUser())
}

func TestConfigImpl_GetDatabaseUser_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_USER")

	config := NewConfig()

	assert.Equal(t, "admin", config.GetDatabaseUser())
}

func TestConfigImpl_GetDatabasePass(t *testing.T) {
	expectedValue := "foo-pass"
	_ = os.Setenv("DATABASE_PASS", expectedValue)

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetDatabasePass())
}

func TestConfigImpl_GetDatabasePass_default(t *testing.T) {
	_ = os.Unsetenv("DATABASE_PASS")

	config := NewConfig()

	assert.Equal(t, "news-hub.2022", config.GetDatabasePass())
}

func TestConfigImpl_GetBCryptCost(t *testing.T) {
	expectedValue := bcrypt.DefaultCost
	_ = os.Setenv("BCRYPT_COST", strconv.Itoa(expectedValue))

	config := NewConfig()

	assert.Equal(t, expectedValue, config.GetBCryptCost())
}

func TestConfigImpl_GetBCryptCost_Panic_invalid_int(t *testing.T) {
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

func TestConfigImpl_GetBCryptCost_default(t *testing.T) {
	_ = os.Unsetenv("BCRYPT_COST")

	config := NewConfig()

	assert.Equal(t, bcrypt.MinCost, config.GetBCryptCost())
}

func Test_configImpl_GetTokenUserSecretKey(t *testing.T) {
	_ = os.Unsetenv("USER_TOKEN_SECRET_KEY")

	config := NewConfig()

	assert.Equal(t, "foo", config.GetTokenUserSecretKey())
}

func Test_configImpl_GetTokenUserExpirationHours(t *testing.T) {
	_ = os.Unsetenv("USER_TOKEN_EXPIRATION_HOURS")

	config := NewConfig()

	assert.Equal(t, 1, config.GetTokenUserExpirationHours())
}
