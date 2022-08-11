package leetcode134

import "testing"

type given struct {
	gas  []int
	cost []int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
		},
		want: 3,
	}, {
		name: "example2",
		given: given{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
		},
		want: -1,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.given.gas, tt.given.cost)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			canCompleteCircuit(tt.given.gas, tt.given.cost)
		}
	}
}
