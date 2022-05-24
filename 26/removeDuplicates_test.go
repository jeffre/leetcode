package leetcode26

import (
	"fmt"
	"reflect"
	"testing"
)

type want struct {
	nums []int
	k    int
}

var cases = []struct {
	given []int
	want  want
}{
	{
		given: []int{1, 1, 2},
		want: want{
			nums: []int{1, 2},
			k:    2,
		},
	},
	{
		given: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		want: want{
			nums: []int{0, 1, 2, 3, 4},
			k:    5,
		},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := removeDuplicates(tt.given)
			if got != tt.want.k {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want.k)
			}
			if !reflect.DeepEqual(tt.given[:got], tt.want.nums) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", tt.given, tt.want.nums)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			removeDuplicates(tt.given)
		}
	}
}
