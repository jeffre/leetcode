package leetcode349

import (
	"reflect"
	"testing"
)

type given struct {
	nums1 []int
	nums2 []int
}

var cases = []struct {
	name  string
	given given
	want  []int
}{
	{
		name: "example1",
		given: given{
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
		},
		want: []int{2},
	}, {
		name: "example2",
		given: given{
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
		},
		want: []int{9, 4},
	},
}

func TestCases(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.given.nums1, tt.given.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot:   %#v\nwant:  %#v\n", got, tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			intersection(tt.given.nums1, tt.given.nums2)
		}
	}
}
