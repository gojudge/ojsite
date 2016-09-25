package service

import (
	"time"
)

const (
	DATE = "DATE"
	TIME = "TIME"
)

func TimeFormatSimple(t time.Time, types string) string {
	switch types {
	case DATE:
		return TimeFormat(t, "2006-01-02")
	case TIME:
		return TimeFormat(t, "2006-01-02 15:04:05")
	default:
		return t.Format("2006-01-02T15:04:05Z07:00")
	}
}

func TimeFormat(t time.Time, format string) string {
	return t.Format(format)
}
