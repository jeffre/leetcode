package leetcode7

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given int
	want  int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{-2147483648, 0},
	{1534236469, 0},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := reverse(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			reverse(c.given)
		}
	}
}
