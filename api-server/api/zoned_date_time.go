package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ZonedDateTime_ marshals and unmarshals a [time.Time] to/from JSON, using [RFC
// 9557](https://www.rfc-editor.org/rfc/rfc9557.html) IXDTF extensions to include the time zone
// name.
type ZonedDateTime_ struct {
	time.Time
}

func (t *ZonedDateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s[%s]", t.Time.Format(time.RFC3339), t.Time.Location().String()))
}

func (t *ZonedDateTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	idx := strings.Index(s, "[")
	if idx == -1 || !strings.HasSuffix(s, "]") {
		return fmt.Errorf("invalid ZonedDateTime (no time zone name): %q", s)
	}

	tzName := s[idx+1 : len(s)-1]
	timeWithoutTzName := s[:idx]

	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return fmt.Errorf("invalid ZonedDateTime (invalid time zone name): %q", tzName)
	}

	parsed, err := time.ParseInLocation(time.RFC3339, timeWithoutTzName, loc)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

var _ json.Marshaler = (*ZonedDateTime)(nil)
