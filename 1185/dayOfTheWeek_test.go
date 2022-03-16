package leetcode1185

import (
	"fmt"
	"testing"
)

type given struct {
	day   int
	month int
	year  int
}

var cases = []struct {
	given given
	want  string
}{
	{
		given: given{31, 8, 2019},
		want:  "Saturday",
	},
	{
		given: given{18, 7, 1999},
		want:  "Sunday",
	},
	{
		given: given{15, 8, 1993},
		want:  "Sunday",
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := dayOfTheWeek(tt.given.day, tt.given.month, tt.given.year)
			if got != tt.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", tt.given, got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			dayOfTheWeek(tt.given.day, tt.given.month, tt.given.year)
		}
	}
}
