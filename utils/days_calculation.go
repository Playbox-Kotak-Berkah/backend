package utils

import (
	"fmt"
	"time"
)

func CountDays(startDate, endDate string) int {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Println("Failed to parse start date:", err)
		return 0
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		fmt.Println("Failed to parse end date:", err)
		return 0
	}

	duration := end.Sub(start)

	days := int(duration.Hours() / 24)

	return days
}
