package leetcode278

import "testing"

type given struct {
	n int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			n: 5,
		},
		want: 4,
	}, {
		name: "example2",
		given: given{
			n: 1,
		},
		want: 1,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			bad = tt.want
			got := firstBadVersion(tt.given.n)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			firstBadVersion(tt.given.n)
		}
	}
}
