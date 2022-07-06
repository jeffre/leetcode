package leetcode31

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
		given: []int{1, 2, 3},
		want:  []int{1, 3, 2},
	},
	{
		name:  "example2",
		given: []int{3, 2, 1},
		want:  []int{1, 2, 3},
	},
	{
		name:  "example3",
		given: []int{1, 1, 5},
		want:  []int{1, 5, 1},
	},
	{
		name:  "test1",
		given: []int{1, 2, 1},
		want:  []int{2, 1, 1},
	},
	{
		name:  "test2",
		given: []int{1, 3, 2},
		want:  []int{2, 1, 3},
	},
	{
		name:  "test3",
		given: []int{2, 3, 1},
		want:  []int{3, 1, 2},
	},
	{
		name:  "test4",
		given: []int{4, 2, 0, 2, 3, 2, 0},
		want:  []int{4, 2, 0, 3, 0, 2, 2},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			nextPermutation(tt.given)
			if !reflect.DeepEqual(tt.given, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", tt.given, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			nextPermutation(tt.given)
		}
	}
}
