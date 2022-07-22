package leetcode237

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {

	/*
		The two-pointer trick here is to have one iterate headA+headB while the
		other iterates headB+headA.
		This gives the pointers a common length which is an obstacle that would
		otherwise need to be accounted for.
	*/
	pA, pB := headA, headB
	for pA != pB {

		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
