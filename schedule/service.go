package schedule

type Service interface {
	FindAll(UserID uint) ([]Schedule, error)
	FindByID(ID int) (Schedule, error)
	Create(scheduleRequest ScheduleRequest, UserID uint) (Schedule, error)
	Update(ID int, scheduleRequest ScheduleRequest) (Schedule, error)
	Delete(ID int) (Schedule, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]Schedule, error) {
	sch, err := s.repository.FindAll()
	return sch, err
}
func (s *service) FindByID(ID int) (Schedule, error) {
	sch, err := s.repository.FindByID(ID)
	return sch, err
}
func (s *service) Create(scheduleRequest ScheduleRequest, UserID uint) (Schedule, error) {
	sch := Schedule{
		ScheduleName:     scheduleRequest.ScheduleName,
		PIC:              scheduleRequest.PIC,
		StartAt:          scheduleRequest.StartAt,
		EndAt:            scheduleRequest.EndAt,
		TotalParticipant: scheduleRequest.TotalParticipant,
		Room:             scheduleRequest.Room,
	}
	newsch, err := s.repository.Create(sch)
	return newsch, err
}
func (s *service) Update(ID int, scheduleRequest ScheduleRequest) (Schedule, error) {
	sch, err := s.repository.FindByID(ID)

	if scheduleRequest.ScheduleName != "" {
		sch.ScheduleName = scheduleRequest.ScheduleName
	}

	if scheduleRequest.PIC != "" {
		sch.PIC = scheduleRequest.PIC
	}

	if scheduleRequest.StartAt.String() != "" {
		sch.StartAt = scheduleRequest.StartAt
	}

	if scheduleRequest.EndAt.String() != "" {
		sch.EndAt = scheduleRequest.EndAt
	}

	if scheduleRequest.TotalParticipant != 0 {
		sch.TotalParticipant = scheduleRequest.TotalParticipant
	}

	if scheduleRequest.TotalParticipant != 0 {
		sch.TotalParticipant = scheduleRequest.TotalParticipant
	}

	newsch, err := s.repository.Update(sch)
	return newsch, err
}

func (s *service) Delete(ID int) (Schedule, error) {
	room, err := s.repository.FindByID(ID)
	_, err = s.repository.Delete(room)
	return room, err
}
