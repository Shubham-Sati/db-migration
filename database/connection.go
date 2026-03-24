package database

import (
	"chat-analytics-db-migration/configs"
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection establishes a connection to PostgreSQL database
// Returns both GORM and SQL database instances for different use cases
func Connection() (*gorm.DB, *sql.DB) {
	// Build PostgreSQL connection string (DSN - Data Source Name)
	// Using configuration values from the configs package
	dsn := "host=" + configs.DB.Host +
		" user=" + configs.DB.Username +
		" password=" + configs.DB.Password +
		" dbname=" + configs.DB.Database +
		" port=" + configs.DB.Port +
		" sslmode=" + configs.DB.SSLMode

	// Log connection attempt (hide password for security)
	log.Printf("[Connection] Connecting to database: %s@%s:%s/%s",
		configs.DB.Username, configs.DB.Host, configs.DB.Port, configs.DB.Database)

	// Open GORM connection with PostgreSQL driver
	// GORM provides ORM functionality for easier database operations
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false, // Ensure foreign keys work properly
	})
	if err != nil {
		log.Fatalf("[Connection] Error opening GORM connection: %v", err)
	}

	// Get underlying SQL database instance
	// This is needed for connection pooling and low-level operations
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("[Connection] Error getting SQL DB instance: %v", err)
	}

	// Configure connection pooling for optimal performance
	// These settings prevent connection exhaustion and improve performance
	sqlDB.SetMaxOpenConns(configs.DB.MaxOpenConns) // Maximum connections to database
	sqlDB.SetMaxIdleConns(configs.DB.MaxIdleConns) // Maximum idle connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute)     // Maximum connection lifetime (prevent stale connections)

	log.Printf("[Connection] Database connected successfully with %d max connections", configs.DB.MaxOpenConns)

	return db, sqlDB
}
