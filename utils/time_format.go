package utils

import (
	"time"
)

// FormatDate formats a date.
func FormatDate(date time.Time) string {
	return date.Format("Mon _2 Jan 2006 15:04:05")
}

func FormatTimestamp2String(timestamp int64) string {
	// Replace 1705154613 with your timestamp
	// timestamp := int64(1705154613)

	// Load the "Asia/Shanghai" time zone
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// fmt.Println("Error loading location:", err)
		return ""
	}

	// Convert Unix timestamp to time.Time in the specified time zone
	t := time.Unix(timestamp, 0).In(loc)

	// Print the resulting time in a formatted string
	formattedTime := t.Format("2006-01-02 15:04:05 MST")
	// fmt.Println("Formatted Time (Asia/Shanghai):", formattedTime)
	return formattedTime
}
