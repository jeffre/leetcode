package leetcode24

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	given *ListNode
	want  *ListNode
}{
	{
		given: nil,
		want:  nil,
	},
	{
		given: &ListNode{1, nil},
		want:  &ListNode{1, nil},
	},
	{
		given: &ListNode{1, &ListNode{2, nil}},
		want:  &ListNode{2, &ListNode{1, nil}},
	},
	{
		given: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
		want:  &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
	},
}

func TestSwapPairs(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.given), func(t *testing.T) {
			got := swapPairs(tt.given)
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
			swapPairs(tt.given)
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
