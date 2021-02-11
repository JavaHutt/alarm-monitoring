package model

type crit int

const (
	ok crit = iota
	minor
	major
	critical
)

type status string

const (
	// Ongoing issue
	Ongoing status = "ONGOING"
	// Resolved issue
	Resolved status = "RESOLVED"
)

// Alarm is the data model
type Alarm struct {
	Component string `json:"component"  bson:"component"`
	Resource  string `json:"resource"   bson:"resource"`
	Crit      crit   `json:"crit"       bson:"crit"`
	LastMsg   string `json:"last_msg"   bson:"last_msg"`
	FirstMsg  string `json:"first_msg"  bson:"first_msg"`
	StartTime int    `json:"start_time" bson:"start_time"`
	LastTime  int    `json:"last_time"  bson:"last_time"`
	Status    status `json:"status"     bson:"status"`
}

func (c crit) String() string {
	return [4]string{"OK", "Minor", "Major", "Critical"}[c]
}
