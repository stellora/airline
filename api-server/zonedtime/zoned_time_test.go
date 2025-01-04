package zonedtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestZonedTime(t *testing.T) {
	// Pick an obscure but real time zone name so that this test doesn't pass just because it is being
	// run on a machine set to this time zone.
	loc, err := time.LoadLocation("America/Adak")
	if err != nil {
		t.Fatal(err)
	}
	zonedDateTime := ZonedTime{time.Date(2025, 1, 2, 3, 4, 5, 0, loc)}
	rfc9557String := "2025-01-02T03:04:05-10:00[America/Adak]"
	jsonString := fmt.Sprintf(`"%s"`, rfc9557String)

	t.Run("MarshalJSON", func(t *testing.T) {
		got, err := json.Marshal(zonedDateTime)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != jsonString {
			t.Errorf("got %q, want %q", string(got), jsonString)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var zdt ZonedTime
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

	t.Run("no automatic mapping of Europe/London to UTC", func(t *testing.T) {
		t.Run("MarshalText", func(t *testing.T) {
			europeLondon, err := time.LoadLocation("Europe/London")
			if err != nil {
				t.Fatal(err)
			}
			tm := time.Date(2025, 2, 6, 10, 30, 0, 0, europeLondon)
			zdt := ZonedTime{tm}
			text, err := zdt.MarshalText()
			if err != nil {
				t.Fatal(err)
			}
			if got, want := string(text), "2025-02-06T10:30:00+00:00[Europe/London]"; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("UnmarshalText", func(t *testing.T) {
			rfc9557String := "2025-02-06T10:30:00+00:00[Europe/London]"
			var zdt ZonedTime
			if err := zdt.UnmarshalText([]byte(rfc9557String)); err != nil {
				t.Fatal(err)
			}
			if got, want := zdt.Location().String(), "Europe/London"; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	})
}
