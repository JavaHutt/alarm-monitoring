package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Crit is alarm critical level
type Crit int

const (
	ok Crit = iota
	minor
	major
	critical
)

// Status is alarm status
type Status string

const (
	// Ongoing issue
	Ongoing Status = "ONGOING"
	// Resolved issue
	Resolved Status = "RESOLVED"
)

// Alarm is the data model
// I used
type Alarm struct {
	ID        primitive.ObjectID `json:"_id"        bson:"_id"`
	Component string             `json:"component"  bson:"component"`
	Resource  string             `json:"resource"   bson:"resource"`
	Crit      Crit               `json:"crit"       bson:"crit"`
	LastMsg   string             `json:"last_msg"   bson:"last_msg"`
	FirstMsg  string             `json:"first_msg"  bson:"first_msg"`
	StartTime time.Time          `json:"start_time" bson:"start_time"`
	LastTime  time.Time          `json:"last_time"  bson:"last_time"`
	Status    Status             `json:"status"     bson:"status"`
}

func (c Crit) String() string {
	return [4]string{"OK", "Minor", "Major", "Critical"}[c]
}
