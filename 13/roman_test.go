package leetcode13

import (
	"fmt"
	"testing"
)

var cases = []struct {
	want  int
	given string
}{
	{1, "I"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{12, "XII"},
	{13, "XIII"},
	{14, "XIV"},
	{49, "XLIX"},
	{50, "L"},
	{58, "LVIII"},
	{798, "DCCXCVIII"},
	{900, "CM"},
	{1000, "M"},
	{1006, "MVI"},
	{1984, "MCMLXXXIV"},
	{1994, "MCMXCIV"},
	{2014, "MMXIV"},
	{3999, "MMMCMXCIX"},
}

func TestCases(t *testing.T) {
	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := romanToInt(c.given)
			if got != c.want {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", c.given, got, c.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, c := range cases {
			romanToInt(c.given)
		}
	}
}
