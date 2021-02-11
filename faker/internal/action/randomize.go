package action

import (
	"faker/internal/model"
	"math/rand"
	"time"
)

const messageLength = 7

// GetRandomAlarm mocks Alarm data
func GetRandomAlarm() model.Alarm {
	component := getRandomComponent()
	resource := getRandomResource()
	crit := getRandomCrit()
	lastMsg := getRandomMessage(messageLength)
	firstMsg := getRandomMessage(messageLength)
	startTime := time.Now()
	lastTime := time.Now()
	status := getRandomStatus()

	return model.Alarm{
		Component: component,
		Resource:  resource,
		Crit:      crit,
		LastMsg:   lastMsg,
		FirstMsg:  firstMsg,
		StartTime: startTime,
		LastTime:  lastTime,
		Status:    status,
	}
}

func getRandomComponent() string {
	rand.Seed(time.Now().Unix())
	return components[rand.Intn(len(components))]
}

func getRandomResource() string {
	rand.Seed(time.Now().Unix() * 2)
	return components[rand.Intn(len(components))]
}

func getRandomCrit() model.Crit {
	rand.Seed(time.Now().Unix() * 3)
	return crits[rand.Intn(len(crits))]
}

func getRandomStatus() model.Status {
	rand.Seed(time.Now().Unix() * 4)
	return statuses[rand.Intn(len(statuses))]
}

func getRandomMessage(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var components = []string{
	"CPU",
	"RAM",
	"DIMM",
	"HDD",
	"SSD",
	"Graphic card",
}

var statuses = []model.Status{model.Ongoing, model.Resolved}

var crits = []model.Crit{0, 1, 2, 3}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
