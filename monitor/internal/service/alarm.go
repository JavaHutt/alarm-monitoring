package service

type AlarmService struct {
	repo int
}

func NewAlarmService(repo int) *AlarmService {
	return &AlarmService{}
}

func (s *AlarmService) InsertAlarm() (int, error) {
	return 1, nil
}
