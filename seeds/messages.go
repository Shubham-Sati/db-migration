package seeds

import (
	"chat-analytics-db-migration/constants"
	"chat-analytics-db-migration/tables/chat"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func Messages(db *gorm.DB) error {

	metadata1, _ := json.Marshal(map[string]interface{}{
		"edited": false,
		"edited_at": nil,
		"client": "web",
		"platform": "chrome",
	})

	metadata2, _ := json.Marshal(map[string]interface{}{
		"edited": true,
		"edited_at": time.Now().Add(-2 * time.Hour).Format(time.RFC3339),
		"original_content": "Welcome everyone to the team!",
		"client": "mobile",
		"platform": "ios",
	})

	messages := []chat.Message{
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_WELCOME, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			Content:     sql.NullString{String: "Welcome everyone to the General Discussion room! Feel free to share ideas and collaborate.", Valid: true},
			MessageType: sql.NullString{String: "text", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    metadata2,
			IsEdited:    true,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-2 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_THANKS, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			Content:     sql.NullString{String: "Thanks! Excited to be part of the team 🎉", Valid: true},
			MessageType: sql.NullString{String: "text", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    metadata1,
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-29 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-29 * 24 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_ENGINEERING, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_ENGINEERING, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			Content:     sql.NullString{String: "Hey team, let's discuss the new feature requirements for next sprint.", Valid: true},
			MessageType: sql.NullString{String: "text", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    json.RawMessage(`{"client": "desktop", "urgent": true}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-25 * 24 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_TECH_SPECS, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_ENGINEERING, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_DESIGNER, Valid: true},
			Content:     sql.NullString{String: "I've uploaded the technical specs document. Please review when you have time.", Valid: true},
			MessageType: sql.NullString{String: "file", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    json.RawMessage(`{"file_name": "tech_specs_v2.pdf", "file_size": "2.3MB", "file_type": "pdf"}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-24 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-24 * 24 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_COFFEE, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_RANDOM, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_MARKETING, Valid: true},
			Content:     sql.NullString{String: "Anyone want to grab coffee? ☕", Valid: true},
			MessageType: sql.NullString{String: "text", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    json.RawMessage(`{"emoji_count": 1, "reaction_enabled": true}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-3 * time.Hour),
			UpdatedAt:   time.Now().Add(-3 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_COFFEE_REPLY, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_RANDOM, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_QA, Valid: true},
			Content:     sql.NullString{String: "Count me in! I'll be down in 5 minutes.", Valid: true},
			MessageType: sql.NullString{String: "text", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false}, // This would be the PID of the coffee message in real scenario
			Metadata:    json.RawMessage(`{"reply": true, "quick_response": true}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-2*time.Hour - 50*time.Minute),
			UpdatedAt:   time.Now().Add(-2*time.Hour - 50*time.Minute),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_SUPPORT, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_SUPPORT, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			Content:     sql.NullString{String: "Hi everyone! If you need help with any technical issues, please describe your problem here.", Valid: true},
			MessageType: sql.NullString{String: "system", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    json.RawMessage(`{"system_message": true, "pinned": true, "auto_generated": false}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-18 * 24 * time.Hour),
			UpdatedAt:   time.Now().Add(-18 * 24 * time.Hour),
		},
		{
			PID:         sql.NullString{String: constants.SeedConstants.MESSAGE_ANNOUNCEMENT, Valid: true},
			RoomPID:     sql.NullString{String: constants.SeedConstants.ROOM_PRODUCT, Valid: true},
			UserPID:     sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			Content:     sql.NullString{String: "🚀 New feature release v2.1.0 is now live! Check out the improved dashboard and analytics.", Valid: true},
			MessageType: sql.NullString{String: "announcement", Valid: true},
			ReplyToPID:  sql.NullString{Valid: false},
			Metadata:    json.RawMessage(`{"announcement": true, "version": "2.1.0", "priority": "high"}`),
			IsEdited:    false,
			IsActive:    true,
			CreatedAt:   time.Now().Add(-1 * time.Hour),
			UpdatedAt:   time.Now().Add(-1 * time.Hour),
		},
	}

	return db.Create(&messages).Error
}

