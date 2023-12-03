package util

import (
	"fmt"
	"time"
)

func TimeIsoToFormat(date string) (string, error) {

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return "", err
	}

	desiredLayout := "02-01-2006 15:04:05"
	formattedTime := parsedTime.Format(desiredLayout)

	return formattedTime, nil
}

func GetTimeOnly(date string) (string, error) {

	parsedTime, err := time.Parse(time.TimeOnly, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return "", err
	}

	return parsedTime.String(), nil
}

func GetTimeOneMinuteAgo() *time.Time {
	currentTime := time.Now()
	oneMinuteAgo := currentTime.Add(-5 * time.Minute)

	return &oneMinuteAgo
}
