package service

import (
	"context"
	"monitor/internal/adaptor"
)

type Service struct {
	Alarm
}

func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		Alarm: NewAlarmService(adaptor),
	}
}

type Alarm interface {
	InsertAlarm(ctx context.Context) (int, error)
}
