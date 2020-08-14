package config

import (
	"os"
	"strconv"
)

// Config contains all the necessary configurations
type Config struct {
	App         appConfig
	environment string
	DBConfig    DBConfig
}

// GetEnv returns the current development environment
func (c Config) GetEnv() string {
	return c.environment
}

// Load reads all config from env to config
func Load() Config {
	timeout, _ := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if timeout == 0 {
		timeout = 5
	}
	return Config{
		environment: os.Getenv("APP_ENV"),
		App: appConfig{
			port: os.Getenv("APP_PORT"),
		},
		DBConfig: DBConfig{
			user:       os.Getenv("DB_USER"),
			pass:       os.Getenv("DB_PASS"),
			host:       os.Getenv("DB_HOST"),
			port:       os.Getenv("DB_PORT"),
			database:   os.Getenv("DB_NAME"),
			collection: os.Getenv("DB_COLLECTION"),
			env:        os.Getenv("APP_ENV"),
			timeout:    timeout,
		},
	}
}
