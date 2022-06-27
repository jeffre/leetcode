package leetcode27

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	nums []int
	val  int
}

type want struct {
	nums []int
	k    int
}

var cases = []struct {
	given given
	want  want
}{
	{
		given: given{
			nums: []int{3, 2, 2, 3},
			val:  3,
		},
		want: want{
			nums: []int{2, 2},
			k:    2,
		},
	},
	{
		given: given{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
		},
		want: want{
			nums: []int{0, 1, 3, 0, 4},
			k:    5,
		},
	},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := removeElement(tt.given.nums, tt.given.val)
			if got != tt.want.k {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want.k)
			}
			if !reflect.DeepEqual(tt.given.nums[:got], tt.want.nums) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", tt.given.nums[:got], tt.want.nums)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			removeElement(tt.given.nums, tt.given.val)
		}
	}
}
