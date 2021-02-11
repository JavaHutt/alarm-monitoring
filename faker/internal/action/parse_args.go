package action

import (
	"strconv"
	"time"
)

const defaultDuration = 10

// GetDuration gets the duration passed in script arguments
func GetDuration(args []string) time.Duration {
	var d int
	if len(args) < 2 || args[1] == "" {
		return defaultDuration
	}

	d, err := strconv.Atoi(args[1])
	if err != nil {
		return defaultDuration
	}
	if d < 1 {
		return defaultDuration
	}
	return time.Duration(d)
}
