package database

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

// MigrationHistory tracks which migrations have been applied
type MigrationHistory struct {
	Version       string    `gorm:"primaryKey;type:varchar(255)"`
	Description   string    `gorm:"type:text"`
	AppliedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ExecutionTime int64     `gorm:"type:bigint"` // milliseconds
	Success       bool      `gorm:"default:true"`
	RollbackSQL   string    `gorm:"type:text"` // Store DOWN migration for easy rollback
}

// MigrationFile represents a migration file
type MigrationFile struct {
	Version     string
	Filename    string
	UpSQL       string
	DownSQL     string
	Description string
}

// CreateMigrationHistoryTable ensures the migration_history table exists
func CreateMigrationHistoryTable(db *gorm.DB) error {
	// Create table if it doesn't exist
	if err := db.AutoMigrate(&MigrationHistory{}); err != nil {
		return fmt.Errorf("failed to create migration_history table: %w", err)
	}

	fmt.Println("[Migration] Migration history table ready")
	return nil
}

// GetAppliedMigrations returns list of already applied migration versions
func GetAppliedMigrations(db *gorm.DB) (map[string]bool, error) {
	var migrations []MigrationHistory
	if err := db.Where("success = ?", true).Find(&migrations).Error; err != nil {
		return nil, fmt.Errorf("failed to get applied migrations: %w", err)
	}

	applied := make(map[string]bool)
	for _, m := range migrations {
		applied[m.Version] = true
	}

	return applied, nil
}

// LoadMigrationFiles reads all migration files from the migrations directory
func LoadMigrationFiles(migrationsPath string) ([]MigrationFile, error) {
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var migrations []MigrationFile

	for _, file := range files {
		// Skip directories and non-SQL files
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		// Parse filename: YYYYMMDD_HHMMSS_description.sql
		parts := strings.SplitN(file.Name(), "_", 3)
		if len(parts) < 3 {
			fmt.Printf("[Migration] Skipping invalid filename format: %s\n", file.Name())
			continue
		}

		version := parts[0] + "_" + parts[1] // YYYYMMDD_HHMMSS
		description := strings.TrimSuffix(parts[2], ".sql")

		// Read file content
		content, err := ioutil.ReadFile(filepath.Join(migrationsPath, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
		}

		// Split UP and DOWN migrations
		upSQL, downSQL := parseMigrationContent(string(content))

		migrations = append(migrations, MigrationFile{
			Version:     version,
			Filename:    file.Name(),
			UpSQL:       upSQL,
			DownSQL:     downSQL,
			Description: description,
		})
	}

	// Sort migrations by version (timestamp)
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

// parseMigrationContent splits migration file into UP and DOWN sections
func parseMigrationContent(content string) (up string, down string) {
	lines := strings.Split(content, "\n")
	var currentSection string
	var upLines, downLines []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(strings.ToUpper(line))

		// Check for section markers
		if strings.HasPrefix(trimmed, "-- UP") || strings.HasPrefix(trimmed, "--UP") {
			currentSection = "UP"
			continue
		}
		if strings.HasPrefix(trimmed, "-- DOWN") || strings.HasPrefix(trimmed, "--DOWN") {
			currentSection = "DOWN"
			continue
		}

		// Add line to appropriate section
		switch currentSection {
		case "UP":
			upLines = append(upLines, line)
		case "DOWN":
			downLines = append(downLines, line)
		}
	}

	up = strings.TrimSpace(strings.Join(upLines, "\n"))
	down = strings.TrimSpace(strings.Join(downLines, "\n"))

	// If no DOWN section found, try to auto-generate for simple cases
	if down == "" && up != "" {
		down = generateRollback(up)
	}

	return up, down
}

// generateRollback attempts to generate a simple rollback for common operations
func generateRollback(upSQL string) string {
	upper := strings.ToUpper(strings.TrimSpace(upSQL))

	// Simple ALTER TABLE ADD COLUMN -> DROP COLUMN
	if strings.Contains(upper, "ALTER TABLE") && strings.Contains(upper, "ADD COLUMN") {
		// This is a simplified example - real implementation would need proper SQL parsing
		return "-- Rollback not auto-generated. Please define manually in -- DOWN section"
	}

	return "-- Rollback not defined. Please add -- DOWN section to migration file"
}

// RecordMigration records a successful migration in the history table
func RecordMigration(db *gorm.DB, migration MigrationFile, executionTime int64) error {
	record := MigrationHistory{
		Version:       migration.Version,
		Description:   migration.Description,
		AppliedAt:     time.Now(),
		ExecutionTime: executionTime,
		Success:       true,
		RollbackSQL:   migration.DownSQL,
	}

	if err := db.Create(&record).Error; err != nil {
		return fmt.Errorf("failed to record migration %s: %w", migration.Version, err)
	}

	return nil
}

// GetLastMigration returns the most recently applied migration
func GetLastMigration(db *gorm.DB) (*MigrationHistory, error) {
	var migration MigrationHistory
	err := db.Where("success = ?", true).
		Order("applied_at DESC").
		First(&migration).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get last migration: %w", err)
	}

	return &migration, nil
}

// RemoveMigrationRecord removes a migration from the history (used during rollback)
func RemoveMigrationRecord(db *gorm.DB, version string) error {
	if err := db.Where("version = ?", version).Delete(&MigrationHistory{}).Error; err != nil {
		return fmt.Errorf("failed to remove migration record %s: %w", version, err)
	}
	return nil
}
