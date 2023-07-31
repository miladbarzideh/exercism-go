package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	return convertToString("1/2/2006 15:04:05", date)
}

func convertToString(layout string, value string) time.Time {
	t, error := time.Parse(layout, value)
	if (error != nil) {
		fmt.Println(error)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := convertToString("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t := convertToString("Monday, January 2, 2006 15:04:05", date)
	h := t.Hour()
	return h >= 12 && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := convertToString("1/2/2006 15:04:05", date)
	formattedDate := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedDate)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Now()
	return time.Date(t.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
