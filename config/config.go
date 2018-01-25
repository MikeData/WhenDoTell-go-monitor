package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

// Configuration structure which hold information for configuring the import API
type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	APIURL           				string        `envconfig:"DATASET_API_URL"`
	SecretKey               string        `envconfig:"SECRET_KEY"`
	MongoConfig             MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr   string `envconfig:"MONGODB_BIND_ADDR"`
	Collection string `envconfig:"MONGODB_COLLECTION"`
	Database   string `envconfig:"MONGODB_MONITOR"`
}

var cfg *Configuration

// Get the application and returns the configuration structure
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                          ":22123",
		APIURL:                            "http://localhost:22123",
		SecretKey:                         "FD0108EA-825D-411C-9B1D-41EF7727F465",
		MongoConfig: MongoConfig{
			BindAddr:                        "localhost:27017",
			Collection:                      "monitor",
			Database:                        "monitor",
		},
	}

	sanitized := *cfg
	sanitized.SecretKey = ""
	log.Printf("config on startup, config: %s", sanitized)

	return cfg, envconfig.Process("", cfg)
}
