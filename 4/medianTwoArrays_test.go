package leetcode4

import (
	"fmt"
	"testing"
)

var cases = []struct {
	given [][]int
	want  float64
}{
	{
		given: [][]int{
			{1, 3},
			{2},
		},
		want: 2.0,
	}, {
		given: [][]int{
			{1, 2},
			{3, 4},
		},
		want: 2.5,
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := findMedianSortedArrays(c.given[0], c.given[1])
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}

		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			findMedianSortedArrays(c.given[0], c.given[1])
		}
	}
}
