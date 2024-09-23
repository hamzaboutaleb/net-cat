package utils

import "time"

func FormatDate(now time.Time) string {
	layout := "2006-01-02 15:04:05"
	return now.Format(layout)
}
