package main

import (
	"chat-analytics-db-migration/commands"
	"chat-analytics-db-migration/configs"

	"github.com/spf13/cobra"
)

func main() {
	// Load environment variables and configuration settings
	// This reads from .env file and initializes database connection configs
	configs.LoadConfigs()

	// Create the root CLI command with basic information
	cmd := &cobra.Command{
		Use:   "db-migration",
		Short: "Database migration tool for chat and analytics services",
		Long:  "A CLI tool to manage database migrations, seeding, and schema management for chat and analytics services",
	}

	// Register all available CLI commands
	// Each command handles a specific database operation
	cmd.AddCommand(commands.DropTables())  // Drop all database tables
	cmd.AddCommand(commands.Migrate())     // Run database migrations (old - creates tables)
	cmd.AddCommand(commands.MigrateNew())  // New migration system with up/down/status
	cmd.AddCommand(commands.Seed())        // Populate tables with seed data
	cmd.AddCommand(commands.Alter())       // Alter existing table structures (deprecated - use migration)
	cmd.AddCommand(commands.RunServer())   // Start a test server (optional)

	// Execute the CLI command based on user input
	// If command fails, panic to stop execution
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
