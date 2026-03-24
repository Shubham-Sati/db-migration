package database

import (
	"chat-analytics-db-migration/tables/analytics"
	"chat-analytics-db-migration/tables/chat"
	"chat-analytics-db-migration/tables/shared"

	"gorm.io/gorm"
)

// Migration represents a single database migration operation
// Contains the table model and logic to execute the migration
type Migration struct {
	TableName string               // Name of the table being migrated
	Model     interface{}          // GORM model struct
	Run       func(*gorm.DB) error // Function to execute the migration
}

// AutoMigrate returns a list of all migrations to be executed
// Migrations are ordered to respect foreign key dependencies
func AutoMigrate(db *gorm.DB) []Migration {
	migrations := []Migration{
		// Step 1: Create shared tables (no dependencies)
		{
			TableName: "users",
			Model:     &shared.User{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&shared.User{})
			},
		},

		// Step 2: Create chat tables (depend on users)
		{
			TableName: "chat_rooms",
			Model:     &chat.Room{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&chat.Room{})
			},
		},
		{
			TableName: "chat_room_members",
			Model:     &chat.RoomMember{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&chat.RoomMember{})
			},
		},
		{
			TableName: "chat_messages",
			Model:     &chat.Message{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&chat.Message{})
			},
		},
		{
			TableName: "chat_message_reactions",
			Model:     &chat.MessageReaction{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&chat.MessageReaction{})
			},
		},
		{
			TableName: "chat_message_attachments",
			Model:     &chat.MessageAttachment{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&chat.MessageAttachment{})
			},
		},

		// Step 3: Create analytics tables (depend on users and chat tables)
		{
			TableName: "analytics_events",
			Model:     &analytics.Event{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&analytics.Event{})
			},
		},
		{
			TableName: "analytics_user_sessions",
			Model:     &analytics.UserSession{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&analytics.UserSession{})
			},
		},
		{
			TableName: "analytics_daily_metrics",
			Model:     &analytics.DailyMetric{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&analytics.DailyMetric{})
			},
		},
		{
			TableName: "analytics_room_metrics",
			Model:     &analytics.RoomMetric{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&analytics.RoomMetric{})
			},
		},
		{
			TableName: "analytics_user_metrics",
			Model:     &analytics.UserMetric{},
			Run: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&analytics.UserMetric{})
			},
		},
	}

	return migrations
}
