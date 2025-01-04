package localtime

import (
	"reflect"
	"testing"
)

func TestParseTimeOfDay(t *testing.T) {
	tests := []struct {
		input     string
		want      TimeOfDay
		wantError bool
	}{
		{input: "19:00", want: TimeOfDay{Hour: 19, Minute: 0}},
		{input: "07:04", want: TimeOfDay{Hour: 7, Minute: 4}},
		{input: "07:04:05", wantError: true},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := ParseTimeOfDay(test.input)
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
