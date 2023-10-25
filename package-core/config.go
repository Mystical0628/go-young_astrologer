package package_core

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort int
	NasaApiKey string
	DB         ConfigDB
}

type ConfigDB struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: getEnvAsInt("SERVER_PORT", 80),
		NasaApiKey: getEnv("NASA_API_KEY", ""),
		DB: ConfigDB{
			Driver:   getEnv("DB_DRIVER", "postgres"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Database: getEnv("DB_DATABASE", "postgres"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
