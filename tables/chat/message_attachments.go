package chat

import (
	"database/sql"
	"time"
)

// MessageAttachment represents files/images attached to messages
// Handles file uploads and media sharing in chat
type MessageAttachment struct {
	ID         int            `gorm:"column:chat_message_attachments_id;primaryKey;autoIncrement" json:"id"`
	PID        sql.NullString `gorm:"column:chat_message_attachments_pid;unique;not null;type:varchar(40)" json:"pid"` // Public identifier
	MessagePID sql.NullString `gorm:"column:message_pid;not null;type:varchar(40);index" json:"message_pid"`           // Message PID reference
	FileName   sql.NullString `gorm:"column:file_name;not null;type:varchar(255)" json:"file_name"`                    // Original filename
	FileURL    sql.NullString `gorm:"column:file_url;not null;type:varchar(500)" json:"file_url"`                      // URL to access the file
	FileType   sql.NullString `gorm:"column:file_type;not null;type:varchar(50)" json:"file_type"`                     // MIME type (image/png, etc.)
	FileSize   int64          `gorm:"column:file_size;not null" json:"file_size"`                                      // File size in bytes
	IsActive   bool           `gorm:"column:is_active;default:true" json:"is_active"`                                  // Attachment status
	IsDeleted  bool           `gorm:"column:is_deleted;default:false" json:"is_deleted"`                               // Soft delete flag
	CreatedAt  time.Time      `json:"created_at"`                                                                      // When file was uploaded
	UpdatedAt  time.Time      `json:"updated_at"`                                                                      // Last update time
}
