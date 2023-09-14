package initializers

import (
	"booking-room/room"
	"booking-room/schedule"
	"booking-room/user"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	db.AutoMigrate(user.User{})
	db.AutoMigrate(schedule.Schedule{})
	db.AutoMigrate(room.Room{})
}