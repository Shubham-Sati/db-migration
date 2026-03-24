package commands

import (
	"fmt"

	"chat-analytics-db-migration/configs"
	"chat-analytics-db-migration/database"

	"github.com/spf13/cobra"
)

// DropTables creates and returns the drop tables command
// This command removes all database tables (useful for clean reset)
func DropTables() *cobra.Command {
	return &cobra.Command{
		Use:   "droptables",
		Short: "Drop all database tables",
		Long:  "Remove all database tables for chat and analytics services (destructive operation)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if configs.App.Env != "local" {
				fmt.Println("Warning: Environment is not local. Tables wont be dropped")
				return nil
			}
			fmt.Println("App env is local")

			// Establish database connection
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Start a database transaction
			// This ensures all drops either succeed or fail together
			begin := dbConnection.Begin()

			// Get list of all tables to drop
			// DropAllTables function returns slice of drop operations
			dropOperations := database.DropAllTables(begin)

			fmt.Printf("[DropTables] Starting drop of %d tables...\n", len(dropOperations))

			// Execute each drop operation in sequence
			for i, dropOp := range dropOperations {
				// Run individual table drop
				if err := dropOp.Run(begin); err != nil {
					// If any drop fails, rollback the entire transaction
					begin.Rollback()
					fmt.Printf("[DropTables] Failed to drop table: %s, Error: %v\n", dropOp.TableName, err)
					return fmt.Errorf("drop failed at table %s: %w", dropOp.TableName, err)
				}

				// Log successful drop
				fmt.Printf("[%d]: Dropped table: %s\n", i+1, dropOp.TableName)
			}

			// Commit all successful drops
			if err := begin.Commit().Error; err != nil {
				fmt.Printf("[DropTables] Failed to commit table drops: %v\n", err)
				return fmt.Errorf("failed to commit table drops: %w", err)
			}

			fmt.Println("[DropTables] All tables dropped successfully!")
			return nil
		},
	}
}
