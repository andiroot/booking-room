package schedule

import (
	"time"
)

type Schedule struct {
	ScheduleID       int
	ScheduleName     string
	PIC              string
	StartAt          time.Time
	EndAt            time.Time
	TotalParticipant int
	Room             uint
}
