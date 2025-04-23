package core

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	SSLMode   string
	AppPort   string
	JWTSecret string
	JWTExp    string
}

var configInstance *Config

func ConfigInstance() *Config {
	if configInstance == nil {
		initConfig()
	}
	return configInstance
}

func initConfig() {
	required := []string{
		"DB_HOST", "DB_PORT", "DB_USER", "DB_NAME", "SSL_MODE", "APP_PORT",
	}

	for _, k := range required {
		if os.Getenv(k) == "" {
			panic(fmt.Sprintf("missing required env var: %s", k))
		}
	}

	configInstance = &Config{
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASSWORD"),
		DBName:    os.Getenv("DB_NAME"),
		SSLMode:   os.Getenv("SSL_MODE"),
		AppPort:   os.Getenv("APP_PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExp:    os.Getenv("JWT_EXPIRATION"),
	}
}
