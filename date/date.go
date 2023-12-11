package main

import (
	"time"
)

func GenerateUTCDateBefore(inputDuration string) (string, error) {
	duration, err := time.ParseDuration(inputDuration)
	if err != nil {
		return "", err
	}

	currentTime := time.Now().UTC()

	targetTime := currentTime.Add(-duration)

	utcFormat := "2006-01-02T15:04:05.999Z"
	result := targetTime.Format(utcFormat)

	return result, nil
}
