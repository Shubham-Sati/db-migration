package analytics

import (
	"database/sql"
	"time"
)

type UserMetric struct {
	ID                int            `gorm:"column:user_metrics_id;primaryKey;autoIncrement"`
	PID               sql.NullString `gorm:"column:user_metrics_pid;unique;not null;type:varchar(40)"`
	UserPID           sql.NullString `gorm:"column:users_pid;not null;type:varchar(40);index"`
	Date              time.Time      `gorm:"column:date;not null;index;type:date"`
	MessagesSent      int            `gorm:"column:messages_sent;default:0"`
	RoomsJoined       int            `gorm:"column:rooms_joined;default:0"`
	SessionDuration   int            `gorm:"column:session_duration;default:0"`
	ReactionsGiven    int            `gorm:"column:reactions_given;default:0"`
	ReactionsReceived int            `gorm:"column:reactions_received;default:0"`
	FirstActivity     sql.NullTime   `gorm:"column:first_activity;default:null"`
	LastActivity      sql.NullTime   `gorm:"column:last_activity;default:null"`
	IsActive          bool           `gorm:"column:is_active;default:true"`
	IsDeleted         bool           `gorm:"column:is_deleted;default:false"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
