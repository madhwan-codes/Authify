// internal/db/redis.go

package db

import (
	"github.com/madhwan-codes/authify/internal/config"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

// initRedis initializes the Redis connection.
func initRedis(cfg config.RedisConfig) (*redis.Client, error) {
	options := &redis.Options{
		Addr:         cfg.RedisUri,
		PoolSize:     cfg.MaxActiveConnections,
		MinIdleConns: cfg.MaxIdleConnections,
	}

	if cfg.Wait {
		options.MaxRetries = 3
	}

	client := redis.NewClient(options)

	// Optionally, you can ping the Redis server to check if the connection is established
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
