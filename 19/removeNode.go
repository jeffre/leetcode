package leetcode19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// Extract values from linked list
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	// Create new linked list
	newHead := &ListNode{Val: 0, Next: nil}
	current := newHead
	for i, v := range vals {

		// Skip node n from the end
		if len(vals)-i == n {
			continue
		}

		current.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		current = current.Next
	}
	return newHead.Next
}
