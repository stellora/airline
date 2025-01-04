package api

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestZonedDateTime(t *testing.T) {
	// Pick an obscure but real time zone name so that this test doesn't pass just because it is being
	// run on a machine set to this time zone.
	loc, err := time.LoadLocation("America/Adak")
	if err != nil {
		t.Fatal(err)
	}
	zonedDateTime := ZonedDateTime{time.Date(2025, 1, 2, 3, 4, 5, 0, loc)}
	rfc9557String := "2025-01-02T03:04:05-10:00[America/Adak]"
	jsonString := fmt.Sprintf(`"%s"`, rfc9557String)

	t.Run("MarshalJSON", func(t *testing.T) {
		got, err := zonedDateTime.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != jsonString {
			t.Errorf("got %q, want %q", string(got), jsonString)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var zdt ZonedDateTime
		if err := json.Unmarshal([]byte(jsonString), &zdt); err != nil {
			t.Fatal(err)
		}
		if got, want := zdt.Time.String(), zonedDateTime.Time.String(); got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		if got, want := zdt.Location().String(), zonedDateTime.Location().String(); got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
