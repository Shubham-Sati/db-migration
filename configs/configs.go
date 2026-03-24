package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// LoadConfigs initializes all configuration settings
// This function is called once at application startup
func LoadConfigs() {
	// Get the current file path to locate .env file relative to this config file
	_, currentFile, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(currentFile), "../")

	// Log the root path to verify we're looking in the correct location
	log.Printf("[LoadConfigs] currentFile path determined as: %s", currentFile)
	log.Printf("[LoadConfigs] Root path determined as: %s", rootPath)
	log.Printf("[LoadConfigs] Looking for .env file at: %s", rootPath+"/.env")

	// Load environment variables from .env file
	// .env file contains sensitive data like database credentials
	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize application configuration
	// This loads general app settings like port, environment, etc.
	loadAppConfig()

	// Initialize database configuration
	// This loads database connection parameters
	loadDBConfig()
}
