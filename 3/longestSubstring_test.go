package leetcode3

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given string
	want  int
}{
	{
		given: "abcabcbb",
		want:  3,
	}, {
		given: "bbbbb",
		want:  1,
	}, {
		given: "pwwkew",
		want:  3,
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := lengthOfLongestSubstring(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}

		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			lengthOfLongestSubstring(c.given)
		}
	}
}
