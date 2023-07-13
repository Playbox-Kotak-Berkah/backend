package utils

import (
	"fmt"
	"time"
)

func MissingDates(startDate, endDate string) ([]string, int) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Println("Failed to parse start date:", err)
		return nil, 0
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		fmt.Println("Failed to parse end date:", err)
		return nil, 0
	}

	duration := end.Sub(start)
	totalDays := int(duration.Hours() / 24)

	missingDates := make([]string, 0)
	for i := 1; i < totalDays; i++ {
		missingDate := start.AddDate(0, 0, i).Format("2006-01-02")
		missingDates = append(missingDates, missingDate)
	}

	return missingDates, len(missingDates)
}
