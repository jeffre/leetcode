package leetcode28

import (
	"fmt"
	"testing"
)

type given struct {
	haystack string
	needle   string
}

var cases = []struct {
	given given
	want  int
}{
	{
		given: given{haystack: "hello", needle: "ll"},
		want:  2,
	}, {
		given: given{haystack: "aaaaa", needle: "bba"},
		want:  -1,
	}, {
		given: given{haystack: "aaaaa", needle: ""},
		want:  0,
	},
}

func TestStrStr(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := strStr(tt.given.haystack, tt.given.needle)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			strStr(tt.given.haystack, tt.given.needle)
		}
	}
}
