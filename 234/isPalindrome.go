package leetcode141

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1, p2 := head, secondHalfStart
	result := true
	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1, p2 = p1.Next, p2.Next
	}

	// The last half of the linked-list was reversed. Because we should not
	// assume that the given linked-list is ok to modify we will undo the
	// reversal. Although, it should be noted then that this code is not thread
	// safe and should have exclusive access to the linked-list while running.
	firstHalfEnd.Next = reverseList(secondHalfStart)

	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		nextTemp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTemp
	}
	return prev
}
