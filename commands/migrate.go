package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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

			// Now run SQL migrations automatically
			fmt.Println("\n[Migrate] Running SQL migrations...")

			// Create migration history table if it doesn't exist
			if err := database.CreateMigrationHistoryTable(dbConnection); err != nil {
				fmt.Printf("[Migrate] Warning: Could not create migration history table: %v\n", err)
				// Continue anyway as this is not critical for base migration
			}

			// Get applied migrations
			applied, err := database.GetAppliedMigrations(dbConnection)
			if err != nil {
				fmt.Printf("[Migrate] Warning: Could not get applied migrations: %v\n", err)
				applied = make(map[string]bool)
			}

			// Load all migration files
			migrationsPath := filepath.Join(".", "migrations")
			sqlMigrations, err := database.LoadMigrationFiles(migrationsPath)
			if err != nil {
				fmt.Printf("[Migrate] Info: No SQL migrations found (this is okay): %v\n", err)
				// This is okay if migrations folder doesn't exist or has no valid migrations
				return nil
			}

			// Find pending migrations
			var pending []database.MigrationFile
			for _, m := range sqlMigrations {
				if !applied[m.Version] {
					pending = append(pending, m)
				}
			}

			if len(pending) == 0 {
				fmt.Println("[Migrate] No pending SQL migrations to apply")
				return nil
			}

			fmt.Printf("[Migrate] Found %d pending SQL migrations\n", len(pending))

			// Apply each migration
			for _, migration := range pending {
				fmt.Printf("[Migrate] Applying SQL migration %s: %s\n", migration.Version, migration.Description)

				startTime := time.Now()
				tx := dbConnection.Begin()

				// Read and execute UP migration
				upPath := filepath.Join(migrationsPath, migration.Version+"_"+migration.Description, "up.sql")
				upSQL, err := os.ReadFile(upPath)
				if err != nil {
					tx.Rollback()
					fmt.Printf("[Migrate] Failed to read migration file: %v\n", err)
					return fmt.Errorf("failed to read migration %s: %w", migration.Version, err)
				}

				if err := tx.Exec(string(upSQL)).Error; err != nil {
					tx.Rollback()
					fmt.Printf("[Migrate] Failed to execute migration: %v\n", err)
					return fmt.Errorf("failed to execute migration %s: %w", migration.Version, err)
				}

				// Record successful migration
				history := &database.MigrationHistory{
					Version:       migration.Version,
					Description:   migration.Description,
					AppliedAt:     time.Now(),
					ExecutionTime: time.Since(startTime).Milliseconds(),
				}

				if err := tx.Create(history).Error; err != nil {
					tx.Rollback()
					fmt.Printf("[Migrate] Failed to record migration: %v\n", err)
					return fmt.Errorf("failed to record migration %s: %w", migration.Version, err)
				}

				tx.Commit()
				fmt.Printf("[Migrate] ✓ Applied %s successfully (took %dms)\n", migration.Version, history.ExecutionTime)
			}

			fmt.Println("\n[Migrate] All base tables and SQL migrations completed successfully!")
			return nil
		},
	}
}

