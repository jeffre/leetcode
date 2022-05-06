package leetcode24

/*
  Given a linked list, swap every two adjacent nodes and return its head. You
  must solve the problem without modifying the values in the list's nodes
  (i.e., only nodes themselves may be changed.)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	// Return *ListNode as-is if the current or next node is nil
	if head == nil || head.Next == nil {
		return head
	}

	// Create newHead as second node
	newHead := head.Next

	// Change (old) head's Next to point to any remaining nodes
	head.Next = swapPairs(newHead.Next)

	// Change newHead's Next to point to (old) head
	newHead.Next = head

	return newHead
}
