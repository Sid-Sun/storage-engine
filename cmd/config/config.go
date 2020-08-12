package config

import (
	"os"
)

// Config contains all the neccessary configurations
type Config struct {
	App         appConfig
	environment string
	DBConfig    DBConfig
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
		DBConfig: DBConfig{
			user:       os.Getenv("DB_USER"),
			pass:       os.Getenv("DB_PASS"),
			host:       os.Getenv("DB_HOST"),
			port:       os.Getenv("DB_PORT"),
			database:   os.Getenv("DB_NAME"),
			collection: os.Getenv("DB_COLLECTION"),
			env:        os.Getenv("APP_ENV"),
		},
	}
}
