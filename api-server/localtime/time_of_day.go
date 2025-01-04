package localtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

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

// Scan implements the [sql.Scanner] interface.
func (t *TimeOfDay) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case string:
		dv, err := ParseTimeOfDay(v)
		if err != nil {
			return fmt.Errorf("scan TimeOfDay: %v", err)
		}
		*t = dv
		return nil
	default:
		return fmt.Errorf("scan TimeOfDay: invalid type %T", value)
	}
}

// Value implements the [driver.Valuer] interface.
func (t TimeOfDay) Value() (driver.Value, error) {
	return t.String(), nil
}
