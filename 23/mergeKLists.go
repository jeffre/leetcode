package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	dummy := &ListNode{0, nil}

	helper(lists, dummy)

	return dummy.Next
}

func helper(lists []*ListNode, result *ListNode) {

	// Remove nilled ListNode from 'lists'
	for i := 0; i < len(lists); {
		if lists[i] != nil {
			i++
			continue
		}

		if i < len(lists)-1 {
			copy(lists[i:], lists[i+1:])
		}

		lists[len(lists)-1] = nil
		lists = lists[:len(lists)-1]
	}

	switch len(lists) {
	case 0:
		// All lists have been fully processed
		return
	case 1:
		// All remaining nodes are in the same list so we can inherit the
		// parent and be done
		result.Next = lists[0]
		return
	}

	// Find which list has the lowest value
	lowest := 0
	for i := 1; i < len(lists); i++ {
		if lists[i].Val < lists[lowest].Val {
			lowest = i
		}
	}

	// Add lowest value to solution and advance that list
	result.Next = &ListNode{lists[lowest].Val, nil}
	result = result.Next
	lists[lowest] = lists[lowest].Next

	helper(lists, result)
}
