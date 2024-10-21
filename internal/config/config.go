package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

// AppConfig holds the overall configuration for the application.
type AppConfig struct {
	Environment   string        `mapstructure:"environment"`     // Environment (e.g., dev, stag, prod)
	ServiceConfig ServiceConfig `mapstructure:"service_config"`  // Configuration for the service
	DBConfig      DBConfig      `mapstructure:"database_config"` // Configuration for the database
	LoggingConfig LoggingConfig `mapstructure:"logging_config"`  // Configuration for logging
}

// ServiceConfig holds the configuration for the service.
type ServiceConfig struct {
	Name    string `mapstructure:"name"`    // Name of the service
	Port    int    `mapstructure:"port"`    // Port on which the service runs
	Timeout int    `mapstructure:"timeout"` // Request timeout in seconds
	Retries int    `mapstructure:"retries"` // Number of retries on failure
}

// DBConfig holds the configuration for the database.
type DBConfig struct {
	PostgresUri string `mapstructure:"postgres_uri"` // URI for PostgreSQL database
	RedisUri    string `mapstructure:"redis_uri"`    // URI for Redis database
}

// LoggingConfig holds the configuration for logging.
type LoggingConfig struct {
	Level string `mapstructure:"level"` // Logging level (e.g., info, debug)
}

// LoadConfig loads the configuration from config.yaml and .env file
func LoadConfig() (*AppConfig, error) {
	var config AppConfig

	// Load config.yaml using Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // Path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Load environment variables from .env file (for dev purposes)
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Get the actual values from the environment using keys from the config.yaml
	config.DBConfig.PostgresUri = os.Getenv(config.DBConfig.PostgresUri)
	config.DBConfig.RedisUri = os.Getenv(config.DBConfig.RedisUri)

	// Ensure the DB URIs are populated
	if config.DBConfig.PostgresUri == "" {
		return nil, fmt.Errorf("POSTGRES_URI is not set in the environment")
	}

	if config.DBConfig.PostgresUri == "" {
		return nil, fmt.Errorf("REDIS_URI is not set in the environment")
	}

	return &config, nil
}
