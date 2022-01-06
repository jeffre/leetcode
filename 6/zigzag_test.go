package leetcode6

import (
	"fmt"
	"testing"
)

type given struct {
	s       string
	numRows int
}

var cases = []struct {
	given given
	want  string
}{
	{
		given: given{"PAYPALISHIRING", 3},
		want:  "PAHNAPLSIIGYIR",
	}, {
		given: given{"PAYPALISHIRING", 4},
		want:  "PINALSIGYAHRPI",
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := convert(c.given.s, c.given.numRows)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			convert(c.given.s, c.given.numRows)
		}
	}
}
