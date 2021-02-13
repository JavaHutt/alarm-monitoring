package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

// ParseAlarm parses incoming JSON message to struct
func (s *AlarmService) ParseAlarm(body []byte) (*model.Alarm, error) {
	parsed := new(model.Alarm)

	err := json.Unmarshal(body, parsed)
	if err != nil {
		log.Println("Failed to unmarshal incoming alarm")
		return nil, err
	}
	log.Println("Incoming alarm: ", *parsed)
	return parsed, nil
}

// InsertAlarm creates or updates alarm depending on Alarm state
func (s *AlarmService) InsertAlarm(ctx context.Context, alarm model.Alarm) error {
	fmt.Println("INSERT ALARM")
	existOngoing, err := s.adaptor.GetOngoingTechByComponentName(ctx, alarm.Component, alarm.Resource)
	if err != nil {
		return err
	}
	log.Printf("Crit level: %d", alarm.Crit)
	if alarm.Crit > model.Ok {
		log.Println("Crit is not OK")
		if existOngoing != nil {
			updateFields := model.Alarm{
				Crit:     alarm.Crit,
				LastMsg:  alarm.LastMsg,
				LastTime: alarm.LastTime,
			}

			return s.adaptor.UpdateOngoingTech(ctx, existOngoing.ID, updateFields)
		}
		return s.adaptor.CreateTech(ctx, alarm)
	} else if alarm.Crit == model.Ok {
		log.Println("Crit is OK")
		if existOngoing != nil {
			updateFields := model.Alarm{
				Crit:     alarm.Crit,
				LastMsg:  alarm.LastMsg,
				LastTime: alarm.LastTime,
			}

			return s.adaptor.UpdateResolvedTech(ctx, existOngoing.ID, updateFields)
		}
		log.Println("Continue recieving")
	}

	return nil
}
