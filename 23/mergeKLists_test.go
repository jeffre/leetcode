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

func getTests() (tests []test) {

	tests = append(tests, test{
		name:  "0 nodes",
		given: []*ListNode{},
		want:  nil,
	})

	tests = append(tests, test{
		name: "1 listnode",
		given: []*ListNode{
			intsToList([]int{1, 4, 5}),
		},
		want: intsToList([]int{1, 4, 5}),
	})

	tests = append(tests, test{
		name: "3 listnodes",
		given: []*ListNode{
			nil,
			nil,
			intsToList([]int{1, 4, 5}),
		},
		want: intsToList([]int{1, 4, 5}),
	})

	tests = append(tests, test{
		name: "simple-merge",
		given: []*ListNode{
			nil,
			intsToList([]int{6}),
			intsToList([]int{1, 4, 5}),
		},
		want: intsToList([]int{1, 4, 5, 6}),
	})

	tests = append(tests, test{
		name: "triple-merge",
		given: []*ListNode{
			intsToList([]int{1, 4, 5}),
			intsToList([]int{1, 3, 4}),
			intsToList([]int{2, 6}),
		},
		want: intsToList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
	})

	tests = append(tests, test{
		name: "negative-vals",
		given: []*ListNode{
			intsToList([]int{-1, 1}),
			intsToList([]int{-3, 1, 4}),
			intsToList([]int{-2, -1, 0, 2}),
		},
		want: intsToList([]int{-3, -2, -1, -1, 0, 1, 1, 2, 4}),
	})

	return tests
}

func TestMergeKLists(t *testing.T) {
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
