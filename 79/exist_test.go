package leetcode79

import "testing"

type given struct {
	board [][]byte
	word  string
}

type result struct {
	bool
}

func (tt *test) run() result {
	got := exist(tt.given.board, tt.given.word)
	return result{got}
}

var tests []test

// Add example1
func init() {
	tests = append(tests, test{
		name: "example1",
		given: given{
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word: "ABCCED",
		},
		want: result{true},
	})
}

type test struct {
	name  string
	given given
	want  result
}

func TestCases(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.run()
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			tt.run()
		}
	}
}
