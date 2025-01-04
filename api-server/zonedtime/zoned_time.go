package zonedtime

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// ZonedTime marshals and unmarshals a [time.Time] to/from JSON, using [RFC
// 9557](https://www.rfc-editor.org/rfc/rfc9557.html) IXDTF extensions to include the time zone
// name.
type ZonedTime struct {
	time.Time
}

func (t ZonedTime) FormatRFC9557() string {
	return fmt.Sprintf("%s[%s]", t.Time.Format(time.RFC3339), t.Time.Location().String())
}

func (t ZonedTime) MarshalText() ([]byte, error) {
	return []byte(t.FormatRFC9557()), nil
}

func (t *ZonedTime) UnmarshalText(data []byte) error {
	idx := bytes.IndexByte(data, '[')
	if idx == -1 || len(data) == 0 || data[len(data)-1] != ']' {
		return fmt.Errorf("invalid ZonedTime (no time zone name): %q", data)
	}

	tzName := string(data[idx+1 : len(data)-1])
	timeWithoutTzName := string(data[:idx])

	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return fmt.Errorf("invalid ZonedTime (invalid time zone name): %q", tzName)
	}

	parsed, err := time.ParseInLocation(time.RFC3339, timeWithoutTzName, loc)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

func (t ZonedTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.FormatRFC9557())
}

func (t *ZonedTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return t.UnmarshalText([]byte(s))
}

// Value implements the [database/sql.Valuer] interface.
func (t ZonedTime) Value() (driver.Value, error) {
	return t.FormatRFC9557(), nil
}

// Scan implements the [database/sql.Scanner] interface.
func (t *ZonedTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case string:
		return t.UnmarshalText([]byte(v))
	default:
		return fmt.Errorf("scan ZonedTime: invalid type %T", value)
	}
}
