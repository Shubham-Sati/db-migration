package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig holds general application configuration
// These settings control how the migration tool behaves
type AppConfig struct {
	Env      string `envconfig:"ENVIRONMENT" default:"development"` // development, staging, production
	Port     string `envconfig:"PORT" default:"8080"`               // Port for test server
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`          // Logging verbosity
	Debug    bool   `envconfig:"DEBUG" default:"false"`             // Enable debug mode
}

// Global variable to access app config throughout the application
var App AppConfig

// loadAppConfig reads application settings from environment variables
// Uses envconfig to automatically map environment variables to struct fields
func loadAppConfig() {
	err := envconfig.Process("", &App)
	if err != nil {
		log.Fatalf("[AppConfig] Error loading app configuration: %v", err)
	}

	log.Printf("[AppConfig] Loaded configuration - Environment: %s, Port: %s", App.Env, App.Port)
}
