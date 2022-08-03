package leetcode1482

import "testing"

type given struct {
	bloomDay []int
	m        int
	k        int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        1,
		},
		want: 3,
	}, {
		name: "example2",
		given: given{
			bloomDay: []int{1, 10, 3, 10, 2},
			m:        3,
			k:        2,
		},
		want: -1,
	}, {
		name: "example3",
		given: given{
			bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
			m:        2,
			k:        3,
		},
		want: 12,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := minDays(tt.given.bloomDay, tt.given.m, tt.given.k)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			minDays(tt.given.bloomDay, tt.given.m, tt.given.k)
		}
	}
}
