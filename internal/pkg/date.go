package pkg

import "time"

func GetTime() string {
	currentTime := time.Now()
	formattedDate := currentTime.Format(time.RFC3339)

	return formattedDate
}
