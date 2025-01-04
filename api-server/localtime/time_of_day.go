package localtime

import (
	"fmt"
	"time"
)

func (d LocalDate) TimeOfDay(loc *time.Location, at TimeOfDay) time.Time {
	if loc == nil || loc.String() == "UTC" {
		panic("converting a LocalDate to UTC is almost always a mistake")
	}
	return time.Date(d.Year, d.Month, d.Day, at.Hour, at.Minute, 0, 0, loc)
}

// A TimeOfDay is a timezone-naive time of day. It is a time of day that does not have a
// timezone attached to it.
type TimeOfDay struct {
	Hour, Minute int
}

func NewTimeOfDay(hour, minute int) TimeOfDay {
	if hour < 0 || hour > 23 {
		panic("hour must be between 0 and 23")
	}
	if minute < 0 || minute > 59 {
		panic("minute must be between 0 and 59")
	}
	return TimeOfDay{Hour: hour, Minute: minute}
}

func ParseTimeOfDay(s string) (TimeOfDay, error) {
	const HourMinuteOnly = "15:04"
	t, err := time.Parse(HourMinuteOnly, s)
	if err != nil {
		return TimeOfDay{}, err
	}
	return TimeOfDay{Hour: t.Hour(), Minute: t.Minute()}, nil
}

func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d", t.Hour, t.Minute)
}
