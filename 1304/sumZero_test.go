package leetcode1304

import (
	"reflect"
	"testing"
)

var cases = []struct {
	name  string
	given int
	want  []int
}{
	{
		name:  "example1",
		given: 5,
		want:  []int{-2, -1, 0, 1, 2},
	}, {
		name:  "example2",
		given: 3,
		want:  []int{-1, 0, 1},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := sumZero(tt.given)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			sumZero(tt.given)
		}
	}
}
