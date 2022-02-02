package leetcode21

import "fmt"

// ListNodeToInts converts ListNode to []int with a max length of 100 to
// prevent infinite loops
func ListNodeToInts(head *ListNode) []int {

	limit := 100

	res := []int{}
	for i := 0; head != nil; i++ {
		if i > limit {
			msg := fmt.Sprintf("ListNode limit (%v) reached!\n", limit)
			panic(msg)
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// IntsToListNode converts a []int into a ListNode
func IntsToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	dummy := &ListNode{}
	head := dummy
	for _, v := range nums {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return dummy.Next
}
