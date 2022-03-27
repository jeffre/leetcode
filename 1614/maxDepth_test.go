package leetcode1614

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given string
	want  int
}{
	{"(1+(2*3)+((8)/4))+1", 3},
	{"(1)+((2))+(((3)))", 3},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := maxDepth(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			maxDepth(c.given)
		}
	}
}
