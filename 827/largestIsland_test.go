package leetcode827

import "testing"

type given struct {
	grid [][]int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			grid: [][]int{
				{1, 0},
				{0, 1},
			},
		},
		want: 3,
	}, {
		name: "example2",
		given: given{
			grid: [][]int{
				{1, 1},
				{1, 0},
			},
		},
		want: 4,
	}, {
		name: "example3",
		given: given{
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		want: 4,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := largestIsland(tt.given.grid)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			largestIsland(tt.given.grid)
		}
	}
}
