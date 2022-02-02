package leetcode21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// Early termination if either list is null
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Use recursive calls to cycle through both lists at the same time

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
