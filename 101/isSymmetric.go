package leetcode101

/*
	Challenge Description: Given the root of a binary tree, check whether it is
	a mirror of itself (i.e., symmetric around its center).
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(Left, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}
	if Left == nil || Right == nil {
		return false
	}
	if Left.Val != Right.Val {
		return false
	}
	return isMirror(Left.Left, Right.Right) && isMirror(Left.Right, Right.Left)
}
