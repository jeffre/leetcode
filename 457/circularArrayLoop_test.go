package leetcode457

import (
	"testing"
)

var cases = []struct {
	name  string
	given []int
	want  bool
}{
	{
		name:  "example1",
		given: []int{2, -1, 1, 2, 2},
		want:  true,
	},
	{
		name:  "example2",
		given: []int{-1, 2},
		want:  false,
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := circularArrayLoop(tt.given)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

// Go modulus operator outputs values differently than other languages (inc
// python). This test appreciates that difference and shows how to use it
// for both +/- cycle directions in this leetcode challenge.
func TestMod(t *testing.T) {

	tests := []struct {
		name   string
		length int
		start  int
		move   int
		want   int
	}{
		{
			name:   "cycle-",
			length: 10,
			move:   -1,
			want:   9,
		},
		{
			name:   "cycle+",
			length: 10,
			move:   12,
			want:   2,
		},
		{
			name:   "step",
			length: 10,
			move:   2,
			want:   2,
		},
		{
			name:   "cycle++",
			length: 10,
			move:   2200002,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.start + tt.move
			if got < 0 {
				got = got%tt.length + tt.length
			} else if got >= tt.length {
				got %= tt.length
			}

			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			circularArrayLoop(tt.given)
		}
	}
}
