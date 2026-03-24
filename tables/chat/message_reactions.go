package chat

import (
	"database/sql"
	"time"
)

// MessageReaction represents reactions to messages (like, love, laugh, etc.)
// Allows users to react to messages without sending new messages
type MessageReaction struct {
	ID         int            `gorm:"column:chat_message_reactions_id;primaryKey;autoIncrement" json:"id"`
	PID        sql.NullString `gorm:"column:chat_message_reactions_pid;unique;not null;type:varchar(40)" json:"pid"` // Public identifier
	MessagePID sql.NullString `gorm:"column:message_pid;not null;type:varchar(40);index" json:"message_pid"`         // Message PID reference
	UserPID    sql.NullString `gorm:"column:user_pid;not null;type:varchar(40);index" json:"user_pid"`               // User PID reference
	Reaction   sql.NullString `gorm:"column:reaction_type;not null;type:varchar(20)" json:"reaction"`                // like, love, laugh, angry, sad
	IsActive   bool           `gorm:"column:is_active;default:true" json:"is_active"`                                // Reaction status
	IsDeleted  bool           `gorm:"column:is_deleted;default:false" json:"is_deleted"`                             // Soft delete flag
	CreatedAt  time.Time      `json:"created_at"`                                                                    // When reaction was added
	UpdatedAt  time.Time      `json:"updated_at"`                                                                    // Last update time
}
