package leetcode10

import (
	"fmt"
	"testing"
)

type given struct {
	string  string
	pattern string
}

var cases = []struct {
	given given
	want  bool
}{
	{given{"aa", "a"}, false},
	{given{"aa", "a*"}, true},
	{given{"ab", ".*"}, true},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := isMatch(c.given.string, c.given.pattern)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			isMatch(c.given.string, c.given.pattern)
		}
	}
}
