package localtime

import (
	"reflect"
	"testing"
	"time"
)

func TestParseLocalDate(t *testing.T) {
	tests := []struct {
		input     string
		want      LocalDate
		wantError bool
	}{
		{input: "2024-08-13", want: LocalDate{Year: 2024, Month: time.August, Day: 13}},
		{input: "2024-08-13Z", wantError: true},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := ParseLocalDate(test.input)
			if err != nil {
				if test.wantError {
					return
				}
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("got %#v, want %#v", got, test.want)
			}
		})
	}
}
