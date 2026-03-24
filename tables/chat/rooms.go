package chat

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Room represents a chat room where users can communicate
// Each room can have multiple users and messages
type Room struct {
	ID          int             `gorm:"column:chat_rooms_id;primaryKey;autoIncrement" json:"id"`
	PID         sql.NullString  `gorm:"column:chat_rooms_pid;unique;not null;type:varchar(40)" json:"pid"` // Public identifier
	Name        sql.NullString  `gorm:"column:room_name;not null;type:varchar(100)" json:"name"`           // Room display name
	Description sql.NullString  `gorm:"column:room_description;type:varchar(500)" json:"description"`      // Room description/topic
	OwnerPID    sql.NullString  `gorm:"column:owner_pid;not null;type:varchar(40);index" json:"owner_pid"` // Room creator/admin PID
	RoomType    sql.NullString  `gorm:"column:room_type;type:varchar(20);default:public" json:"room_type"` // public, private, direct
	MaxUsers    int             `gorm:"column:max_users;default:100" json:"max_users"`                     // Maximum users allowed
	Settings    json.RawMessage `gorm:"column:room_settings;type:jsonb;default:null" json:"settings"`      // Room configuration
	IsActive    bool            `gorm:"column:is_active;default:true" json:"is_active"`                    // Room status
	IsDeleted   bool            `gorm:"column:is_deleted;default:false" json:"is_deleted"`                 // Soft delete flag
	CreatedAt   time.Time       `json:"created_at"`                                                        // Room creation time
	UpdatedAt   time.Time       `json:"updated_at"`                                                        // Last room update
}
