package leetcodeDummyNum

import (
	"testing"
)

type given struct {
	s     string
	words []string
}

var cases = []struct {
	name  string
	given given
	want  int
}{
	{
		name: "example1",
		given: given{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
		},
		want: 0,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := dummyFunc(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			dummyFunc(tt.given)
		}
	}
}
