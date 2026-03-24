package chat

import (
	"database/sql"
	"encoding/json"
	"time"
)

// Message represents a single chat message
// Contains the actual message content and metadata
type Message struct {
	ID          int             `gorm:"column:chat_messages_id;primaryKey;autoIncrement" json:"id"`
	PID         sql.NullString  `gorm:"column:chat_messages_pid;unique;not null;type:varchar(40)" json:"pid"`  // Public identifier
	RoomPID     sql.NullString  `gorm:"column:room_pid;not null;type:varchar(40);index" json:"room_pid"`       // Room PID reference
	UserPID     sql.NullString  `gorm:"column:user_pid;not null;type:varchar(40);index" json:"user_pid"`       // User PID reference
	Content     sql.NullString  `gorm:"column:message_content;not null;type:text" json:"content"`              // Message text content
	MessageType sql.NullString  `gorm:"column:message_type;type:varchar(20);default:text" json:"message_type"` // text, image, file, system
	ReplyToPID  sql.NullString  `gorm:"column:reply_to_pid;type:varchar(40);index" json:"reply_to_pid"`        // PID of message being replied to
	Metadata    json.RawMessage `gorm:"column:metadata;type:jsonb;default:null" json:"metadata"`               // Additional message data
	IsEdited    bool            `gorm:"column:is_edited;default:false" json:"is_edited"`                       // Whether message was edited
	IsActive    bool            `gorm:"column:is_active;default:true" json:"is_active"`                        // Message status
	IsDeleted   bool            `gorm:"column:is_deleted;default:false" json:"is_deleted"`                     // Soft delete flag
	CreatedAt   time.Time       `json:"created_at"`                                                            // When message was sent
	UpdatedAt   time.Time       `json:"updated_at"`                                                            // Last edit time
}
