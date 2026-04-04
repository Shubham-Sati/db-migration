package database

import (
	"chat-analytics-db-migration/tables/analytics"
	"chat-analytics-db-migration/tables/chat"
	"chat-analytics-db-migration/tables/shared"

	"gorm.io/gorm"
)

// DropAllTables returns a list of operations to drop all tables
// Order is reversed to respect foreign key constraints
func DropAllTables(db *gorm.DB) []Migration {
	// Reverse order of creation to handle dependencies
	dropOperations := []Migration{
		// Drop migration history table first (no dependencies)
		{
			TableName: "migration_histories",
			Run: func(tx *gorm.DB) error {
				// Use DROP TABLE IF EXISTS to avoid errors if table doesn't exist
				return tx.Exec("DROP TABLE IF EXISTS migration_histories CASCADE").Error
			},
		},
		
		// Drop analytics tables (they depend on others)
		{
			TableName: "analytics_user_metrics",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&analytics.UserMetric{})
			},
		},
		{
			TableName: "analytics_room_metrics",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&analytics.RoomMetric{})
			},
		},
		{
			TableName: "analytics_daily_metrics",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&analytics.DailyMetric{})
			},
		},
		{
			TableName: "analytics_user_sessions",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&analytics.UserSession{})
			},
		},
		{
			TableName: "analytics_events",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&analytics.Event{})
			},
		},

		// Drop chat tables (they depend on users)
		{
			TableName: "chat_message_attachments",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&chat.MessageAttachment{})
			},
		},
		{
			TableName: "chat_message_reactions",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&chat.MessageReaction{})
			},
		},
		{
			TableName: "chat_messages",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&chat.Message{})
			},
		},
		{
			TableName: "chat_room_members",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&chat.RoomMember{})
			},
		},
		{
			TableName: "chat_rooms",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&chat.Room{})
			},
		},

		// Drop shared tables last (others depend on them)
		{
			TableName: "users",
			Run: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&shared.User{})
			},
		},
	}

	return dropOperations
}
