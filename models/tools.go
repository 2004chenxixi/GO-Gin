package tools

import (
	"time"
)

func GetDay() string {
	test := "1650422138"
	return time.Now().Format(test)
}
func GetUnix() int64 {
	return time.Now().Unix()
}

func Day() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	return timeStr
}
func DAY() string {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr
}
