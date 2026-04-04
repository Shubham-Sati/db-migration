package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"chat-analytics-db-migration/database"

	"github.com/spf13/cobra"
)

// MigrateNew creates the new migration command with up/down/status subcommands
func MigrateNew() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migration",
		Short: "Manage database migrations",
		Long:  "Run, rollback, and check status of database migrations",
	}

	// Add subcommands
	cmd.AddCommand(migrateUp())
	cmd.AddCommand(migrateDown())
	cmd.AddCommand(migrateStatus())
	cmd.AddCommand(migrateCreate())

	return cmd
}

// migrateUp runs all pending migrations
func migrateUp() *cobra.Command {
	return &cobra.Command{
		Use:   "up",
		Short: "Apply all pending migrations",
		Long:  "Runs all migrations that haven't been applied yet, in chronological order",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Connect to database
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Create migration history table if it doesn't exist
			if err := database.CreateMigrationHistoryTable(dbConnection); err != nil {
				return err
			}

			// Get applied migrations
			applied, err := database.GetAppliedMigrations(dbConnection)
			if err != nil {
				return err
			}

			// Load all migration files
			migrationsPath := filepath.Join(".", "migrations")
			migrations, err := database.LoadMigrationFiles(migrationsPath)
			if err != nil {
				return err
			}

			// Find pending migrations
			var pending []database.MigrationFile
			for _, m := range migrations {
				if !applied[m.Version] {
					pending = append(pending, m)
				}
			}

			if len(pending) == 0 {
				fmt.Println("[Migration] No pending migrations to apply")
				return nil
			}

			fmt.Printf("[Migration] Found %d pending migrations\n", len(pending))

			// Apply each migration in a transaction
			for _, migration := range pending {
				fmt.Printf("[Migration] Applying %s: %s\n", migration.Version, migration.Description)

				startTime := time.Now()

				// Start transaction
				tx := dbConnection.Begin()

				// Execute UP migration
				if err := tx.Exec(migration.UpSQL).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to apply migration %s: %w", migration.Version, err)
				}

				// Record in history
				executionTime := time.Since(startTime).Milliseconds()
				if err := database.RecordMigration(tx, migration, executionTime); err != nil {
					tx.Rollback()
					return err
				}

				// Commit transaction
				if err := tx.Commit().Error; err != nil {
					return fmt.Errorf("failed to commit migration %s: %w", migration.Version, err)
				}

				fmt.Printf("[Migration] ✅ Applied %s (%dms)\n", migration.Version, executionTime)
			}

			fmt.Printf("[Migration] Successfully applied %d migrations\n", len(pending))
			return nil
		},
	}
}

// migrateDown rolls back the last migration
func migrateDown() *cobra.Command {
	return &cobra.Command{
		Use:   "down",
		Short: "Rollback the last applied migration",
		Long:  "Rolls back the most recently applied migration using its DOWN script",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Connect to database
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Get last migration
			lastMigration, err := database.GetLastMigration(dbConnection)
			if err != nil {
				return err
			}

			if lastMigration == nil {
				fmt.Println("[Migration] No migrations to rollback")
				return nil
			}

			fmt.Printf("[Migration] Rolling back %s: %s\n", lastMigration.Version, lastMigration.Description)

			// Start transaction
			tx := dbConnection.Begin()

			// Execute DOWN migration
			if lastMigration.RollbackSQL != "" &&
				!contains(lastMigration.RollbackSQL, "Rollback not") {
				if err := tx.Exec(lastMigration.RollbackSQL).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to rollback migration %s: %w", lastMigration.Version, err)
				}
			} else {
				tx.Rollback()
				return fmt.Errorf("no rollback SQL defined for migration %s", lastMigration.Version)
			}

			// Remove from history
			if err := database.RemoveMigrationRecord(tx, lastMigration.Version); err != nil {
				tx.Rollback()
				return err
			}

			// Commit transaction
			if err := tx.Commit().Error; err != nil {
				return fmt.Errorf("failed to commit rollback for %s: %w", lastMigration.Version, err)
			}

			fmt.Printf("[Migration] ✅ Rolled back %s\n", lastMigration.Version)
			return nil
		},
	}
}

