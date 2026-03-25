package seeds

import (
	"chat-analytics-db-migration/constants"
	"chat-analytics-db-migration/tables/analytics"
	"database/sql"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func Events(db *gorm.DB) error {

	props1, _ := json.Marshal(map[string]interface{}{
		"page":     "/dashboard",
		"duration": 3200,
		"referrer": "direct",
	})

	props2, _ := json.Marshal(map[string]interface{}{
		"room_id":     "room_general",
		"message_count": 1,
		"characters":   45,
	})

	props3, _ := json.Marshal(map[string]interface{}{
		"feature":    "dark_mode",
		"enabled":    true,
		"section":    "preferences",
	})

	events := []analytics.Event{
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_DASHBOARD, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			SessionID:  sql.NullString{String: "sess_abc123def456", Valid: true},
			EventType:  sql.NullString{String: "page_view", Valid: true},
			EventName:  sql.NullString{String: "dashboard_visit", Valid: true},
			Properties: props1,
			Timestamp:  time.Now().Add(-4 * time.Hour),
			IPAddress:  sql.NullString{String: "192.168.1.100", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-4 * time.Hour),
			UpdatedAt:  time.Now().Add(-4 * time.Hour),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_MESSAGE_SENT, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			SessionID:  sql.NullString{String: "sess_xyz789ghi012", Valid: true},
			EventType:  sql.NullString{String: "user_action", Valid: true},
			EventName:  sql.NullString{String: "message_sent", Valid: true},
			Properties: props2,
			Timestamp:  time.Now().Add(-3 * time.Hour),
			IPAddress:  sql.NullString{String: "192.168.1.101", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-3 * time.Hour),
			UpdatedAt:  time.Now().Add(-3 * time.Hour),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_PREFERENCE, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			SessionID:  sql.NullString{String: "sess_mno345pqr678", Valid: true},
			EventType:  sql.NullString{String: "settings", Valid: true},
			EventName:  sql.NullString{String: "preference_change", Valid: true},
			Properties: props3,
			Timestamp:  time.Now().Add(-2 * time.Hour),
			IPAddress:  sql.NullString{String: "10.0.0.50", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X)", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-2 * time.Hour),
			UpdatedAt:  time.Now().Add(-2 * time.Hour),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_ROOM_JOIN, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_MARKETING, Valid: true},
			SessionID:  sql.NullString{String: "sess_stu901vwx234", Valid: true},
			EventType:  sql.NullString{String: "navigation", Valid: true},
			EventName:  sql.NullString{String: "room_join", Valid: true},
			Properties: json.RawMessage(`{"room_name": "Engineering Team", "room_type": "private"}`),
			Timestamp:  time.Now().Add(-90 * time.Minute),
			IPAddress:  sql.NullString{String: "172.16.0.25", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-90 * time.Minute),
			UpdatedAt:  time.Now().Add(-90 * time.Minute),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_FILE_UPLOAD, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_DESIGNER, Valid: true},
			SessionID:  sql.NullString{String: "sess_abc567def890", Valid: true},
			EventType:  sql.NullString{String: "file_action", Valid: true},
			EventName:  sql.NullString{String: "file_upload", Valid: true},
			Properties: json.RawMessage(`{"file_name": "report.pdf", "file_size": 1024000, "room": "support"}`),
			Timestamp:  time.Now().Add(-60 * time.Minute),
			IPAddress:  sql.NullString{String: "192.168.1.102", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Edge/96.0", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-60 * time.Minute),
			UpdatedAt:  time.Now().Add(-60 * time.Minute),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_LOGOUT, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			SessionID:  sql.NullString{String: "sess_ghi123jkl456", Valid: true},
			EventType:  sql.NullString{String: "user_action", Valid: true},
			EventName:  sql.NullString{String: "logout", Valid: true},
			Properties: json.RawMessage(`{"session_duration": 7200, "pages_visited": 12}`),
			Timestamp:  time.Now().Add(-30 * time.Minute),
			IPAddress:  sql.NullString{String: "192.168.1.100", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-30 * time.Minute),
			UpdatedAt:  time.Now().Add(-30 * time.Minute),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_LANDING, Valid: true},
			UserPID:    sql.NullString{String: "", Valid: false}, // Anonymous event
			SessionID:  sql.NullString{String: "sess_anon789xyz", Valid: true},
			EventType:  sql.NullString{String: "page_view", Valid: true},
			EventName:  sql.NullString{String: "landing_page", Valid: true},
			Properties: json.RawMessage(`{"page": "/", "referrer": "google.com", "campaign": "organic"}`),
			Timestamp:  time.Now().Add(-15 * time.Minute),
			IPAddress:  sql.NullString{String: "203.0.113.45", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Android; Mobile) AppleWebKit/537.36", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-15 * time.Minute),
			UpdatedAt:  time.Now().Add(-15 * time.Minute),
		},
		{
			PID:        sql.NullString{String: constants.SeedConstants.EVENT_API_ERROR, Valid: true},
			UserPID:    sql.NullString{String: constants.SeedConstants.USER_QA, Valid: true},
			SessionID:  sql.NullString{String: "sess_final123end", Valid: true},
			EventType:  sql.NullString{String: "error", Valid: true},
			EventName:  sql.NullString{String: "api_error", Valid: true},
			Properties: json.RawMessage(`{"error_code": "404", "endpoint": "/api/rooms/invalid", "retry_count": 3}`),
			Timestamp:  time.Now().Add(-10 * time.Minute),
			IPAddress:  sql.NullString{String: "192.168.1.103", Valid: true},
			UserAgent:  sql.NullString{String: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15", Valid: true},
			IsActive:   true,
			CreatedAt:  time.Now().Add(-10 * time.Minute),
			UpdatedAt:  time.Now().Add(-10 * time.Minute),
		},
	}

	return db.Create(&events).Error
}

