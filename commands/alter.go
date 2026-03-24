package commands

import (
	"fmt"

	"chat-analytics-db-migration/database"

	"github.com/spf13/cobra"
)

// Alter creates and returns the alter command
// This command modifies existing table structures (add columns, indexes, etc.)
func Alter() *cobra.Command {
	return &cobra.Command{
		Use:   "alter",
		Short: "Alter existing database tables",
		Long:  "Modify existing table structures by adding columns, indexes, or constraints",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Establish database connection
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Start a database transaction
			// This ensures all alterations either succeed or fail together
			begin := dbConnection.Begin()

			// Get list of all alter operations
			// AlterAllTables function returns slice of alteration operations
			alterOperations := database.AlterAllTables(begin)

			fmt.Printf("[Alter] Starting alteration of %d table operations...\n", len(alterOperations))

			// Execute each alter operation in sequence
			for i, alterOp := range alterOperations {
				// Run individual alter operation
				if err := alterOp.Run(begin); err != nil {
					// If any alter fails, rollback the entire transaction
					begin.Rollback()
					fmt.Printf("[Alter] Failed to alter: %s, Error: %v\n", alterOp.TableName, err)
					return fmt.Errorf("alteration failed for %s: %w", alterOp.TableName, err)
				}

				// Log successful alter
				fmt.Printf("[%d]: Altered: %s\n", i+1, alterOp.TableName)
			}

			// Commit all successful alterations
			if err := begin.Commit().Error; err != nil {
				fmt.Printf("[Alter] Failed to commit table alterations: %v\n", err)
				return fmt.Errorf("failed to commit table alterations: %w", err)
			}

			fmt.Println("[Alter] All table alterations completed successfully!")
			return nil
		},
	}
}
