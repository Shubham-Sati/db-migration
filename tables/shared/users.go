package shared

import (
	"database/sql"
	"encoding/json"
	"time"
)

// User represents a user in the system
// This table is shared between chat and analytics services
type User struct {
	ID        int             `gorm:"column:user_id;primaryKey;autoIncrement" json:"id"`
	PID       sql.NullString  `gorm:"column:user_pid;unique;not null;type:varchar(40)" json:"pid"`      // Public identifier
	Username  sql.NullString  `gorm:"column:username;unique;not null;type:varchar(50)" json:"username"` // Unique username for login
	Email     sql.NullString  `gorm:"column:email;unique;not null;type:varchar(255)" json:"email"`      // Unique email address
	Password  sql.NullString  `gorm:"column:password_hash;not null;type:varchar(255)" json:"-"`         // Hashed password (hidden from JSON)
	FullName  sql.NullString  `gorm:"column:full_name;type:varchar(100)" json:"full_name"`              // Display name
	Avatar    sql.NullString  `gorm:"column:avatar_url;type:varchar(500)" json:"avatar"`                // Profile picture URL
	Metadata  json.RawMessage `gorm:"column:metadata;type:jsonb;default:null" json:"metadata"`          // Additional user data
	LastSeen  sql.NullTime    `gorm:"column:last_seen_at" json:"last_seen"`                             // Last activity timestamp
	IsActive  bool            `gorm:"column:is_active;default:true" json:"is_active"`                   // Account status
	IsDeleted bool            `gorm:"column:is_deleted;default:false" json:"is_deleted"`                // Soft delete flag
	CreatedAt time.Time       `json:"created_at"`                                                       // Account creation time
	UpdatedAt time.Time       `json:"updated_at"`                                                       // Last profile update
}
