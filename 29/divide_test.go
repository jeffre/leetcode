package leetcode29

import (
	"fmt"
	"testing"
)

type given struct {
	dividend int
	divisor  int
}

var cases = []struct {
	given given
	want  int
}{
	{
		given: given{dividend: 10, divisor: 3},
		want:  3,
	}, {
		given: given{dividend: 7, divisor: -3},
		want:  -2,
	}, {
		given: given{dividend: 1, divisor: 1},
		want:  1,
	}, {
		given: given{dividend: -2147483648, divisor: -1},
		want:  2147483647,
	}, {
		given: given{dividend: -2147483648, divisor: 1},
		want:  -2147483648,
	},
}

func TestDivide(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := divide(tt.given.dividend, tt.given.divisor)
			if got != tt.want {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			divide(tt.given.dividend, tt.given.divisor)
		}
	}
}
