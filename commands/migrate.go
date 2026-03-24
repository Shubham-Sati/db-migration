package commands

import (
	"fmt"

	"chat-analytics-db-migration/database"

	"github.com/spf13/cobra"
)

// Migrate creates and returns the migrate command
// This command runs all database migrations to create tables and schema
func Migrate() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Long:  "Create all database tables and schema for chat and analytics services",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Establish database connection
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Start a database transaction
			// This ensures all migrations either succeed or fail together
			begin := dbConnection.Begin()

			// Get list of all migrations to run
			// AutoMigrate function returns slice of migration operations
			migrations := database.AutoMigrate(begin)
			
			fmt.Printf("[Migrate] Starting migration of %d tables...\n", len(migrations))

			// Execute each migration in sequence
			for i, migrate := range migrations {
				// Run individual migration
				if err := migrate.Run(begin); err != nil {
					// If any migration fails, rollback the entire transaction
					begin.Rollback()
					fmt.Printf("[Migrate] Failed at table: %s, Error: %v\n", migrate.TableName, err)
					return fmt.Errorf("migration failed at table %s: %w", migrate.TableName, err)
				}
				
				// Log successful migration
				fmt.Printf("[%d]: Migrated table: %s\n", i+1, migrate.TableName)
			}

			// Commit all successful migrations
			if err := begin.Commit().Error; err != nil {
				fmt.Printf("[Migrate] Failed to commit migrations: %v\n", err)
				return fmt.Errorf("failed to commit migrations: %w", err)
			}

			fmt.Println("[Migrate] All migrations completed successfully!")
			return nil
		},
	}
}