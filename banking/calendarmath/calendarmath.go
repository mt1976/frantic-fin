package calendarmath

import (
	"fmt"
	"time"

	"github.com/scmhub/calendar"
)

// Direction for date math
type Direction int

const (
	Forward  Direction = 1
	Backward Direction = -1
)

// Options defines how the date calculation works
type Options struct {
	Months            int                // calendar months to add/subtract
	Days              int                // calendar days to add/subtract
	BusinessDayOffset int                // business-day offset after calendar math
	AdjustToBusiness  bool               // snap final date to workday if off
	Direction         Direction          // forward or backward
	ExchangeCalendar  *calendar.Calendar // e.g. NYSE, LSE, etc.
}

// GetNextWorkingDay computes next working day
func GetNextWorkingDay(start time.Time) (time.Time, error) {
	// Apply calendar math

	// Default to NYSE if no calendar provided

	cal2 := calendar.NewCalendar("test", time.Local, time.Now().Year())

	nbd := cal2.NextBusinessDay(start)

	return nbd, nil
}

// GetPreviousWorkingDay computes the prvious working day
func GetPreviousWorkingDay(start time.Time) (time.Time, error) {
	cal2 := calendar.NewCalendar("test", time.Local, time.Now().Year())
	upd := start.AddDate(0, 0, -2)
	pbd := cal2.NextBusinessDay(upd)
	if pbd.IsZero() {
		return time.Time{}, fmt.Errorf("no previous business day found")
	}
	return pbd, nil
}

// IsWorkingDay checks if the date is a working day
func IsWorkingDay(date time.Time) (bool, error) {
	cal2 := calendar.NewCalendar("test", time.Local, time.Now().Year())
	wd := cal2.IsBusinessDay(date)
	if wd {
		return true, nil
	}
	return false, fmt.Errorf("not a working day")
}
