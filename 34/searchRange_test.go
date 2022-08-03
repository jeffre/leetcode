package leetcode34

import (
	"reflect"
	"testing"
)

type given struct {
	nums   []int
	target int
}

var cases = []struct {
	name  string
	given given
	want  []int
}{
	{
		name: "example1",
		given: given{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
		},
		want: []int{3, 4},
	}, {
		name: "example2",
		given: given{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
		},
		want: []int{-1, -1},
	}, {
		name: "example3",
		given: given{
			nums:   []int{},
			target: 0,
		},
		want: []int{-1, -1},
	}, {
		name: "test1",
		given: given{
			nums:   []int{1},
			target: 1,
		},
		want: []int{0, 0},
	}, {
		name: "test2",
		given: given{
			nums:   []int{2, 2},
			target: 3,
		},
		want: []int{-1, -1},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange(tt.given.nums, tt.given.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			searchRange(tt.given.nums, tt.given.target)
		}
	}
}
