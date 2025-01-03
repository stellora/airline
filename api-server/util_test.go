package main

import (
	"testing"
)

func TestParseDaysOfWeek(t *testing.T) {
	tests := []struct {
		str      string
		wantDays []int
		wantErr  string
	}{
		{str: "01356", wantDays: []int{0, 1, 3, 5, 6}},
		{str: "123456", wantDays: []int{1, 2, 3, 4, 5, 6}},
		{str: "0", wantDays: []int{0}},
		{str: "6", wantDays: []int{6}},
		{str: "", wantDays: []int{}},
		{str: "0456", wantDays: []int{4, 5, 6}},
		{str: "27", wantErr: "invalid day of week"},
		{str: "1234567", wantErr: "invalid day of week"},
	}

	for _, tc := range tests {
		t.Run(tc.str, func(t *testing.T) {
			days, err := parseDaysOfWeek(tc.str)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatal(err)
				}
				assertEqual(t, days, tc.wantDays)
			} else {
				if err == nil {
					t.Fatalf("expected error: %q", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("got error %s, want %s", err.Error(), tc.wantErr)
				}
			}
		})
	}
}
