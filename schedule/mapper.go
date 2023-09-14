package schedule

func ConvertToScheduleResponse(s Schedule) ScheduleResponse {
	return ScheduleResponse{
		ScheduleID:       s.ScheduleID,
		ScheduleName:     s.ScheduleName,
		PIC:              s.PIC,
		StartAt:          s.StartAt,
		EndAt:            s.EndAt,
		TotalParticipant: s.TotalParticipant,
	}
}
