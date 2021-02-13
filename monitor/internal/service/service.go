package service

type Service struct {
	Alarm
}

func NewService() *Service {
	return &Service{
		Alarm: NewAlarmService(1),
	}
}

type Alarm interface {
	InsertAlarm() (int, error)
}
