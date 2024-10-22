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
	PostgresConfig PostgresConfig `mapstructure:"postgres_config"` // URI for PostgreSQL database
	RedisConfig    RedisConfig `mapstructure:"redis_config"`    // URI for Redis database
}

// RedisConfig holds the configuration for Redis.
type RedisConfig struct {
	RedisUri             string `mapstructure:"redis_uri"`              // URI for Redis database
	MaxIdleConnections   int    `mapstructure:"max_idle_connections"`   // Max Idle Connections
	MaxActiveConnections int    `mapstructure:"max_active_connections"` // Max Active Connections
	IdleTimeout          int    `mapstructure:"idle_timeout"`           // Idle timeout in seconds
	Wait                 bool   `mapstructure:"wait"`                   // Wait for a connection to be available
}

// PostgresConfig holds the configuration for PostgreSQL.
type PostgresConfig struct {
	PostgresUri          string `mapstructure:"postgres_uri"`           // URI for PostgreSQL database
	MaxIdleConnections   int    `mapstructure:"max_idle_connections"`   // Max Idle Connections
	MaxOpenConnections   int    `mapstructure:"max_open_connections"`   // Max Open Connections
	ConnectionTimeout    int    `mapstructure:"connection_timeout"`     // Connection timeout in seconds
	IdleTimeout          int    `mapstructure:"idle_timeout"`           // Idle timeout in seconds
	ConnMaxLifetime      int    `mapstructure:"conn_max_lifetime"`      // Maximum amount of time a connection may be reused
	ConnMaxIdleTime      int    `mapstructure:"conn_max_idle_time"`     // Maximum amount of time a connection may be idle
	PreferSimpleProtocol bool   `mapstructure:"prefer_simple_protocol"` // Prefer simple protocol
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
	config.DBConfig.PostgresConfig.PostgresUri = os.Getenv(config.DBConfig.PostgresConfig.PostgresUri)
	config.DBConfig.RedisConfig.RedisUri = os.Getenv(config.DBConfig.RedisConfig.RedisUri)

	// Ensure the DB URIs are populated
	if config.DBConfig.PostgresConfig.PostgresUri == "" {
		return nil, fmt.Errorf("POSTGRES_URI is not set in the environment")
	}

	if config.DBConfig.RedisConfig.RedisUri == "" {
		return nil, fmt.Errorf("REDIS_URI is not set in the environment")
	}

	return &config, nil
}
