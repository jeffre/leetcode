package leetcode5

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given string
	want  string
}{
	{
		given: "",
		want:  "",
	}, {
		given: "aaa",
		want:  "aaa",
	}, {
		given: "123",
		want:  "1",
	}, {
		given: "babad",
		want:  "bab",
	}, {
		given: "cbbd",
		want:  "bb",
	}, {
		given: "tattarrattat",
		want:  "tattarrattat",
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := longestPalindrome(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}

		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			longestPalindrome(c.given)
		}
	}
}
