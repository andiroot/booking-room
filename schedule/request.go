package schedule

import "time"

type ScheduleRequest struct {
	ScheduleName     string    `binding:"required"`
	PIC              string    `binding:"required"`
	StartAt          time.Time `binding:"required"`
	EndAt            time.Time `binding:"required"`
	TotalParticipant int       `binding:"required"`
	Room             uint      `binding:"required"`
}
