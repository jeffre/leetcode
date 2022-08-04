package leetcode33

import "testing"

type given struct {
	nums   []int
	target int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
		},
		want: 4,
	},
	{
		name: "example2",
		given: given{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
		},
		want: -1,
	}, {
		name: "example3",
		given: given{
			nums:   []int{1},
			target: 0,
		},
		want: -1,
	}, {
		name: "sorted",
		given: given{
			nums:   []int{1, 2, 3, 4, 5},
			target: 4,
		},
		want: 3,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.given.nums, tt.given.target)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			search(tt.given.nums, tt.given.target)
		}
	}
}