// migrateStatus shows the status of all migrations
func migrateStatus() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show migration status",
		Long:  "Lists all migrations and shows which have been applied",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Connect to database
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			// Create migration history table if it doesn't exist
			if err := database.CreateMigrationHistoryTable(dbConnection); err != nil {
				return err
			}

			// Get applied migrations
			applied, err := database.GetAppliedMigrations(dbConnection)
			if err != nil {
				return err
			}

			// Load all migration files
			migrationsPath := filepath.Join(".", "migrations")
			migrations, err := database.LoadMigrationFiles(migrationsPath)
			if err != nil {
				return err
			}

			fmt.Println("\n[Migration Status]")
			fmt.Println("=====================================")

			if len(migrations) == 0 {
				fmt.Println("No migration files found")
				return nil
			}

			for _, m := range migrations {
				status := "⏳ Pending"
				if applied[m.Version] {
					status = "✅ Applied"
				}
				fmt.Printf("%s | %s | %s\n", status, m.Version, m.Description)
			}

			fmt.Printf("\nTotal: %d migrations (%d applied, %d pending)\n",
				len(migrations), len(applied), len(migrations)-len(applied))

			return nil
		},
	}
}

// migrateCreate creates a new migration file
func migrateCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "create [description]",
		Short: "Create a new migration file",
		Long:  "Creates a new timestamped migration folder with separate up.sql and down.sql files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			description := args[0]

			// Generate timestamp
			timestamp := time.Now().Format("20060102_150405")

			// Create folder name
			folderName := fmt.Sprintf("%s_%s", timestamp, description)
			folderPath := filepath.Join(".", "migrations", folderName)

			// Create migration folder
			if err := os.MkdirAll(folderPath, 0755); err != nil {
				return fmt.Errorf("failed to create migration folder: %w", err)
			}

			// Template for up.sql
			upTemplate := `-- Migration: %s
-- Description: %s
-- Created: %s
-- Direction: UP

-- Add your forward migration SQL statements here
-- Example:
-- ALTER TABLE users ADD COLUMN new_field VARCHAR(255);
-- CREATE INDEX idx_new_field ON users(new_field);
`

			// Template for down.sql
			downTemplate := `-- Migration: %s
-- Description: %s
-- Created: %s
-- Direction: DOWN

-- Add your rollback SQL statements here
-- These should undo what the UP migration does
-- Example:
-- DROP INDEX IF EXISTS idx_new_field;
-- ALTER TABLE users DROP COLUMN IF EXISTS new_field;
`

			createdAt := time.Now().Format("2006-01-02 15:04:05")
			migrationName := timestamp + "_" + description

			// Create up.sql file
			upContent := fmt.Sprintf(upTemplate, migrationName, description, createdAt)
			upPath := filepath.Join(folderPath, "up.sql")
			if err := writeFile(upPath, upContent); err != nil {
				return fmt.Errorf("failed to create up.sql file: %w", err)
			}

			// Create down.sql file
			downContent := fmt.Sprintf(downTemplate, migrationName, description, createdAt)
			downPath := filepath.Join(folderPath, "down.sql")
			if err := writeFile(downPath, downContent); err != nil {
				return fmt.Errorf("failed to create down.sql file: %w", err)
			}

			fmt.Printf("[Migration] Created new migration: %s\n", folderName)
			fmt.Printf("[Migration] Files created:\n")
			fmt.Printf("  - %s (apply migration)\n", upPath)
			fmt.Printf("  - %s (rollback migration)\n", downPath)
			fmt.Printf("[Migration] Edit these files to add your SQL statements\n")

			return nil
		},
	}
}

// Helper function to check if string contains substring
func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (s == substr || (len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || (len(substr) < len(s) && findSubstring(s, substr)))))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Helper to write file
func writeFile(filepath, content string) error {
	// Create file with proper permissions
	return os.WriteFile(filepath, []byte(content), 0644)
}
