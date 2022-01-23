package leetcode19

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/*
		Examples below are based on the following input:
		head: [1, 2, 3, 4, 5]
		n: 2
	*/

	// Prepend 'head' with a dummy value
	dummy := &ListNode{Val: 0, Next: head} // [0, 1, 2, 3, 4, 5]

	// Fast forward into 'head' n times. To be continued...
	for ; n > 0; n-- {
		head = head.Next // [3, 4, 5]
	}

	// Create two new variables that will trail behind on future interations of
	// 'head'
	prev := dummy     // [0, 1, 2, 3, 4, 5]
	cur := dummy.Next // [1, 2, 3, 4, 5]

	// By Iterating 'head' until it runs out, 'prev' will be pointing at the
	// node that is to be removed, while 'cur' is one node ahead.
	for ; head != nil; head = head.Next {
		prev = prev.Next // [3, 4, 5]
		cur = cur.Next   // [4, 5]
	}

	// Instead of 'prev.Next' pointing to (what is now) 'cur', we can now skip
	// it by setting it to 'cur.Next'
	prev.Next = cur.Next

	// And because this change was done to a pointer that is part of the
	// 'dummy' ListNode, 'dummy.Next' just happens to be a convenient variable
	// to return.
	return dummy.Next
}

func printListNode(l *ListNode) {
	fmt.Printf("[")
	first := true
	for l != nil {
		if !first {
			fmt.Printf(", ")
		}
		fmt.Printf("%v", l.Val)
		first = false
		l = l.Next
	}
	fmt.Printf("]\n")
}
