package service

import (
	"context"
	"monitor/internal/adaptor"
	"monitor/internal/model"
)

// AlarmService service for Alarms
type AlarmService struct {
	adaptor *adaptor.Adaptor
}

// NewAlarmService is a constructor
func NewAlarmService(adaptor *adaptor.Adaptor) *AlarmService {
	return &AlarmService{
		adaptor: adaptor,
	}
}

// InsertAlarm creates or updates alarm depending on Alarm state
func (s *AlarmService) InsertAlarm(ctx context.Context, alarm model.Alarm) error {
	existOngoing, _ := s.adaptor.GetOngoingTechByComponentName(ctx, "Abc", "Dce")
	// if alarm.Crit > model.Ok {
	// 	existOngoing, err := s.adaptor.GetOngoingTechByComponentName(ctx, alarm.Component, alarm.Resource)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if existOngoing != nil {
	// 		updateFields := model.Alarm{
	// 			Crit:     alarm.Crit,
	// 			LastMsg:  alarm.LastMsg,
	// 			LastTime: alarm.LastTime,
	// 		}

	// 		return s.adaptor.UpdateTech(ctx, existOngoing.ID, updateFields)
	// 	}
	// 	return s.adaptor.CreateTech(ctx, alarm)
	// }

	// s.adaptor.GetAllTechniques(ctx)
	// s.adaptor.GetOngoingTechByComponentName(ctx, "Abc", "Dce")
	// s.adaptor.CreateTech(ctx)
	s.adaptor.UpdateOngoingTech(ctx, existOngoing.ID, model.Alarm{LastMsg: "updated by GO 3"})
	return nil
}
