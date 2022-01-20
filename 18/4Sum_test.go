package leetcode18

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	nums   []int
	target int
}

var cases = []struct {
	given given
	want  [][]int
}{
	{
		given: given{
			[]int{1, 0, -1, 0, -2, 2},
			0,
		},
		want: [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		},
	}, {
		given: given{
			[]int{2, 2, 2, 2, 2},
			8,
		},
		want: [][]int{
			{2, 2, 2, 2},
		},
	}, {
		given: given{
			[]int{0, 0, 0, 0},
			0,
		},
		want: [][]int{
			{0, 0, 0, 0},
		},
	}, {
		given: given{
			[]int{1, -2, -5, -4, -3, 3, 3, 5},
			-11,
		},
		want: [][]int{
			{-5, -4, -3, 1},
		},
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := fourSum(c.given.nums, c.given.target)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			fourSum(c.given.nums, c.given.target)
		}
	}
}
