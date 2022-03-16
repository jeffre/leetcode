package leetcode374

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given int
	want  int
}{
	{
		given: 10,
		want:  1,
	},
	{
		given: 10,
		want:  10,
	},
	{
		given: 10,
		want:  5,
	},
	{
		given: 200,
		want:  4,
	},
	{
		given: 2000,
		want:  4,
	},
	{
		given: 9871235,
		want:  881,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			answer = tt.want
			got := guessNumber(tt.given)
			if got != tt.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", tt.given, got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			answer = tt.want
			guessNumber(tt.given)
		}
	}
}
