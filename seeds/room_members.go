package seeds

import (
	"chat-analytics-db-migration/constants"
	"chat-analytics-db-migration/tables/chat"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func RoomMembers(db *gorm.DB) error {

	roomMembers := []chat.RoomMember{
		// General Discussion room members
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_1, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_ADMIN, Valid: true},
			Role:     sql.NullString{String: "admin", Valid: true},
			JoinedAt: time.Now().Add(-30 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-30 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_2, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			Role:     sql.NullString{String: "member", Valid: true},
			JoinedAt: time.Now().Add(-28 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-28 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_3, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_GENERAL, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			Role:     sql.NullString{String: "moderator", Valid: true},
			JoinedAt: time.Now().Add(-25 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-30 * time.Minute),
		},
		// Engineering Team room members
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_4, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_ENGINEERING, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_DEV, Valid: true},
			Role:     sql.NullString{String: "admin", Valid: true},
			JoinedAt: time.Now().Add(-25 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-25 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-45 * time.Minute),
		},
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_5, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_ENGINEERING, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_DESIGNER, Valid: true},
			Role:     sql.NullString{String: "member", Valid: true},
			JoinedAt: time.Now().Add(-20 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-20 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-1 * time.Hour),
		},
		// Random room members
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_6, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_RANDOM, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_MARKETING, Valid: true},
			Role:     sql.NullString{String: "member", Valid: true},
			JoinedAt: time.Now().Add(-18 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-18 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-2 * time.Hour),
		},
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_7, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_RANDOM, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_QA, Valid: true},
			Role:     sql.NullString{String: "member", Valid: true},
			JoinedAt: time.Now().Add(-15 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-15 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-3 * time.Hour),
		},
		// Support Help room member
		{
			PID:      sql.NullString{String: constants.SeedConstants.MEMBER_8, Valid: true},
			RoomPID:  sql.NullString{String: constants.SeedConstants.ROOM_SUPPORT, Valid: true},
			UserPID:  sql.NullString{String: constants.SeedConstants.USER_SUPPORT, Valid: true},
			Role:     sql.NullString{String: "admin", Valid: true},
			JoinedAt: time.Now().Add(-18 * 24 * time.Hour),
			IsActive: true,
			CreatedAt: time.Now().Add(-18 * 24 * time.Hour),
			UpdatedAt: time.Now().Add(-4 * time.Hour),
		},
	}

	return db.Create(&roomMembers).Error
}

