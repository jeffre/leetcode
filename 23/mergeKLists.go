package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)

	// Merge lists into list[0] (3844 ns/op)
	for i := 1; i < length; i++ {
		lists[0] = merge2Lists(lists[0], lists[i])
	}

	// Alternatively to the above merging method (3854 ns/op):
	// Merge lists down but reducing the number of repeated occurrances of any
	// given list. IE: merge lists as such: 0+1, 2+3, 4+5, 6+7; then 0+2, 4+6;
	// finally 0+4 for total of 7 mostly distinct passes.
	//
	// `interval` represents the distance between two lists.
	/*
		for interval := 1; interval < length; interval *= 2 {
			for i := 0; i < length-interval; i += interval * 2 {
				lists[i] = merge2Lists(lists[i], lists[i+interval])
			}
		}
	*/

	if length > 0 {
		return lists[0]
	}

	return nil
}

func merge2Lists(l1, l2 *ListNode) *ListNode {
	sentinel := &ListNode{0, nil}
	curr := sentinel

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}
	return sentinel.Next
}
