package seeds

import (
	"chat-analytics-db-migration/constants"
	"chat-analytics-db-migration/tables/shared"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	metadata1, _ := json.Marshal(map[string]interface{}{
		"role":        "admin",
		"department":  "engineering",
		"timezone":    "UTC",
		"preferences": map[string]bool{"notifications": true, "dark_mode": false},
	})

	metadata2, _ := json.Marshal(map[string]interface{}{
		"role":        "user",
		"department":  "marketing",
		"timezone":    "PST",
		"preferences": map[string]bool{"notifications": true, "dark_mode": true},
	})

	metadata3, _ := json.Marshal(map[string]interface{}{
		"role":        "moderator",
		"department":  "support",
		"timezone":    "EST",
		"preferences": map[string]bool{"notifications": false, "dark_mode": false},
	})

	users := []shared.User{
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			Username:  sql.NullString{String: "john_admin", Valid: true},
			Email:     sql.NullString{String: "john@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true}, // hashed "password123"
			FullName:  sql.NullString{String: "John Anderson", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/john.jpg", Valid: true},
			Metadata:  metadata1,
			LastSeen:  sql.NullTime{Time: time.Now().Add(-30 * time.Minute), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-5 * time.Minute),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			Username:  sql.NullString{String: "sarah_dev", Valid: true},
			Email:     sql.NullString{String: "sarah@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Sarah Wilson", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/sarah.jpg", Valid: true},
			Metadata:  metadata2,
			LastSeen:  sql.NullTime{Time: time.Now().Add(-10 * time.Minute), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Minute),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			Username:  sql.NullString{String: "mike_support", Valid: true},
			Email:     sql.NullString{String: "mike@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Mike Thompson", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/mike.jpg", Valid: true},
			Metadata:  metadata3,
			LastSeen:  sql.NullTime{Time: time.Now().Add(-2 * time.Hour), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-20 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_MARKETING, Valid: true},
			Username:  sql.NullString{String: "emma_marketing", Valid: true},
			Email:     sql.NullString{String: "emma@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Emma Rodriguez", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/emma.jpg", Valid: true},
			Metadata:  json.RawMessage(`{"role": "user", "department": "marketing"}`),
			LastSeen:  sql.NullTime{Time: time.Now().Add(-1 * time.Hour), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-15 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-30 * time.Minute),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_DESIGNER, Valid: true},
			Username:  sql.NullString{String: "alex_designer", Valid: true},
			Email:     sql.NullString{String: "alex@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Alex Chen", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/alex.jpg", Valid: true},
			Metadata:  json.RawMessage(`{"role": "user", "department": "design", "timezone": "JST"}`),
			LastSeen:  sql.NullTime{Time: time.Now().Add(-45 * time.Minute), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-12 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-20 * time.Minute),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_QA, Valid: true},
			Username:  sql.NullString{String: "lisa_qa", Valid: true},
			Email:     sql.NullString{String: "lisa@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Lisa Johnson", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/lisa.jpg", Valid: true},
			Metadata:  json.RawMessage(`{"role": "user", "department": "qa"}`),
			LastSeen:  sql.NullTime{Time: time.Now().Add(-3 * time.Hour), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-8 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_INACTIVE, Valid: true},
			Username:  sql.NullString{String: "david_inactive", Valid: true},
			Email:     sql.NullString{String: "david@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "David Miller", Valid: true},
			Avatar:    sql.NullString{String: "https://avatar.example.com/david.jpg", Valid: true},
			Metadata:  json.RawMessage(`{"role": "user", "department": "sales", "status": "inactive"}`),
			LastSeen:  sql.NullTime{Time: time.Now().Add(-7 * 24 * time.Hour), Valid: true},
			IsActive:  false,
			CreatedAt: time.Now().Add(-45 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-7 * 24 * time.Hour),
		},
		{
			PID:       sql.NullString{String: constants.SeedConstants.USER_GUEST, Valid: true},
			Username:  sql.NullString{String: "guest_user", Valid: true},
			Email:     sql.NullString{String: "guest@company.com", Valid: true},
			Password:  sql.NullString{String: "$2a$10$HgtaXOgAzWCteVOB.8nuMuLIs/xjb.p5fm3Kk2GZL7vytaIJjcnlW", Valid: true},
			FullName:  sql.NullString{String: "Guest User", Valid: true},
			Avatar:    sql.NullString{String: "", Valid: false},
			Metadata:  json.RawMessage(`{"role": "guest", "temporary": true}`),
			LastSeen:  sql.NullTime{Time: time.Now().Add(-15 * time.Minute), Valid: true},
			IsActive:  true,
			CreatedAt: time.Now().Add(-1 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-15 * time.Minute),
		},
	}

	return db.Create(&users).Error
}
