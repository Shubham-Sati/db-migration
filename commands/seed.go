package commands

import (
	"fmt"

	"chat-analytics-db-migration/database"

	"github.com/spf13/cobra"
)

// Seed creates and returns the seed command
// This command populates database tables with initial/test data
func Seed() *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "Populate database with seed data",
		Long:  "Insert initial/test data into chat and analytics service tables",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Establish database connection
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Start a database transaction
			// This ensures all seeds either succeed or fail together
			begin := dbConnection.Begin()

			// Get list of all seed operations
			// SeedAllTables function returns slice of seeding operations
			seedOperations := database.SeedAllTables(begin)

			fmt.Printf("[Seed] Starting seeding of %d table groups...\n", len(seedOperations))

			// Execute each seed operation in sequence
			for i, seedOp := range seedOperations {
				// Run individual seed operation
				if err := seedOp.Run(begin); err != nil {
					// If any seed fails, rollback the entire transaction
					begin.Rollback()
					fmt.Printf("[Seed] Failed to seed: %s, Error: %v\n", seedOp.TableName, err)
					return fmt.Errorf("seeding failed for %s: %w", seedOp.TableName, err)
				}

				// Log successful seed
				fmt.Printf("[%d]: Seeded: %s\n", i+1, seedOp.TableName)
			}

			// Commit all successful seeds
			if err := begin.Commit().Error; err != nil {
				fmt.Printf("[Seed] Failed to commit seed data: %v\n", err)
				return fmt.Errorf("failed to commit seed data: %w", err)
			}

			fmt.Println("[Seed] All seed data inserted successfully!")
			return nil
		},
	}
}
