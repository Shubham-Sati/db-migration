package chat

import (
	"database/sql"
	"time"
)

// RoomMember represents the many-to-many relationship between users and rooms
// Tracks which users are members of which rooms with additional metadata
type RoomMember struct {
	ID        int            `gorm:"column:chat_room_members_id;primaryKey;autoIncrement" json:"id"`
	PID       sql.NullString `gorm:"column:chat_room_members_pid;unique;not null;type:varchar(40)" json:"pid"` // Public identifier
	RoomPID   sql.NullString `gorm:"column:room_pid;not null;type:varchar(40);index" json:"room_pid"`          // Room PID reference
	UserPID   sql.NullString `gorm:"column:user_pid;not null;type:varchar(40);index" json:"user_pid"`          // User PID reference
	Role      sql.NullString `gorm:"column:member_role;type:varchar(20);default:member" json:"role"`           // member, moderator, admin
	JoinedAt  time.Time      `gorm:"column:joined_at;autoCreateTime" json:"joined_at"`                         // When user joined room
	IsActive  bool           `gorm:"column:is_active;default:true" json:"is_active"`                           // Active membership status
	IsDeleted bool           `gorm:"column:is_deleted;default:false" json:"is_deleted"`                        // Soft delete flag
	CreatedAt time.Time      `json:"created_at"`                                                               // Record creation time
	UpdatedAt time.Time      `json:"updated_at"`                                                               // Last update time
}
