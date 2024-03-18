package utilities

import (
	"strconv"
	"time"
)

func StringToInt(value string) (int, error) {
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func StringToDuration(value string) (time.Duration, error) {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, err
	}

	return duration, nil
}
