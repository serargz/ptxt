package calendar

import "time"

type CalendarEvent struct {
	date        time.Time
	title       string
	description string
	notes       string
}
