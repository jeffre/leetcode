package leetcode20

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given string
	want  bool
}{
	{given: "([])({})", want: true},
	{given: "()[]{}", want: true},
	{given: "([)]{}", want: false},
	{given: "([]){}", want: true},
	{given: "([]{)}", want: false},
	{given: "([]{})", want: true},
	{given: "([]{})", want: true},
	{given: "(]", want: false},
	{given: "(){}}{", want: false},
	{given: "}{", want: false},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := isValid(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			isValid(c.given)
		}
	}
}
