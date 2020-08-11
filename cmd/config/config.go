package config

import (
	"os"
)

// Config contains all the neccessary configurations
type Config struct {
	App         appConfig
	environment string
}

// GetEnv returns the current developemnt environment
func (c Config) GetEnv() string {
	return c.environment
}

// Load reads all config from env to config
func Load() Config {
	return Config{
		environment: os.Getenv("APP_ENV"),
		App: appConfig{
			port: os.Getenv("APP_PORT"),
		},
	}
}
