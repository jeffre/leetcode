package leetcode2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// head's Val will be thrown away and Next will be properly set before
	// returning
	head := &ListNode{0, nil}
	current := head
	carry := 0

	// A non-empty ListNode or a non-zero carry require an additional ListNode
	// be appended
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		// Consume from l1
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		// Consume from l2
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// Create a new ListNode and give its address to the current node's
		// pointer
		current.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		// Each nodes only holds 1 digit so we need to carry anything more
		// than that into a future node
		carry = sum / 10

		// Cycle into the next (nil for now) node
		current = current.Next
	}

	// Use dummy head to return the first node
	return head.Next
}
