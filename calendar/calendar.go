package calendar

import "time"

type CalendarType string

const (
	TypeGoogleCalendar CalendarType = "google_calendar"
)

// Calendar is an interface that allows interation with an external
// calendar (e.g. Google Calendar)
type Fetcher interface {
	fetchEvents(time.Time) ([]CalendarEvent, error)
}

func Sync(t time.Time) error {
	gcal, err := NewCalendarGoogle("../config/credentials.json")
	if err != nil {
		return err
	}

	gcal.fetchEvents(t)
	return nil
}
