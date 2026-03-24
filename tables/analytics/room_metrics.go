package analytics

import (
	"database/sql"
	"time"
)

type RoomMetric struct {
	ID             int             `gorm:"column:room_metrics_id;primaryKey;autoIncrement"`
	PID            sql.NullString  `gorm:"column:room_metrics_pid;unique;not null;type:varchar(40)"`
	RoomPID        sql.NullString  `gorm:"column:rooms_pid;not null;type:varchar(40);index"`
	Date           time.Time       `gorm:"column:date;not null;index;type:date"`
	MessageCount   int             `gorm:"column:message_count;default:0"`
	UniqueUsers    int             `gorm:"column:unique_users;default:0"`
	PeakConcurrent int             `gorm:"column:peak_concurrent;default:0"`
	AvgMessageLen  sql.NullFloat64 `gorm:"column:avg_message_length;default:null"`
	TotalReactions int             `gorm:"column:total_reactions;default:0"`
	IsActive       bool            `gorm:"column:is_active;default:true"`
	IsDeleted      bool            `gorm:"column:is_deleted;default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
