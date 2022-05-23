package leetcode25

import (
	"fmt"
	"reflect"
	"testing"
)

type given struct {
	head *ListNode
	k    int
}

var cases = []struct {
	given given
	want  *ListNode
}{
	{
		given: given{
			head: &ListNode{1, nil},
			k:    1,
		},
		want: &ListNode{1, nil},
	},
	{
		given: given{
			head: &ListNode{1, &ListNode{2, nil}},
			k:    1,
		},
		want: &ListNode{1, &ListNode{2, nil}},
	},
	{
		given: given{
			head: &ListNode{1, &ListNode{2, nil}},
			k:    2,
		},
		want: &ListNode{2, &ListNode{1, nil}},
	},
}

func TestSwapPairs(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(ListToInts(tt.given.head), tt.given.k), func(t *testing.T) {
			got := reverseKGroup(tt.given.head, tt.given.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngiven: %#v\ngot:   %#v\nwant:  %#v\n", tt.given, got, tt.want)
				t.Errorf("got  %+v", ListToInts(got))
				t.Errorf("want %+v", ListToInts(tt.want))
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range cases {
			reverseKGroup(tt.given.head, tt.given.k)
		}
	}
}

/*
	Helpers for converting leetcode's sample cases into real linked lists
*/

// ListToInts convert List to []int
func ListToInts(head *ListNode) (result []int) {
	limit := 100
	times := 0

	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("The chain depth exceeds %d.", limit)
			panic(msg)
		}

		result = append(result, head.Val)
		head = head.Next
	}

	return result

}

// IntsToList convert []int to List
func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}
