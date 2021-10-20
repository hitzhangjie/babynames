package main

import "testing"

func Test_getChineseHour(t *testing.T) {
	tests := []struct {
		name string
		hour int
		want string
	}{
		{name: "1", hour: 1, want: "丑时"},
		{name: "2", hour: 2, want: "丑时"},
		{name: "3", hour: 3, want: "寅时"},
		{name: "4", hour: 4, want: "寅时"},
		{name: "5", hour: 5, want: "卯时"},
		{name: "6", hour: 6, want: "卯时"},
		{name: "7", hour: 7, want: "辰时"},
		{name: "8", hour: 8, want: "辰时"},
		{name: "23", hour: 23, want: "子时"},
		{name: "24", hour: 24, want: "子时"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChineseHour(tt.hour); got != tt.want {
				t.Errorf("getChineseHour() = %s, want %s", got, tt.want)
			}
		})
	}
}
