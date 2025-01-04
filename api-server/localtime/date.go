package localtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// A LocalDate is a timezone-naive date. It is a calendar date that does not have a timezone
// attached to it.
type LocalDate struct {
	Year  int
	Month time.Month
	Day   int
}

// ParseLocalDate parses a date in RFC3339 format. It returns an error if the date is not in the
// format "YYYY-MM-DD".
func ParseLocalDate(s string) (LocalDate, error) {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return LocalDate{}, err
	}
	return NewLocalDate(t.Year(), t.Month(), t.Day()), nil
}

func NewLocalDate(year int, month time.Month, day int) LocalDate {
	return LocalDate{Year: year, Month: month, Day: day}
}

func (d LocalDate) Date(loc *time.Location) time.Time {
	if loc == nil || loc.String() == "UTC" {
		panic("converting a LocalDate to UTC is almost always a mistake")
	}
	return d.date(loc)
}

func (d LocalDate) date(loc *time.Location) time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
}

func (d LocalDate) TimeOfDay(loc *time.Location, at TimeOfDay) time.Time {
	if loc == nil || loc.String() == "UTC" {
		panic("converting a LocalDate to UTC is almost always a mistake")
	}
	return time.Date(d.Year, d.Month, d.Day, at.Hour, at.Minute, 0, 0, loc)
}

func (d LocalDate) Equal(other LocalDate) bool {
	return d.Year == other.Year && d.Month == other.Month && d.Day == other.Day
}

func (d LocalDate) After(other LocalDate) bool {
	if d.Year != other.Year {
		return d.Year > other.Year
	}
	if d.Month != other.Month {
		return d.Month > other.Month
	}
	return d.Day > other.Day
}

func (d LocalDate) Before(other LocalDate) bool {
	if d.Year != other.Year {
		return d.Year < other.Year
	}
	if d.Month != other.Month {
		return d.Month < other.Month
	}
	return d.Day < other.Day
}

func (d LocalDate) Weekday() time.Weekday {
	return d.date(time.UTC).Weekday()
}

func (d LocalDate) AddDays(days int) LocalDate {
	t := d.date(time.UTC).AddDate(0, 0, days)
	return NewLocalDate(t.Year(), t.Month(), t.Day())
}

func (d LocalDate) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

// Scan implements the [sql.Scanner] interface.
func (d *LocalDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case string:
		dv, err := ParseLocalDate(v)
		if err != nil {
			return fmt.Errorf("scan LocalDate: %v", err)
		}
		*d = dv
		return nil
	default:
		return fmt.Errorf("scan LocalDate: invalid type %T", value)
	}
}

// Value implements the [driver.Valuer] interface.
func (d LocalDate) Value() (driver.Value, error) {
	return d.String(), nil
}
