// internal/db/db.go

package db

import (
	"fmt"
	"github.com/madhwan-codes/authify/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// DB holds the database connection pools.
type DB struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}

// NewDB initializes the database connections.
func NewDB(cfg config.DBConfig) (*DB, error) {
	// Initialize PostgreSQL connection
	postgresDB, err := initPostgres(cfg.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf("could not initialize PostgreSQL: %w", err)
	}

	// Initialize Redis connection
	redisClient, err := initRedis(cfg.RedisConfig)
	if err != nil {
		return nil, fmt.Errorf("could not initialize Redis: %w", err)
	}

	return &DB{
		Postgres: postgresDB,
		Redis:    redisClient,
	}, nil
}
