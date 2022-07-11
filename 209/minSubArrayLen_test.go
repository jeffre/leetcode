package leetcode209

import (
	"testing"
)

type given struct {
	target int
	nums   []int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
		},
		want: 2,
	}, {
		name: "example2",
		given: given{
			target: 4,
			nums:   []int{1, 4, 4},
		},
		want: 1,
	}, {
		name: "example3",
		given: given{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		want: 0,
	}, {
		name: "secondsubarray",
		given: given{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3, 0, 7, 0},
		},
		want: 1,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.given.target, tt.given.nums)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			minSubArrayLen(tt.given.target, tt.given.nums)
		}
	}
}
