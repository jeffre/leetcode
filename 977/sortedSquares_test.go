package leetcode977

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name  string
	given []int
	want  []int
}{
	{
		name:  "example1",
		given: []int{-4, -1, 0, 3, 10},
		want:  []int{0, 1, 9, 16, 100},
	}, {
		name:  "example1",
		given: []int{-7, -3, 2, 3, 11},
		want:  []int{4, 9, 9, 49, 121},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.given)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			sortedSquares(tt.given)
		}
	}
}
