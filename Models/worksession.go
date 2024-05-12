package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkSession struct {
	gorm.Model
	StartTime   time.Time
	EndTime     time.Time
	HoursWorked float64 `gorm:"-"`
}

func (ws *WorkSession) CalculateHoursWorked() {
	if ws.EndTime.After(ws.StartTime) {
		duration := ws.EndTime.Sub(ws.StartTime)
		ws.HoursWorked = duration.Hours()
	}
}
