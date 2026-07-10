package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func openWeeklyLogFile() (*os.File, error) {
	now := time.Now()
	year := now.Year()
	month := now.Month()

	startDay := ((now.Day()-1)/7)*7 + 1
	endDay := startDay + 6

	lastDay := time.Date(
		year,
		month+1,
		0, 0, 0, 0, 0, time.Local).Day()
	if endDay > lastDay {
		endDay = lastDay
	}

	week := ((now.Day() - 1) / 7) + 1

	dir := filepath.Join(
		"logs",
		fmt.Sprintf("%d", year),
		fmt.Sprintf("%02d", month),
	)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	filename := fmt.Sprintf(
		"week%d_%02d_%02d.log",
		week,
		startDay,
		endDay,
	)

	return os.OpenFile(
		filepath.Join(dir, filename),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
}
