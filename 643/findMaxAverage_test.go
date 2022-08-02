package leetcode643

import (
	"testing"
)

type given struct {
	nums []int
	k    int
}

var cases = []struct {
	name  string
	given given
	want  float64
}{
	{
		name: "example1",
		given: given{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
		},
		want: 12.75000,
	}, {
		name: "example2",
		given: given{
			nums: []int{5},
			k:    1,
		},
		want: 5.0,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaxAverage(tt.given.nums, tt.given.k)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			findMaxAverage(tt.given.nums, tt.given.k)
		}
	}
}
