package leetcode21

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	list1 []int
	list2 []int
}

var cases = []struct {
	given given
	want  []int
}{
	{
		given: given{
			list1: []int{1, 2, 3},
			list2: []int{1, 3, 4},
		},
		want: []int{1, 1, 2, 3, 3, 4},
	}, {
		given: given{
			list1: nil,
			list2: nil,
		},
		want: nil,
	}, {
		given: given{
			list1: nil,
			list2: []int{1},
		},
		want: []int{1},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := mergeTwoLists(IntsToListNode(tt.given.list1), IntsToListNode(tt.given.list2))
			if !reflect.DeepEqual(got, IntsToListNode(tt.want)) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", tt.given, ListNodeToInts(got), tt.want)
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			mergeTwoLists(IntsToListNode(tt.given.list1), IntsToListNode(tt.given.list2))
		}
	}
}
