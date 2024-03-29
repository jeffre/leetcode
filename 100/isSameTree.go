package leetcode100

/*
	Challenge Description: Given the roots of two binary trees p and q, write a
	function to check if they are the same or not.

	Two binary trees are considered the same if they are structurally identical,
	and the nodes have the same value.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Right, q.Right) && isSameTree(p.Left, q.Left)
}
