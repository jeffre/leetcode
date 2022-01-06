package leetcode9

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given int
	want  bool
}{
	{0, true},
	{121, true},
	{1221, true},
	{-121, false},
	{10, false},
	{1000021, false},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := isPalindrome(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			isPalindrome(c.given)
		}
	}
}
