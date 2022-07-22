package leetcode237

type ListNode struct {
	Val  int
	Next *ListNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	// Find lengths of a and b
	lengthA, lengthB := 0, 0
	for h := headA; h != nil; h = h.Next {
		lengthA++
	}
	for h := headB; h != nil; h = h.Next {
		lengthB++
	}

	// Determine the long and short lists
	short := headA
	long := headB
	if lengthA > lengthB {
		short = headB
		long = headA
	}

	// Iterate long list until its remaining node count is equal to the short
	// list
	for i := abs(lengthA - lengthB); i > 0; i-- {
		long = long.Next
	}

	// Iterate and compare for sameness
	for long != nil {
		if long == short {
			return long
		}
		long = long.Next
		short = short.Next
	}

	return nil
}
