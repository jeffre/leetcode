package leetcode8

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given string
	want  int
}{
	{"42", 42},
	{"   -42", -42},
	{"4193 with words", 4193},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := myAtoi(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			myAtoi(c.given)
		}
	}
}
