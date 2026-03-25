package seeds

import (
	"chat-analytics-db-migration/constants"
	"chat-analytics-db-migration/tables/chat"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func Rooms(db *gorm.DB) error {

	settings1, _ := json.Marshal(map[string]interface{}{
		"allow_file_upload": true,
		"max_file_size":     "10MB",
		"moderation":        true,
		"auto_archive":      false,
	})

	settings2, _ := json.Marshal(map[string]interface{}{
		"allow_file_upload": false,
		"moderation":        false,
		"auto_archive":      true,
		"archive_after":     "30d",
	})

	rooms := []chat.Room{
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			Name:        sql.NullString{String: "General Discussion", Valid: true},
			Description: sql.NullString{String: "Main chat room for general company discussions", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			RoomType:    sql.NullString{String: "public", Valid: true},
			MaxUsers:    100,
			Settings:    settings1,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_ENGINEERING, Valid: true},
			Name:        sql.NullString{String: "Engineering Team", Valid: true},
			Description: sql.NullString{String: "Private room for engineering team discussions", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			RoomType:    sql.NullString{String: "private", Valid: true},
			MaxUsers:    25,
			Settings:    json.RawMessage(`{"allow_file_upload": true, "code_snippets": true}`),
			IsActive:    true,
			CreatedAt:   time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-30 * time.Minute),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_RANDOM, Valid: true},
			Name:        sql.NullString{String: "Random", Valid: true},
			Description: sql.NullString{String: "Casual conversations and off-topic discussions", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			RoomType:    sql.NullString{String: "public", Valid: true},
			MaxUsers:    50,
			Settings:    settings2,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-20 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_SUPPORT, Valid: true},
			Name:        sql.NullString{String: "Support Help", Valid: true},
			Description: sql.NullString{String: "Customer support and help desk discussions", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			RoomType:    sql.NullString{String: "public", Valid: true},
			MaxUsers:    75,
			Settings:    json.RawMessage(`{"priority_support": true, "ticket_integration": true}`),
			IsActive:    true,
			CreatedAt:   time.Now().Add(-18 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-45 * time.Minute),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_PRODUCT, Valid: true},
			Name:        sql.NullString{String: "Product Updates", Valid: true},
			Description: sql.NullString{String: "Announcements and updates about product releases", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			RoomType:    sql.NullString{String: "public", Valid: true},
			MaxUsers:    200,
			Settings:    json.RawMessage(`{"read_only": false, "announcements": true}`),
			IsActive:    true,
			CreatedAt:   time.Now().Add(-15 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-3 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_MARKETING, Valid: true},
			Name:        sql.NullString{String: "Marketing Team", Valid: true},
			Description: sql.NullString{String: "Marketing team coordination and campaigns", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_MARKETING, Valid: true},
			RoomType:    sql.NullString{String: "private", Valid: true},
			MaxUsers:    15,
			Settings:    json.RawMessage(`{"campaign_tools": true, "analytics": true}`),
			IsActive:    true,
			CreatedAt:   time.Now().Add(-12 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-4 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_DESIGN, Valid: true},
			Name:        sql.NullString{String: "Design Critique", Valid: true},
			Description: sql.NullString{String: "Design reviews, feedback, and creative discussions", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_DESIGNER, Valid: true},
			RoomType:    sql.NullString{String: "private", Valid: true},
			MaxUsers:    20,
			Settings:    json.RawMessage(`{"file_upload": true, "image_preview": true, "creative_tools": true}`),
			IsActive:    true,
			CreatedAt:   time.Now().Add(-10 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-6 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.ROOM_ARCHIVE, Valid: true},
			Name:        sql.NullString{String: "Archive Test Room", Valid: true},
			Description: sql.NullString{String: "Test room for archive functionality - inactive", Valid: true},
			OwnerPID:    sql.NullString{String: constants.SeedConstants.USER_INACTIVE, Valid: true},
			RoomType:    sql.NullString{String: "public", Valid: true},
			MaxUsers:    10,
			Settings:    json.RawMessage(`{"archived": true}`),
			IsActive:    false,
			CreatedAt:   time.Now().Add(-60 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-30 * 24 * time.Hour),
		},
	}

	return db.Create(&rooms).Error
}
