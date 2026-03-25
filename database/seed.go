package database

import (
	"chat-analytics-db-migration/seeds"
	"fmt"

	"gorm.io/gorm"
)

// SeedAllTables returns a list of seed operations for populating initial data
func SeedAllTables(db *gorm.DB) []Migration {
	fmt.Println("[Debug] SeedAllTables function called")
	seedOperations := []Migration{
		{
			TableName: "users_seed",
			Model:     nil,
			Run: func(tx *gorm.DB) error {
				return seeds.Users(tx)
			},
		},
		{
			TableName: "rooms_seed",
			Model:     nil,
			Run: func(tx *gorm.DB) error {
				return seeds.Rooms(tx)
			},
		},
		{
			TableName: "room_members_seed",
			Model:     nil,
			Run: func(tx *gorm.DB) error {
				return seeds.RoomMembers(tx)
			},
		},
		{
			TableName: "messages_seed",
			Model:     nil,
			Run: func(tx *gorm.DB) error {
				return seeds.Messages(tx)
			},
		},
		{
			TableName: "events_seed",
			Model:     nil,
			Run: func(tx *gorm.DB) error {
				return seeds.Events(tx)
			},
		},
	}

	fmt.Printf("[Debug] SeedAllTables returning %d operations\n", len(seedOperations))
	return seedOperations
}
