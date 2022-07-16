package leetcode340

import (
	"testing"
)

type given struct {
	s string
	k int
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name:  "example1",
		given: given{"eceba", 2},
		want:  3,
	}, {
		name:  "example2",
		given: given{"aa", 1},
		want:  2,
	}, {
		name:  "example2",
		given: given{"aa", 1},
		want:  2,
	}, {
		name:  "big k",
		given: given{"ecebaxyzxzy", 5},
		want:  8,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstringKDistinct(tt.given.s, tt.given.k)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			lengthOfLongestSubstringKDistinct(tt.given.s, tt.given.k)
		}
	}
}
