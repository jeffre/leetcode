package leetcode25

/*
  Given the head of a linked list, reverse the nodes of the list k at a time,
  and return the modified list.

  k is a positive integer and is less than or equal to the length of the linked
  list. If the number of nodes is not a multiple of k then left-out nodes, in
  the end, should remain as it is.

  You may not alter the values in the list's nodes, only nodes themselves may
  be changed.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	start := head

	// Iterate head k times
	for i := 0; i < k; i++ {

		// If we run out of nodes return as-is
		if head == nil {
			return start
		}
		head = head.Next
	}

	// Create subset of ListNodes in reverse order
	newHead := reverseGroup(start, head)

	// "start" is now the end of the mini-list. So we do another round of
	// reverseKGroup to set its Next value
	start.Next = reverseKGroup(head, k)

	return newHead
}

// reverseGroup takes two ListNodes and iterates the first one until it matches
// the second one and returns the nodes in reverse order
func reverseGroup(start, end *ListNode) *ListNode {
	var newHead, tempHead *ListNode
	for start != end {
		tempHead = start.Next
		start.Next = newHead
		newHead = start
		start = tempHead
	}
	return newHead
}
