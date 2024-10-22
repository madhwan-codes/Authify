// internal/db/postgres.go

package db

import (
	config "github.com/madhwan-codes/authify/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// initPostgres initializes the PostgreSQL connection pool using GORM.
func initPostgres(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := cfg.PostgresUri // Data Source Name
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	// Optionally, you can ping the database to check if the connection is established
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
