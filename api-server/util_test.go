package main

import (
	"testing"
)

func TestParseDaysOfWeek(t *testing.T) {
	tests := []struct {
		str      string
		wantDays []int
		wantStr  string
		wantErr  string
	}{
		{str: "01356", wantDays: []int{0, 1, 3, 5, 6}},
		{str: "123456", wantDays: []int{1, 2, 3, 4, 5, 6}},
		{str: "0", wantDays: []int{0}},
		{str: "6", wantDays: []int{6}},
		{str: "", wantDays: nil},
		{str: "0456", wantDays: []int{0, 4, 5, 6}},
		{str: "004556", wantDays: []int{0, 4, 5, 6}, wantStr: "0456"}, // non-unique but tolerate it
		{str: "61", wantDays: []int{1, 6}, wantStr: "16"},             // non-unique but tolerate it
		{str: "27", wantErr: "invalid day of week"},
		{str: "1234567", wantErr: "invalid day of week"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			days, err := parseDaysOfWeek(test.str)
			if test.wantErr == "" {
				if err != nil {
					t.Fatal(err)
				}
				assertEqual(t, days, test.wantDays)

				reverse := toDBDaysOfWeek(days)
				wantStr := test.wantStr
				if wantStr == "" {
					wantStr = test.str
				}
				assertEqual(t, reverse, wantStr)
			} else {
				if err == nil {
					t.Fatalf("expected error: %q", test.wantErr)
				}
				if err.Error() != test.wantErr {
					t.Fatalf("got error %q, want %q", err.Error(), test.wantErr)
				}
			}
		})
	}
}
