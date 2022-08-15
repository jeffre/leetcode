package leetcode116

/*
	Challenge Description: You are given a perfect binary tree where all leaves
	are on the same level, and every parent has two children. The binary tree
	has the following definition:

	struct Node {
		int val;
		Node *left;
		Node *right;
		Node *next;
	}

	Populate each next pointer to point to its next right node. If there is no
	next right node, the next pointer should be set to NULL.

	Initially, all next pointers are set to NULL.
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	q := []*Node{root.Left}
	q = append(q, root.Right)

	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[0]
			q = q[1:]

			if i != length-1 {
				node.Next = q[0]
			}
			if node.Left != nil {
				q = append(q, node.Left)
				q = append(q, node.Right)
			}
		}
	}
	return root
}
