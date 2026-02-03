package pkg

import "time"

func GetTodayDate() string {
	currentTime := time.Now()
	formattedDate := currentTime.Format("2006-01-02")

	return formattedDate
}
