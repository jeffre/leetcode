package leetcode30

import (
	"reflect"
	"testing"
)

type given struct {
	s     string
	words []string
}

var cases = []struct {
	name  string
	given given
	want  []int
}{
	{
		name:  "example 1",
		given: given{s: "barfoothefoobarman", words: []string{"foo", "bar"}},
		want:  []int{0, 9},
	}, {
		name:  "example 2",
		given: given{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "word"}},
		want:  []int{},
	}, {
		name:  "example 3",
		given: given{s: "barfoofoobarthefoobarman", words: []string{"bar", "foo", "the"}},
		want:  []int{6, 9, 12},
	}, {
		name:  "submitted 1",
		given: given{s: "wordgoodgoodgoodbestword", words: []string{"word", "good", "best", "good"}},
		want:  []int{8},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.given.s, tt.given.words)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			findSubstring(tt.given.s, tt.given.words)
		}
	}
}
