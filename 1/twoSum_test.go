package leetcode1

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
	want  []int
}{
	{
		given: given{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		want: []int{0, 1},
	}, {
		given: given{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		want: []int{1, 2},
	}, {
		given: given{
			nums:   []int{3, 3},
			target: 6,
		},
		want: []int{0, 1},
	},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := twoSum(c.given.nums, c.given.target)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			twoSum(c.given.nums, c.given.target)
		}
	}
}
