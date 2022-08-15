package leetcode134

import "testing"

type given struct {
	gas  []int
	cost []int
}

type want struct {
	int
}

func runTest(given given) want {
	got := canCompleteCircuit(given.gas, given.cost)
	return want{got}
}

var tests = []test{
	{
		"example1",
		given{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
		},
		want{3},
	}, {
		"example2",
		given{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
		},
		want{-1},
	},
}

type test struct {
	name  string
	given given
	want  want
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := runTest(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			runTest(tt.given)
		}
	}
}
