package leetcode438

import (
	"reflect"
	"testing"
)

type given struct {
	s string
	p string
}

var cases = []struct {
	name  string
	given given
	want  []int
}{
	{
		name:  "example1",
		given: given{"cbaebabacd", "abc"},
		want:  []int{0, 6},
	}, {
		name:  "example2",
		given: given{"abab", "ab"},
		want:  []int{0, 1, 2},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := findAnagrams(tt.given.s, tt.given.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			findAnagrams(tt.given.s, tt.given.p)
		}
	}
}
