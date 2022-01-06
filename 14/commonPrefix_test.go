package leetcode14

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given []string
	want  string
}{
	{
		given: []string{"flower", "flow", "flight"},
		want:  "fl",
	},
	{
		given: []string{"dog", "racecar", "car"},
		want:  "",
	},
	{
		given: []string{"ab", "a", "abc"},
		want:  "a",
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := longestCommonPrefix(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			longestCommonPrefix(c.given)
		}
	}
}
