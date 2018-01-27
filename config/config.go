package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr    string `envconfig:"BIND_ADDR"`
	APIURL      string `envconfig:"DATASET_API_URL"`
	SecretKey   string `envconfig:"SECRET_KEY"`
	MongoConfig MongoConfig
	RedisConfig RedisConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	Addr       string `envconfig:"MONGODB_BIND_ADDR"`
	Collection string `envconfig:"MONGODB_COLLECTION"`
	Database   string `envconfig:"MONGODB_MONITOR"`
}

// RedisConfig contains the config required to connect to MongoDB
type RedisConfig struct {
	Addr     string `envconfig:"REDIS_BIND_ADDR"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int8   `envconfig:"REDIS_DATABASE"`
}

var cfg *Configuration

// Get the application and returns the configuration structure
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:  ":22123",
		APIURL:    "http://localhost:22123",
		SecretKey: "FD0108EA-825D-411C-9B1D-41EF7727F465",
		MongoConfig: MongoConfig{
			Addr:       "localhost:27017",
			Collection: "monitor",
			Database:   "monitor",
		},
		RedisConfig: RedisConfig{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	sanitized := *cfg
	sanitized.SecretKey = ""

	return cfg, envconfig.Process("", cfg)
}
