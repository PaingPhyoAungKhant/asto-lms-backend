// Package config load the environment variables
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string
	Host string
}

// Load - Load & return config
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("GATEWAY_PORT", "8080"),
			Host: getEnv("GATEWAY_HOST", "localhost"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
