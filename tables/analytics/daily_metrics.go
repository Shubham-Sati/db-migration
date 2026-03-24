package analytics

import (
	"database/sql"
	"time"
)

type DailyMetric struct {
	ID             int             `gorm:"column:daily_metrics_id;primaryKey;autoIncrement"`
	PID            sql.NullString  `gorm:"column:daily_metrics_pid;unique;not null;type:varchar(40)"`
	Date           time.Time       `gorm:"column:date;not null;uniqueIndex:idx_date_metric;type:date"`
	MetricType     sql.NullString  `gorm:"column:metric_type;not null;uniqueIndex:idx_date_metric;type:varchar(50)"`
	TotalUsers     int             `gorm:"column:total_users;default:0"`
	NewUsers       int             `gorm:"column:new_users;default:0"`
	TotalMessages  int             `gorm:"column:total_messages;default:0"`
	TotalRooms     int             `gorm:"column:total_rooms;default:0"`
	ActiveRooms    int             `gorm:"column:active_rooms;default:0"`
	AvgSessionTime sql.NullFloat64 `gorm:"column:avg_session_time;default:null"`
	TotalSessions  int             `gorm:"column:total_sessions;default:0"`
	IsActive       bool            `gorm:"column:is_active;default:true"`
	IsDeleted      bool            `gorm:"column:is_deleted;default:false"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
