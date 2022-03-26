package leetcode76

import (
	"fmt"
	"testing"
)

type given struct {
	s string
	t string
}

var cases = []struct {
	given given
	want  string
}{
	{
		given: given{"ADOBECODEBANC", "ABC"},
		want:  "BANC",
	},
	{
		given: given{"a", "a"},
		want:  "a",
	},
	{
		given: given{"a", "aa"},
		want:  "",
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := minWindow(tt.given.s, tt.given.t)
			if got != tt.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", tt.given, got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			minWindow(tt.given.s, tt.given.t)
		}
	}
}
