package leetcode104

/*
	Challenge Description: Given the root of a binary tree, return its maximum
	depth.

	A binary tree's maximum depth is the number of nodes along the longest path
	from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Iteration
// 780.0 ns/op
/*
func maxDepth(root *TreeNode) int {
	type qItem struct {
		Depth int
		Node  *TreeNode
	}
	var q []qItem

	if root != nil {
		q = append(q, qItem{1, root})
	}

	max := 0

	for len(q) > 0 {
		curDepth, node := q[0].Depth, q[0].Node
		q = q[1:]
		if node != nil {
			if curDepth > max {
				max = curDepth
			}
			q = append(q, qItem{Depth: curDepth + 1, Node: node.Left})
			q = append(q, qItem{Depth: curDepth + 1, Node: node.Right})
		}
	}

	return max
}
*/

// DFS
// 30 ns/op
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}
