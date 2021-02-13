package service

import (
	"context"
	"monitor/internal/adaptor"
)

type AlarmService struct {
	adaptor *adaptor.Adaptor
}

// NewAlarmService is a constructor
func NewAlarmService(adaptor *adaptor.Adaptor) *AlarmService {
	return &AlarmService{
		adaptor: adaptor,
	}
}

func (s *AlarmService) InsertAlarm(ctx context.Context) (int, error) {
	s.adaptor.GetAllTechniques(ctx)
	return 1, nil
}
