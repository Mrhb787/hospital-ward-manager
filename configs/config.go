package configs

import (
	"fmt"
	"os"

	"github.com/Mrhb787/hospital-ward-manager/common"
)

type ServiceConfig struct {
	ServicePort string
}

type APIConfig struct {
	ApiKey string
}

type DBConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

type AppConfig struct {

	// api configs
	APIConfig

	// database configs
	DBConfig

	// service configs
	ServiceConfig
}

func (db DBConfig) new() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "postgres://neondb_owner:npg_E7DMQAm4ayRf@ep-quiet-frost-a1oukqpx-pooler.ap-southeast-1.aws.neon.tech/neondb?sslmode=require"),
		Name:     getEnv("DB_NAME", ""),
		User:     getEnv("DB_ADMIN", ""),
		Password: getEnv("DB_ADMIN_PASSWORD", ""),
	}
}

func (api APIConfig) new() APIConfig {
	return APIConfig{
		ApiKey: getEnv("API_SECRET_KEY", ""),
	}
}

func (s ServiceConfig) new() ServiceConfig {
	return ServiceConfig{
		ServicePort: string(common.LOCAL_SERVICE_PORT),
	}
}

func New() *AppConfig {
	return &AppConfig{
		DBConfig:      DBConfig{}.new(),
		APIConfig:     APIConfig{}.new(),
		ServiceConfig: ServiceConfig{}.new(),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	fmt.Println("key", key, "value", value)

	if value == "" {
		return defaultValue
	}
	return value
}
