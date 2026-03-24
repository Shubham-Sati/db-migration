package analytics

import (
	"database/sql"
	"time"
)

type UserSession struct {
	ID        int            `gorm:"column:user_sessions_id;primaryKey;autoIncrement"`
	PID       sql.NullString `gorm:"column:user_sessions_pid;unique;not null;type:varchar(40)"`
	SessionID sql.NullString `gorm:"column:session_id;unique;not null;type:varchar(100)"`
	UserPID   sql.NullString `gorm:"column:users_pid;type:varchar(40);default:null"`
	StartTime time.Time      `gorm:"column:start_time;not null"`
	EndTime   sql.NullTime   `gorm:"column:end_time;default:null"`
	Duration  sql.NullInt64  `gorm:"column:duration;default:null"`
	PageViews int            `gorm:"column:page_views;default:0"`
	Events    int            `gorm:"column:events;default:0"`
	IPAddress sql.NullString `gorm:"column:ip_address;type:varchar(45);default:null"`
	UserAgent sql.NullString `gorm:"column:user_agent;type:varchar(500);default:null"`
	IsActive  bool           `gorm:"column:is_active;default:true"`
	IsDeleted bool           `gorm:"column:is_deleted;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
