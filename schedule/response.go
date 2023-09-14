package schedule

import "time"

type ScheduleResponse struct {
	ScheduleID       int       `json:"schedule_id"`
	ScheduleName     string    `json:"schedule_name"`
	PIC              string    `json:"pic"`
	StartAt          time.Time `json:"start_at"`
	EndAt            time.Time `json:"end_at"`
	TotalParticipant int       `json:"total_participant"`
	Room             uint      `json:"room"`
}
