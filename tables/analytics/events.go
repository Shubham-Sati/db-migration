package analytics

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Event struct {
	ID         int             `gorm:"column:events_id;primaryKey;autoIncrement"`
	PID        sql.NullString  `gorm:"column:events_pid;unique;not null;type:varchar(40)"`
	UserPID    sql.NullString  `gorm:"column:users_pid;type:varchar(40);default:null"`
	SessionID  sql.NullString  `gorm:"column:session_id;type:varchar(100);default:null"`
	EventType  sql.NullString  `gorm:"column:event_type;not null;type:varchar(50);index"`
	EventName  sql.NullString  `gorm:"column:event_name;not null;type:varchar(100)"`
	Properties json.RawMessage `gorm:"column:properties;type:jsonb;default:null"`
	Timestamp  time.Time       `gorm:"column:timestamp;not null;index"`
	IPAddress  sql.NullString  `gorm:"column:ip_address;type:varchar(45);default:null"`
	UserAgent  sql.NullString  `gorm:"column:user_agent;type:varchar(500);default:null"`
	IsActive   bool            `gorm:"column:is_active;default:true"`
	IsDeleted  bool            `gorm:"column:is_deleted;default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
