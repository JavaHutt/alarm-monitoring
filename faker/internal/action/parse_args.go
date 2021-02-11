package action

import (
	"os"
	"strconv"
	"time"
)

const defaultDuration = 10

// GetDuration gets the duration passed in script arguments
func GetDuration(args []string) time.Duration {
	var d int
	if len(os.Args) < 2 || os.Args[1] == "" {
		return defaultDuration
	}

	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return defaultDuration
	}
	if d < 1 {
		return defaultDuration
	}
	return time.Duration(d)
}
