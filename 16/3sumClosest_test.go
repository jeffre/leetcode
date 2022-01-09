package leetcode16

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	num    []int
	target int
}

var cases = []struct {
	given given
	want  int
}{
	{
		given: given{
			num:    []int{-1, 2, 1, -4},
			target: 1,
		},
		want: 2,
	}, {
		given: given{
			num:    []int{0, 0, 0},
			target: 1,
		},
		want: 0,
	}, {
		given: given{
			num:    []int{0, 2, 1, -3},
			target: 1,
		},
		want: 0,
	}, {
		given: given{
			num:    []int{-100, -98, -2, -1},
			target: -101,
		},
		want: -101,
	}, {
		given: given{
			num:    []int{-55, -24, -18, -11, -7, -3, 4, 5, 6, 9, 11, 23, 33},
			target: 0,
		},
		want: 0,
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := threeSumClosest(c.given.num, c.given.target)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			threeSumClosest(c.given.num, c.given.target)
		}
	}
}
