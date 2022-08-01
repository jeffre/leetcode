package leetcode23

import (
	"fmt"
	"testing"
)

type test struct {
	name  string
	given []*ListNode
	want  *ListNode
}

func getTests() []test {

	return []test{
		{
			name:  "0 nodes",
			given: []*ListNode{},
			want:  nil,
		}, {
			name: "1 listnode",
			given: []*ListNode{
				intsToList([]int{1, 4, 5}),
			},
			want: intsToList([]int{1, 4, 5}),
		}, {
			name: "3 listnodes",
			given: []*ListNode{
				nil,
				nil,
				intsToList([]int{1, 4, 5}),
			},
			want: intsToList([]int{1, 4, 5}),
		}, {
			name: "simple-merge",
			given: []*ListNode{
				nil,
				intsToList([]int{6}),
				intsToList([]int{1, 4, 5}),
			},
			want: intsToList([]int{1, 4, 5, 6}),
		}, {
			name: "triple-merge",
			given: []*ListNode{
				intsToList([]int{1, 4, 5}),
				intsToList([]int{1, 3, 4}),
				intsToList([]int{2, 6}),
			},
			want: intsToList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		}, {
			name: "large-merge",
			given: []*ListNode{
				intsToList([]int{1, 4, 5}),
				intsToList([]int{1, 3, 4}),
				intsToList([]int{2, 6, 12, 33, 45}),
				intsToList([]int{2, 61, 63, 64, 66}),
				intsToList([]int{2, 6, 675, 676, 6657}),
				intsToList([]int{2, 6, 15, 17, 17}),
			},
			want: intsToList([]int{1, 1, 2, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6, 12, 15,
				17, 17, 33, 45, 61, 63, 64, 66, 675, 676, 6657}),
		}, {
			name: "negative-vals",
			given: []*ListNode{
				intsToList([]int{-1, 1}),
				intsToList([]int{-3, 1, 4}),
				intsToList([]int{-2, -1, 0, 2}),
			},
			want: intsToList([]int{-3, -2, -1, -1, 0, 1, 1, 2, 4}),
		},
	}
}

func TestCases(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {

			got := mergeKLists(tt.given)
			if !assertIdentical(t, got, tt.want) {
				t.Errorf("\ngiven: %v\ngot:   %v\nwant:  %v\n",
					listsToInts(tt.given), listToInts(got), listToInts(tt.want))
			}
		})
	}
}

func BenchmarkCases(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range getTests() {
			mergeKLists(tt.given)
		}
	}
}

func assertIdentical(t testing.TB, a, b *ListNode) bool {
	t.Helper()
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func listsToInts(nodes []*ListNode) [][]int {
	result := make([][]int, len(nodes))

	for i, n := range nodes {
		result[i] = listToInts(n)
	}
	return result
}

func listToInts(head *ListNode) (result []int) {
	limit := 100

	for head != nil {
		limit--
		if limit < 0 {
			msg := fmt.Sprintf("The chain depth exceeds %d. This may be the "+
				"result of an infinite loop.", limit)
			panic(msg)
		}

		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func intsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	sentinel := &ListNode{}
	cur := sentinel
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return sentinel.Next
}
