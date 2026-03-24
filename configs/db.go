package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// DBConfig holds database connection configuration
// These settings are used to connect to PostgreSQL database
type DBConfig struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`    // Database host address
	Port     string `envconfig:"DB_PORT" default:"5432"`         // Database port
	Username string `envconfig:"DB_USERNAME" default:"postgres"` // Database username
	Password string `envconfig:"DB_PASSWORD" default:"password"` // Database password
	Database string `envconfig:"DB_NAME" default:"chatapp"`      // Database name
	SSLMode  string `envconfig:"DB_SSLMODE" default:"disable"`   // SSL mode (disable, require, etc.)

	// Connection pooling settings
	MaxOpenConns int `envconfig:"DB_MAX_OPEN_CONNS" default:"100"` // Maximum open connections
	MaxIdleConns int `envconfig:"DB_MAX_IDLE_CONNS" default:"10"`  // Maximum idle connections
}

// Global variable to access database config throughout the application
var DB DBConfig

// loadDBConfig reads database settings from environment variables
// Uses envconfig to automatically map environment variables to struct fields
func loadDBConfig() {
	err := envconfig.Process("", &DB)
	if err != nil {
		log.Fatalf("[DBConfig] Error loading database configuration: %v", err)
	}

	// Log database config (but hide password for security)
	log.Printf("[DBConfig] Loaded DB configuration - Host: %s:%s, Database: %s, User: %s",
		DB.Host, DB.Port, DB.Database, DB.Username)
}
