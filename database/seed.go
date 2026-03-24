package database

import (
	"gorm.io/gorm"
)

// SeedAllTables returns a list of seed operations for populating initial data
// Currently returns empty slice - implement seed data as needed
func SeedAllTables(db *gorm.DB) []Migration {
	seedOperations := []Migration{
		// Example: Insert seed users
		// {
		// 	TableName: "users_seed",
		// 	Run: func(tx *gorm.DB) error {
		// 		seedUsers := []shared.User{
		// 			{
		// 				PID:      sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.USER), Valid: true},
		// 				Username: sql.NullString{String: "admin", Valid: true},
		// 				Email:    sql.NullString{String: "admin@example.com", Valid: true},
		// 				FullName: sql.NullString{String: "System Administrator", Valid: true},
		// 			},
		// 			{
		// 				PID:      sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.USER), Valid: true},
		// 				Username: sql.NullString{String: "user1", Valid: true},
		// 				Email:    sql.NullString{String: "user1@example.com", Valid: true},
		// 				FullName: sql.NullString{String: "Test User 1", Valid: true},
		// 			},
		// 		}
		// 		return tx.Create(&seedUsers).Error
		// 	},
		// },

		// Example: Insert seed rooms
		// {
		// 	TableName: "rooms_seed",
		// 	Run: func(tx *gorm.DB) error {
		// 		seedRooms := []chat.Room{
		// 			{
		// 				PID:      sql.NullString{String: utils.UUIDWithPrefix(constants.Prefix.ROOM), Valid: true},
		// 				Name:     sql.NullString{String: "General", Valid: true},
		// 				Description: sql.NullString{String: "General discussion room", Valid: true},
		// 				RoomType: sql.NullString{String: "public", Valid: true},
		// 			},
		// 		}
		// 		return tx.Create(&seedRooms).Error
		// 	},
		// },
	}

	return seedOperations
}
