package service

import (
	"context"
	"monitor/internal/adaptor"
	"monitor/internal/model"
)

// Service is a service
type Service struct {
	Alarm
}

// NewService is a constructor
func NewService(adaptor *adaptor.Adaptor) *Service {
	return &Service{
		Alarm: NewAlarmService(adaptor),
	}
}

// Alarm is interface for handling incoming alarms
type Alarm interface {
	InsertAlarm(ctx context.Context, alarm model.Alarm) error
}
