package leetcode23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	dummy := &ListNode{0, nil}

	// Early return if no lists were provided
	if len(lists) == 0 {
		return dummy.Next
	}

	// Early return if 1 list was provided
	if len(lists) == 1 {
		return lists[0]
	}

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
		// Short circuit if all lists have been fully processed
		return
	case 1:
		// Short circuit if only one list is left (and inherit all nodes)
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
