package calendarmath

import (
	"fmt"
	"time"

	datecalc "github.com/markusmobius/go-dateparser"
)

// Direction represents forward or backward calendar math
type Direction int

const (
	Forward  Direction = 1
	Backward Direction = -1
)

// Options controls how the date math is applied
type Options struct {
	Months            int
	Days              int
	BusinessDayOffset int         // e.g., 2 = +2 business days, -1 = back 1 business day
	AdjustToBusiness  bool        // If true, adjusts final date to nearest business day
	Holidays          []time.Time // Holiday list
	Direction         Direction   // Forward or Backward
}

// AddCalendarBusinessDays performs calendar math + business day offset and adjustment
func AddCalendarBusinessDays(start time.Time, opts Options) (time.Time, error) {
	// Step 1: Calendar math
	months := opts.Months * int(opts.Direction)
	days := opts.Days * int(opts.Direction)
	result := start.AddDate(0, months, days)

	// Step 2: Set up calendar
	cal := datecalc.NewCalendar()
	cal.AddWeekend(time.Saturday)
	cal.AddWeekend(time.Sunday)

	for _, h := range opts.Holidays {
		cal.AddHoliday(h)
	}

	// Step 3: Apply business day offset
	if opts.BusinessDayOffset != 0 {
		var err error
		result, err = cal.AddBusinessDays(result, opts.BusinessDayOffset)
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to add business days: %w", err)
		}
	}

	// Step 4: Adjust final date to business day if requested
	if opts.AdjustToBusiness && !cal.IsBusinessDay(result) {
		var err error
		if opts.Direction == Forward {
			result, err = cal.NextBusinessDay(result)
		} else {
			result, err = cal.PreviousBusinessDay(result)
		}
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to adjust to business day: %w", err)
		}
	}

	return result, nil
}
